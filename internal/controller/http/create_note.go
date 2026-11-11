package http

import (
	"encoding/json"
	"net/http"

	"github.com/dDan4a/go-template/internal/controller/http/converter"
)

func (c *Controller) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req converter.CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid JSON body"})
		return
	}

	note, err := c.service.CreateNote(r.Context(), req.Text)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, converter.NoteToResponse(note))
}
