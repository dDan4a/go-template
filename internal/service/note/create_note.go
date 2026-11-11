package note

import (
	"context"
	"errors"
	"strings"

	"github.com/dDan4a/go-template/internal/domain"
)

var ErrEmptyText = errors.New("note text is empty")

func (s *Service) CreateNote(ctx context.Context, text string) (domain.Note, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return domain.Note{}, ErrEmptyText
	}

	return s.noteRepo.CreateNote(ctx, text)
}
