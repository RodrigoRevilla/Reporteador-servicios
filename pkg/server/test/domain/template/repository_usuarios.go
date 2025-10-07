package template

import (
	"context"
	"reporteador/pkg/server/test/domain"
)

type UsuariosActivosRepository interface {
	Save(ctx context.Context, usuario domain.UsuarioActivo) error
	GetAll(ctx context.Context) ([]domain.UsuarioActivo, error)
	GetByID(ctx context.Context, id string) (*domain.UsuarioActivo, error)
	Update(ctx context.Context, usuario domain.UsuarioActivo) error
	Delete(ctx context.Context, id string) error
}