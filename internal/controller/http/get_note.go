package http

import (
	"net/http"
	"strconv"

	"github.com/dDan4a/go-template/internal/controller/http/converter"
)

func (c *Controller) GetNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "id must be an integer"})
		return
	}

	note, err := c.service.GetNote(r.Context(), id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, converter.NoteToResponse(note))
}
