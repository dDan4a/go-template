package http

import (
	"net/http"

	noteSvc "github.com/dDan4a/go-template/internal/service/note"
)

type Controller struct {
	service *noteSvc.Service
}

func New(service *noteSvc.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", c.Health)
	mux.HandleFunc("POST /notes", c.CreateNote)
	mux.HandleFunc("GET /notes/{id}", c.GetNote)

	return mux
}
