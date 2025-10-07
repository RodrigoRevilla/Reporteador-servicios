package service

import (
	"reporteador/pkg/server/test/app"
	"reporteador/pkg/server/test/app/command"

	"go.uber.org/zap"
)

func NewApplication() (*app.Application, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()

	return &app.Application{
		Commands: app.Commands{
			GenerarPDF:               command.NewGenerarPDFHandler(logger),
			GenerarReporteVentas:     command.NewGenerarReporteVentasHandler(logger),
			GenerarReporteInventario: command.NewGenerarReporteInventarioHandler(logger),
			GenerarRporteUsuarios:    command.NewGenerarReporteUsuariosActivosHandler(logger),
		},
	}, nil
}
