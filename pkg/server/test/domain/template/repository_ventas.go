package template

import (
	"context"
	"reporteador/pkg/server/test/domain"
)

type VentasRepository interface {
	Save(ctx context.Context, reporte domain.ReporteVentas) error
	GetAll(ctx context.Context, limit, offset int) ([]domain.Venta, int, error)
}
