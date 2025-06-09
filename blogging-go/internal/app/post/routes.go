package post

/*
- **GET /posts**: List posts (support pagination & search)
- **POST /posts**: Buat posting baru
- **GET /posts/:postId**: Detail posting
- **GET /posts/:slug**: Detail posting
- **PUT /posts/:postId**: Update posting
- **DELETE /posts/:postId**: Hapus posting
- **GET /posts/top-contributors**: Top Contributors
- **GET /posts/popular-topic**: Popular topic by tag
*/

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// RegisterPostRoutes registers all post-related routes
func RegisterPostRoutes(router *gin.RouterGroup, db *sqlx.DB) {
	// Posts routes group
	posts := router.Group("/posts")
	
	// GET /posts - Get all posts with filtering, sorting, and pagination
	posts.GET("", GetAllPostsHandler(db))
	
	// You can add more routes here later:
	// posts.GET("/top-contributors", GetTopContributorsHandler(db))
	// posts.GET("/popular-topic", GetPopularTopicHandler(db))
	// posts.GET("/:slug", GetPostBySlugHandler(db))
	// posts.GET("/:id", GetPostByIdHandler(db))
	// posts.POST("", CreatePostHandler(db))
	// posts.PUT("/:id", UpdatePostHandler(db))
	// posts.PATCH("/:id", UpdatePostHandler(db))
	// posts.DELETE("/:id", DeletePostHandler(db))
	// posts.POST("/:id/like", CreatePostLikeHandler(db))
	// posts.DELETE("/:id/like", DeletePostLikeHandler(db))
}