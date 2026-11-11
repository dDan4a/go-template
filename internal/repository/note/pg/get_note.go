package pg

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/dDan4a/go-template/internal/domain"
	"github.com/dDan4a/go-template/internal/repository"
)

func (r *Repository) GetNote(ctx context.Context, id int64) (domain.Note, error) {
	rows, _ := r.db.Query(ctx, "SELECT id, text, created_at FROM notes WHERE id = $1", id)

	dao, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[noteDao])
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Note{}, repository.ErrNoteNotFound
	}

	if err != nil {
		return domain.Note{}, fmt.Errorf("select note %d: %w", id, err)
	}

	return dao.toDomain(), nil
}
