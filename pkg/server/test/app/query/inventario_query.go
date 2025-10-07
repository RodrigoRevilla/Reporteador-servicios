package query

import (
	"context"
	"reporteador/pkg/server/test/domain"
	"reporteador/pkg/server/test/domain/template"
)

type ObtenerReporteInventario struct {
	Limit  int
	Offset int
}

type ObtenerReporteInventarioResponse struct {
	Inventario []domain.Inventario `json:"inventario"`
	Limit      int                 `json:"limit"`
	Page       int                 `json:"page"`
	TotalItems int                 `json:"totalItems"`
	TotalPages int                 `json:"totalPages"`
}

type ObtenerReporteInventarioHandler interface {
	Handle(ctx context.Context, query ObtenerReporteInventario) (ObtenerReporteInventarioResponse, error)
}

type obtenerReporteInventarioHandler struct {
	repo template.InventarioRepository
}

func (h obtenerReporteInventarioHandler) Handle(ctx context.Context, query ObtenerReporteInventario) (ObtenerReporteInventarioResponse, error) {
	inventario := []domain.Inventario{}
	total := len(inventario)

	return ObtenerReporteInventarioResponse{
		Inventario: inventario,
		Limit:      query.Limit,
		Page:       query.Offset/query.Limit + 1,
		TotalItems: total,
		TotalPages: (total + query.Limit - 1) / query.Limit,
	}, nil
}

func NewObtenerReporteInventarioHandler(repo template.InventarioRepository) ObtenerReporteInventarioHandler {
	if repo == nil {
		panic("Nil repo")
	}
	return obtenerReporteInventarioHandler{repo: repo}
}
