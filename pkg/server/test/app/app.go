package app

import (
	"reporteador/pkg/server/test/app/command"
)

type Commands struct {
	GenerarPDF command.GenerarPDFHandler
}

type Application struct {
	Commands Commands
}
