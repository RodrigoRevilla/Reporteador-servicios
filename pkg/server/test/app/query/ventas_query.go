package query

import (
	"context"
	"reporteador/pkg/server/test/domain"
	"reporteador/pkg/server/test/domain/template"
)

type ObtenerReporteVentas struct {
	Limit  int
	Offset int
}

type ObtenerReporteVentasResponse struct {
	Ventas     []domain.Venta `json:"ventas"`
	Limit      int            `json:"limit"`
	Page       int            `json:"page"`
	TotalItems int            `json:"totalItems"`
	TotalPages int            `json:"totalPages"`
}

type ObtenerReporteVentasHandler interface {
	Handle(ctx context.Context, query ObtenerReporteVentas) (ObtenerReporteVentasResponse, error)
}

type obtenerReporteVentasHandler struct {
	repo template.VentasRepository 
}

func (h obtenerReporteVentasHandler) Handle(ctx context.Context, query ObtenerReporteVentas) (ObtenerReporteVentasResponse, error) {
	ventas := []domain.Venta{} 
	total := len(ventas)

	return ObtenerReporteVentasResponse{
		Ventas:     ventas,
		Limit:      query.Limit,
		Page:       query.Offset/query.Limit + 1,
		TotalItems: total,
		TotalPages: (total + query.Limit - 1) / query.Limit, 
	}, nil
}

func NewObtenerReporteVentasHandler(repo template.VentasRepository) ObtenerReporteVentasHandler {
	if repo == nil {
		panic("Nil repo")
	}
	return obtenerReporteVentasHandler{repo: repo}
}
