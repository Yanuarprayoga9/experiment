package comment

import (
	"time"

	"github.com/google/uuid"
	user "myblog/internal/app/user"
)

// KomentarBlog represents the blog comments table
type KomentarBlog struct {
	ID              uuid.UUID  `json:"id" db:"Id"`
	KomentarIndukID *uuid.UUID `json:"komentar_induk_id" db:"KomentarIndukId"`
	Konten          string     `json:"konten" db:"Konten"`
	CreatedAt       time.Time  `json:"created_at" db:"CreatedAt"`
	DeletedAt       *time.Time `json:"deleted_at" db:"DeletedAt"`

	Pengguna   *user.Pengguna      `json:"pengguna,omitempty"`
	Balasan    []KomentarBlog `json:"balasan,omitempty"`
	JumlahLike int            `json:"jumlah_like,omitempty"`
}

// KomentarLike represents the comment likes table
type KomentarLike struct {
	ID         uuid.UUID `json:"id" db:"Id"`
	KomentarID uuid.UUID `json:"komentar_id" db:"KomentarId"`
	PenggunaID uuid.UUID `json:"pengguna_id" db:"PenggunaId"`
	CreatedAt  time.Time `json:"created_at" db:"CreatedAt"`
}
