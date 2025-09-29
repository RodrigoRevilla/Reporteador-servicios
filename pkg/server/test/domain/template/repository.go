package template

import (
	"context"
	"fmt"
	"reporteador/pkg/server/test/domain"
)

type NotFoundError struct {
	pdf string
}
type NotFoundDataError struct {
	Message string
}

type NotFoundIdError struct {
	ID int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("pdf no encontrado", e.pdf)
}

type Repository interface {
	Save(ctx context.Context, pdf domain.Template) error
}
