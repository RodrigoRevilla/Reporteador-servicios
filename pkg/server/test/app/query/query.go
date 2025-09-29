package query

import (
	"context"
	"reporteador/pkg/server/test/domain/template"
)

type ObtenerPDF struct {
	Limit  int
	Offset int
}

type ObtenerPDFResponse struct {
	Data       []*template.Data `json:"data"`
	Limit      int              `json:"limit"`
	Page       int              `json:"page"`
	TotalItems int              `json:"totalItems"`
	TotalPages int              `json:"totalPages"`
}

type ObtenerPDFHandler interface {
	Handle(ctx context.Context, query ObtenerPDF) (ObtenerPDFResponse, error)
}

type obtenerPDFHandler struct {
	repo template.Repository
}

func (h obtenerPDFHandler) Handle(ctx context.Context, query ObtenerPDF) (ObtenerPDFResponse, error) {
	return ObtenerPDFResponse{
		Data:       []*template.Data{},
		Limit:      query.Limit,
		Page:       query.Offset/query.Limit + 1,
		TotalItems: 0,
		TotalPages: 0,
	}, nil
}

func NewObtenerPDFHandler(repo template.Repository) ObtenerPDFHandler {
	if repo == nil {
		panic("Nil repo")
	}
	return obtenerPDFHandler{repo: repo}
}
