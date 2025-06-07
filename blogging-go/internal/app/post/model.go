package post

import (
	"time"
	user "myblog/internal/app/user"

	"github.com/google/uuid"
)

// Tag represents the tags table
type Tag struct {
	ID      uuid.UUID `json:"id" db:"Id"`
	NamaTag string    `json:"nama_tag" db:"NamaTag"`
}

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
	Penulis      *user.Pengguna `json:"penulis,omitempty"`
	Tags         []Tag     `json:"tags,omitempty"`
	JumlahLike   int       `json:"jumlah_like,omitempty"`
	JumlahKomentar int     `json:"jumlah_komentar,omitempty"`
}