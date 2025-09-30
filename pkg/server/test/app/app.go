package app

import (
	"reporteador/pkg/server/test/app/command"
)

type Commands struct {
	GenerarPDF command.GenerarPDFHandler
	GenerarReporteVentas command.GenerarReporteVentasHandler
}

type Application struct {
	Commands Commands
}
