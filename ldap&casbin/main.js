// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/casbin/casbin/v2"
// 	"github.com/casbin/casbin/v2/model"
// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/go-ldap/ldap/v3"
// )

// var jwtKey = []byte("secret_key_jwt")

// // JWT claims structure
// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// var enforcer *casbin.Enforcer

// func main() {
// 	initCasbin()

// 	http.HandleFunc("/login", loginHandler)
// 	http.Handle("/public", authMiddleware(http.HandlerFunc(publicHandler)))
// 	http.Handle("/private", authMiddleware(http.HandlerFunc(privateHandler)))
// 	http.Handle("/admin", authMiddleware(http.HandlerFunc(adminHandler)))

// 	log.Println("Server running on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func initCasbin() {
// 	textModel := `
// [request_definition]
// r = sub, obj, act

// [policy_definition]
// p = sub, obj, act

// [role_definition]
// g = _, _

// [policy_effect]
// e = some(where (p.eft == allow))

// [matchers]
// m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act

// `
// 	m, _ := model.NewModelFromString(textModel)
// 	e, err := casbin.NewEnforcer(m)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Tambahkan policies
// added1, err1 := e.AddPolicy("admin", "/private", "read")
// added2, err2 := e.AddPolicy("admin", "/admin", "write")
// added3, err3 := e.AddPolicy("user", "/public", "read")

// fmt.Println("AddPolicy admin /private read:", added1, err1)
// fmt.Println("AddPolicy admin /admin write:", added2, err2)
// fmt.Println("AddPolicy user /public read:", added3, err3)

// // Tambahkan role bindings
// grouped1, err4 := e.AddGroupingPolicy("alice", "admin")
// grouped2, err5 := e.AddGroupingPolicy("bob", "user")

// fmt.Println("AddGroupingPolicy alice admin:", grouped1, err4)
// fmt.Println("AddGroupingPolicy bob user:", grouped2, err5)

// // Cek seluruh policy yang ada
// policies,_ := e.GetPolicy()
// groupings ,_ := e.GetGroupingPolicy()

// fmt.Println("Policies:")
// for _, p := range policies {
//     fmt.Println(p)
// }

// fmt.Println("Grouping Policies:")
// for _, g := range groupings {
//     fmt.Println(g)
// }

// 	enforcer = e
// }

// // LDAP authentication
// func authenticateLDAP(username, password string) bool {
// 	ldapServer := "ldap://localhost:389"
// 	baseDN := "dc=example,dc=org"
// 	// bindUsername := fmt.Sprintf("uid=%s,ou=users,%s", username "admin", baseDN)
// 	bindUsername := fmt.Sprintf("uid=%s,ou=users,%s",  "admin", baseDN)

// 	conn, err := ldap.DialURL(ldapServer)
// 	if err != nil {
// 		log.Println("LDAP connection error:", err)
// 		return false
// 	}
// 	defer conn.Close()

// 	err = conn.Bind(bindUsername, password)
// 	if err != nil {
// 		log.Println("LDAP authentication failed:", err)
// 		return false
// 	}
// 	return true
// }

// // Login handler
// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	type LoginRequest struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	var req LoginRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	if !authenticateLDAP(req.Username, req.Password) {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	expirationTime := time.Now().Add(1 * time.Hour)
// 	claims := &Claims{
// 		Username: req.Username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"token": tokenString,
// 	})
// }

// // Middleware for JWT & Casbin RBAC
// func authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		claims := &Claims{}

// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 		if err != nil || !token.Valid {
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		username := claims.Username
// 		obj := r.URL.Path
// 		act := methodToAction(r.Method)
// 		if act == "" {
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 			return
// 		}

// 		ok, err := enforcer.Enforce(username, obj, act)
// 		if err != nil {
// 			http.Error(w, "Internal server error", http.StatusInternalServerError)
// 			return
// 		}
// 		if !ok {
// 			http.Error(w, "Forbidden", http.StatusForbidden)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// // Method to action mapper
// func methodToAction(method string) string {
// 	switch method {
// 	case "GET":
// 		return "read"
// 	case "POST", "PUT", "PATCH":
// 		return "write"
// 	case "DELETE":
// 		return "delete"
// 	default:
// 		return ""
// 	}
// }

// // Public handler
// func publicHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Public resource accessible by user role")
// }

// // Private handler
// func privateHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Private resource accessible only by admin")
// }

// // Admin handler
// func adminHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Admin area: write access required")
// }
