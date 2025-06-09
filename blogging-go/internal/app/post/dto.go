package post

import (
	"myblog/internal/app/user"
	"time"

	"github.com/google/uuid"
)

// PostinganLike represents the post likes table
type CreatePostinganRequestLikeRequest struct {
	
}
// CreatePostinganRequest represents the request to create a new blog post
type CreatePostinganRequest struct {
	Judul        string   `json:"judul" validate:"required,min=1,max=255"`
	Slug         string   `json:"slug" validate:"required,min=1,max=255"`
	Isi          string   `json:"isi" validate:"required"`
	Visibilitas  string   `json:"visibilitas" validate:"required,oneof=internal publik terbatas"`
	URLThumbnail *string  `json:"url_thumbnail"`
	Tags         []string `json:"tags"`
	Kolaborators []string `json:"kolaborators"`
}

// UpdatePostinganRequest represents the request to update a blog post
type UpdatePostinganRequest struct {
	Judul        *string  `json:"judul" validate:"omitempty,min=1,max=255"`
	Slug         *string  `json:"slug" validate:"omitempty,min=1,max=255"`
	Isi          *string  `json:"isi" validate:"omitempty"`
	Visibilitas  *string  `json:"visibilitas" validate:"omitempty,oneof=internal publik terbatas"`
	URLThumbnail *string  `json:"url_thumbnail"`
	Tags         []string `json:"tags"`
	Kolaborators []string `json:"kolaborators"`
}


type PostinganBlogListResponse struct {
	ID           uuid.UUID  `json:"id" db:"Id"`
	Judul        string     `json:"judul" db:"Judul"`
	Slug         string     `json:"slug" db:"Slug"`
	Isi          string     `json:"isi" db:"Isi"`
	StatusReview string     `json:"status_review" db:"StatusReview"`
	Visibilitas  string     `json:"visibilitas" db:"Visibilitas"`
	JumlahTayang int        `json:"jumlah_tayang" db:"JumlahTayang"`
	URLThumbnail *string    `json:"url_thumbnail" db:"UrlThumbnail"`
	CreatedAt    time.Time  `json:"created_at" db:"CreatedAt"`
	UpdatedAt    time.Time  `json:"updated_at" db:"UpdatedAt"`
	DeletedAt    *time.Time `json:"deleted_at" db:"DeletedAt"`
	
	// Relations
	Penulis      *user.Pengguna `json:"penulis,omitempty"`
	Tags         []string     `json:"tags,omitempty"`
	JumlahLike   int       `json:"jumlah_like,omitempty"`
	JumlahKomentar int     `json:"jumlah_komentar,omitempty"`
}
type GetAllPostinganBlogListResponse struct {
	Posts []PostinganBlogListResponse `json:"data"`
	Total int    `json:"total"`
}


type PostinganBlogDetailResponse struct {
	ID           uuid.UUID  `json:"id" db:"Id"`
	Judul        string     `json:"judul" db:"Judul"`
	Slug         string     `json:"slug" db:"Slug"`
	Isi          string     `json:"isi" db:"Isi"`
	StatusReview string     `json:"status_review" db:"StatusReview"`
	Visibilitas  string     `json:"visibilitas" db:"Visibilitas"`
	JumlahTayang int        `json:"jumlah_tayang" db:"JumlahTayang"`
	URLThumbnail *string    `json:"url_thumbnail" db:"UrlThumbnail"`
	CreatedAt    time.Time  `json:"created_at" db:"CreatedAt"`
	UpdatedAt    time.Time  `json:"updated_at" db:"UpdatedAt"`
	DeletedAt    *time.Time `json:"deleted_at" db:"DeletedAt"`
	
	// Relations
	Penulis      *user.Pengguna `json:"penulis,omitempty"`
	Tags         []string     `json:"tags,omitempty"`
	Kolaborators []KolaboratorBlog `json:"kolaborators,omitempty"`
	JumlahLike   int       `json:"jumlah_like,omitempty"`
	JumlahKomentar int     `json:"jumlah_komentar,omitempty"`
}
