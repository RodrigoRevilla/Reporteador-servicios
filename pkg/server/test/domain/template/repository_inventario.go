package template

import (
	"context"
	"reporteador/pkg/server/test/domain"
)

type InventarioRepository interface {
	Save(ctx context.Context, item domain.Inventario) error
	GetAll(ctx context.Context, limit, offset int) ([]domain.Inventario, int, error)
	GetByID(ctx context.Context, id string) (*domain.Inventario, error)
	Update(ctx context.Context, item domain.Inventario) error
	Delete(ctx context.Context, id string) error
}
