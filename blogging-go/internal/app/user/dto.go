package user

import (
	"time"

	"github.com/google/uuid"
)

// Pengguna represents the users table
type PenggunaResponse struct {
	ID         uuid.UUID  `json:"id" db:"Id"`
	Nip        string     `json:"nip" db:"Nip"`
	Nama       string     `json:"nama" db:"Nama"`
	Email      string     `json:"email" db:"Email"`
	Peran      *string    `json:"peran" db:"Peran"`
	Jabatan    *string    `json:"jabatan" db:"Jabatan"`
	UnitKerja  *string    `json:"unit_kerja" db:"UnitKerja"`
	DihapusPada *time.Time `json:"dihapus_pada" db:"DihapusPada"`
	DibuatPada time.Time  `json:"dibuat_pada" db:"DibuatPada"`
}
