package note

import "github.com/dDan4a/go-template/internal/repository"

type Service struct {
	noteRepo repository.Note
}

func New(noteRepo repository.Note) *Service {
	return &Service{
		noteRepo: noteRepo,
	}
}
