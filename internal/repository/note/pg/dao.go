package pg

import (
	"time"

	"github.com/dDan4a/go-template/internal/domain"
)

type noteDao struct {
	ID        int64     `db:"id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func (d noteDao) toDomain() domain.Note {
	return domain.Note{
		ID:        d.ID,
		Text:      d.Text,
		CreatedAt: d.CreatedAt,
	}
}
