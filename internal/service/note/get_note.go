package note

import (
	"context"

	"github.com/dDan4a/go-template/internal/domain"
)

func (s *Service) GetNote(ctx context.Context, id int64) (domain.Note, error) {
	return s.noteRepo.GetNote(ctx, id)
}
