package models

import (
	"time"

	"github.com/google/uuid"
)


// PostinganBlog represents the blog posts table
type PostinganBlog struct {
	ID           uuid.UUID  `json:"id" db:"Id"`
	Judul        string     `json:"judul" db:"Judul"`
	Slug         string     `json:"slug" db:"Slug"`
	Isi          string     `json:"isi" db:"Isi"`
	PenulisID    uuid.UUID  `json:"penulis_id" db:"PenulisId"`
	StatusReview string     `json:"status_review" db:"StatusReview"`
	Visibilitas  string     `json:"visibilitas" db:"Visibilitas"`
	JumlahTayang int        `json:"jumlah_tayang" db:"JumlahTayang"`
	URLThumbnail *string    `json:"url_thumbnail" db:"UrlThumbnail"`
	CreatedAt    time.Time  `json:"created_at" db:"CreatedAt"`
	UpdatedAt    time.Time  `json:"updated_at" db:"UpdatedAt"`
	DeletedAt    *time.Time `json:"deleted_at" db:"DeletedAt"`
	
	// Relations
	Penulis      *Pengguna `json:"penulis,omitempty"`
	Tags         []Tag     `json:"tags,omitempty"`
	JumlahLike   int       `json:"jumlah_like,omitempty"`
	JumlahKomentar int     `json:"jumlah_komentar,omitempty"`
}

// PostinganLike represents the post likes table
type PostinganLike struct {
	ID          uuid.UUID `json:"id" db:"Id"`
	PostinganID uuid.UUID `json:"postingan_id" db:"PostinganId"`
	PenggunaID  uuid.UUID `json:"pengguna_id" db:"PenggunaId"`
}

// Tag represents the tags table
type Tag struct {
	ID      uuid.UUID `json:"id" db:"Id"`
	NamaTag string    `json:"nama_tag" db:"NamaTag"`
}

// PostinganTag represents the many-to-many relationship between posts and tags
type PostinganTag struct {
	PostinganID uuid.UUID `json:"postingan_id" db:"PostinganId"`
	TagID       uuid.UUID `json:"tag_id" db:"TagId"`
}

// KolaboratorBlog represents the blog collaborators table
type KolaboratorBlog struct {
	ID                  uuid.UUID  `json:"id" db:"Id"`
	PostinganID         uuid.UUID  `json:"postingan_id" db:"PostinganId"`
	PenggunaID          uuid.UUID  `json:"pengguna_id" db:"PenggunaId"`
	PeranDalamPostingan string     `json:"peran_dalam_postingan" db:"PeranDalamPostingan"`
	CreatedAt           time.Time  `json:"created_at" db:"CreatedAt"`
	DeletedAt           *time.Time `json:"deleted_at" db:"DeletedAt"`
	
	// Relations
	Pengguna *Pengguna `json:"pengguna,omitempty"`
}

// KomentarBlog represents the blog comments table
type KomentarBlog struct {
	ID              uuid.UUID  `json:"id" db:"Id"`
	PostinganID     uuid.UUID  `json:"postingan_id" db:"PostinganId"`
	PenggunaID      uuid.UUID  `json:"pengguna_id" db:"PenggunaId"`
	KomentarIndukID *uuid.UUID `json:"komentar_induk_id" db:"KomentarIndukId"`
	Konten          string     `json:"konten" db:"Konten"`
	CreatedAt       time.Time  `json:"created_at" db:"CreatedAt"`
	DeletedAt       *time.Time `json:"deleted_at" db:"DeletedAt"`
	
	// Relations
	Pengguna    *Pengguna      `json:"pengguna,omitempty"`
	Balasan     []KomentarBlog `json:"balasan,omitempty"`
	JumlahLike  int            `json:"jumlah_like,omitempty"`
}

// KomentarLike represents the comment likes table
type KomentarLike struct {
	ID         uuid.UUID `json:"id" db:"Id"`
	KomentarID uuid.UUID `json:"komentar_id" db:"KomentarId"`
	PenggunaID uuid.UUID `json:"pengguna_id" db:"PenggunaId"`
	CreatedAt  time.Time `json:"created_at" db:"CreatedAt"`
}

// StatistikBlog represents the blog statistics table
type StatistikBlog struct {
	ID             uuid.UUID `json:"id" db:"Id"`
	PostinganID    uuid.UUID `json:"postingan_id" db:"PostinganId"`
	TanggalStat    time.Time `json:"tanggal_stat" db:"TanggalStat"`
	JumlahTayang   int       `json:"jumlah_tayang" db:"JumlahTayang"`
	JumlahKomentar int       `json:"jumlah_komentar" db:"JumlahKomentar"`
	JumlahSuka     int       `json:"jumlah_suka" db:"JumlahSuka"`
}

// CreatePostinganRequest represents the request to create a new blog post
type CreatePostinganRequest struct {
	Judul        string    `json:"judul" validate:"required,min=1,max=255"`
	Slug         string    `json:"slug" validate:"required,min=1,max=255"`
	Isi          string    `json:"isi" validate:"required"`
	Visibilitas  string    `json:"visibilitas" validate:"required,oneof=internal publik terbatas"`
	URLThumbnail *string   `json:"url_thumbnail"`
	Tags         []string  `json:"tags"`
}

// UpdatePostinganRequest represents the request to update a blog post
type UpdatePostinganRequest struct {
	Judul        *string   `json:"judul" validate:"omitempty,min=1,max=255"`
	Slug         *string   `json:"slug" validate:"omitempty,min=1,max=255"`
	Isi          *string   `json:"isi" validate:"omitempty"`
	Visibilitas  *string   `json:"visibilitas" validate:"omitempty,oneof=internal publik terbatas"`
	URLThumbnail *string   `json:"url_thumbnail"`
	Tags         []string  `json:"tags"`
}

// CreateKomentarRequest represents the request to create a new comment
type CreateKomentarRequest struct {
	PostinganID     uuid.UUID  `json:"postingan_id" validate:"required"`
	KomentarIndukID *uuid.UUID `json:"komentar_induk_id"`
	Konten          string     `json:"konten" validate:"required,min=1"`
}

// CreateKolaboratorRequest represents the request to add a collaborator
type CreateKolaboratorRequest struct {
	PostinganID         uuid.UUID `json:"postingan_id" validate:"required"`
	PenggunaID          uuid.UUID `json:"pengguna_id" validate:"required"`
	PeranDalamPostingan string    `json:"peran_dalam_postingan" validate:"required,oneof=penulisBersama editor"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token    string    `json:"token"`
	Pengguna *Pengguna `json:"pengguna"`
}