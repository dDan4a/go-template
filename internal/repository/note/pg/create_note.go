package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/dDan4a/go-template/internal/domain"
)

func (r *Repository) CreateNote(ctx context.Context, text string) (domain.Note, error) {
	rows, _ := r.db.Query(ctx, "INSERT INTO notes (text) VALUES ($1) RETURNING id, text, created_at", text)

	dao, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[noteDao])
	if err != nil {
		return domain.Note{}, fmt.Errorf("insert note: %w", err)
	}

	return dao.toDomain(), nil
}
