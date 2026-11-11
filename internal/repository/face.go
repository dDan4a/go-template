package repository

//go:generate minimock -o ../tests/mocks -s _mock.go

import (
	"context"
	"errors"

	"github.com/dDan4a/go-template/internal/domain"
)

var ErrNoteNotFound = errors.New("note not found")

type Note interface {
	CreateNote(ctx context.Context, text string) (domain.Note, error)
	GetNote(ctx context.Context, id int64) (domain.Note, error)
}
