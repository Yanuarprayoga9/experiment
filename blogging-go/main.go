// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/sql-adapter/v2"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

var db *sql.DB
var enforcer *casbin.Enforcer

func main() {
	var err error
	dsn := "sqlserver://sa:YourPassword@localhost:1433?database=blog"
	db, err = sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	enforcer, err = InitCasbin(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize Casbin: %v", err)
	}

	SyncPolicies(enforcer)

	r := gin.Default()
	r.POST("/posts", CreatePost)
	r.GET("/posts", ListPosts)

	r.Run(":8080")
}

func InitCasbin(dsn string) (*casbin.Enforcer, error) {
	a, err := adapter.NewAdapter("sqlserver", dsn, true)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("casbin_model.conf", a)
	if err != nil {
		return nil, err
	}
	e.LoadPolicy()
	return e, nil
}

func SyncPolicies(e *casbin.Enforcer) {
	// Clear existing policies
	e.ClearPolicy()

	// Load group -> permission
	rows, err := db.Query(`
		SELECT g.name, p.object, p.action
		FROM group_permissions gp
		JOIN user_groups g ON gp.group_id = g.id
		JOIN permissions p ON p.id = gp.permission_id
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var groupName, obj, act string
		rows.Scan(&groupName, &obj, &act)
		e.AddPolicy(groupName, obj, act)
	}

	// Load user -> group
	rows2, err := db.Query(`
		SELECT u.ldap_uid, g.name
		FROM user_group_membership m
		JOIN users u ON m.user_id = u.id
		JOIN user_groups g ON m.group_id = g.id
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()
	for rows2.Next() {
		var uid, group string
		rows2.Scan(&uid, &group)
		e.AddGroupingPolicy(uid, group)
	}
	e.SavePolicy()
}

func CreatePost(c *gin.Context) {
	uid := c.GetHeader("X-User-Id")
	ok, err := enforcer.Enforce(uid, "post", "create")
	if err != nil || !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
		return
	}

	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserID = 1

	_, err = db.Exec("INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)", post.Title, post.Content, post.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func ListPosts(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, content, user_id FROM posts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, p)
	}
	c.JSON(http.StatusOK, posts)
}

// casbin_model.conf
// ------------------
/*
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
*/
