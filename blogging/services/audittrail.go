package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
	"github.com/rs/zerolog/log"
)

func ParseMssqlUUID(id string) *mssql.UniqueIdentifier {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		log.Error().Err(err).Msg("Invalid UUID")
	}
	var parsedId mssql.UniqueIdentifier
	copy(parsedId[:], parsedUUID[:])
	return &parsedId
}

var hariIndo = [...]string{
	"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu",
}

var bulanIndo = [...]string{
	"", "Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

func FormatTanggalIndo(t *time.Time) string {
	if t == nil {
		return "-"
	}
	day := hariIndo[t.Weekday()]
	date := t.Day()
	month := bulanIndo[int(t.Month())]
	year := t.Year()

	return fmt.Sprintf("%s, %d %s %d", day, date, month, year)
}
