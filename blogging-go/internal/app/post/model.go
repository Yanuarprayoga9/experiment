package post

import (
	user "myblog/internal/app/user"
	"time"

	"github.com/google/uuid"
)

// KolaboratorBlog represents the blog collaborators table
type KolaboratorBlog struct {
	ID                  uuid.UUID  `json:"id" db:"Id"`
	PostinganID         uuid.UUID  `json:"postingan_id" db:"PostinganId"`
	PenggunaID          uuid.UUID  `json:"pengguna_id" db:"PenggunaId"`
	PeranDalamPostingan string     `json:"peran_dalam_postingan" db:"PeranDalamPostingan"`
	CreatedAt           time.Time  `json:"created_at" db:"CreatedAt"`
	DeletedAt           *time.Time `json:"deleted_at" db:"DeletedAt"`

	// Relations
	Pengguna *user.Pengguna `json:"pengguna,omitempty"`
}

// Tag represents the tags table
type Tag struct {
	ID      uuid.UUID `json:"id" db:"Id"`
	NamaTag string    `json:"nama_tag" db:"NamaTag"`
}

type PaginationInfo struct {
	CurrentPage int
	PageSize    int
	TotalItems  int
	TotalPages  int
}

// PostinganBlog represents the blog posts table
type PostinganBlog struct {
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

// PostinganLike represents the post likes table
type PostinganLike struct {
	ID          uuid.UUID `json:"id" db:"Id"`
	PostinganID uuid.UUID `json:"postingan_id" db:"PostinganId"`
	PenggunaID  uuid.UUID `json:"pengguna_id" db:"PenggunaId"`
}



