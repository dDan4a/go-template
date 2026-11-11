package pg

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dDan4a/go-template/internal/repository"
)

var _ repository.Note = (*Repository)(nil)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
