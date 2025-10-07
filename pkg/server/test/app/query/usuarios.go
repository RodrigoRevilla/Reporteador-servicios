package query

import (
	"context"
	"reporteador/pkg/server/test/domain"
	"reporteador/pkg/server/test/domain/template"
)

type ObtenerUsuariosActivos struct {
	Limit  int
	Offset int
}

type ObtenerUsuariosActivosResponse struct {
	Usuarios   []domain.UsuarioActivo `json:"usuarios"`
	Limit      int                    `json:"limit"`
	Page       int                    `json:"page"`
	TotalItems int                    `json:"totalItems"`
	TotalPages int                    `json:"totalPages"`
}

type ObtenerUsuariosActivosHandler interface {
	Handle(ctx context.Context, query ObtenerUsuariosActivos) (ObtenerUsuariosActivosResponse, error)
}

type obtenerUsuariosActivosHandler struct {
	repo template.UsuariosActivosRepository
}

func (h obtenerUsuariosActivosHandler) Handle(ctx context.Context, query ObtenerUsuariosActivos) (ObtenerUsuariosActivosResponse, error) {
	usuarios := []domain.UsuarioActivo{}
	total := len(usuarios)

	return ObtenerUsuariosActivosResponse{
		Usuarios:   usuarios,
		Limit:      query.Limit,
		Page:       query.Offset/query.Limit + 1,
		TotalItems: total,
		TotalPages: (total + query.Limit - 1) / query.Limit,
	}, nil
}

func NewObtenerUsuariosActivosHandler(repo template.UsuariosActivosRepository) ObtenerUsuariosActivosHandler {
	if repo == nil {
		panic("Nil repo")
	}
	return obtenerUsuariosActivosHandler{repo: repo}
}
