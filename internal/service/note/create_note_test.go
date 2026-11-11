package note_test

import (
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"

	"github.com/dDan4a/go-template/internal/domain"
	"github.com/dDan4a/go-template/internal/service/note"
	"github.com/dDan4a/go-template/internal/tests/mocks"
)

func TestService_CreateNote(t *testing.T) {
	t.Parallel()

	t.Run("trims text before saving", func(t *testing.T) {
		t.Parallel()

		want := domain.Note{ID: 1, Text: "buy milk"}
		repo := mocks.NewNoteMock(t).CreateNoteMock.Expect(minimock.AnyContext, "buy milk").Return(want, nil)

		got, err := note.New(repo).CreateNote(t.Context(), "  buy milk\n")
		if err != nil {
			t.Fatalf("CreateNote: %v", err)
		}

		if got != want {
			t.Errorf("CreateNote = %+v, want %+v", got, want)
		}
	})

	t.Run("rejects blank text without touching the repository", func(t *testing.T) {
		t.Parallel()

		_, err := note.New(mocks.NewNoteMock(t)).CreateNote(t.Context(), " \t")
		if !errors.Is(err, note.ErrEmptyText) {
			t.Errorf("CreateNote error = %v, want %v", err, note.ErrEmptyText)
		}
	})
}
