package converter

import (
	"time"

	"github.com/dDan4a/go-template/internal/domain"
)

type CreateNoteRequest struct {
	Text string `json:"text"`
}

type NoteResponse struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func NoteToResponse(in domain.Note) NoteResponse {
	return NoteResponse{
		ID:        in.ID,
		Text:      in.Text,
		CreatedAt: in.CreatedAt,
	}
}
