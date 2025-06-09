package comment

import (
	"time"

	user "myblog/internal/app/user"
	"github.com/google/uuid"
)

type CreateKomentarRequest struct {
	Konten          string     `json:"konten" validate:"required"`
	KomentarIndukID *uuid.UUID `json:"komentar_induk_id,omitempty"` // untuk komentar balasan
}

type KomentarResponse struct {
	ID              uuid.UUID          `json:"id"`
	Konten          string             `json:"konten"`
	CreatedAt       time.Time          `json:"created_at"`
	DeletedAt       *time.Time         `json:"deleted_at,omitempty"`
	KomentarIndukID *uuid.UUID         `json:"komentar_induk_id,omitempty"`
	Pengguna        *user.PenggunaResponse  `json:"pengguna,omitempty"`
	Balasan         []KomentarResponse `json:"balasan,omitempty"`
	JumlahLike      int                `json:"jumlah_like"`
}