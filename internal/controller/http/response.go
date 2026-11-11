package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/dDan4a/go-template/internal/repository"
	noteSvc "github.com/dDan4a/go-template/internal/service/note"
)

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Error("write response", slog.Any("error", err))
	}
}

func writeError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, repository.ErrNoteNotFound):
		writeJSON(w, http.StatusNotFound, errorResponse{Error: err.Error()})
	case errors.Is(err, noteSvc.ErrEmptyText):
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
	default:
		slog.Error("handle request", slog.Any("error", err))
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: http.StatusText(http.StatusInternalServerError)})
	}
}
