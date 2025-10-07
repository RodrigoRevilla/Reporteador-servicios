package command

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reporteador/pkg/server/test/domain"
	"reporteador/pkg/server/test/domain/template"
	"time"

	"go.uber.org/zap"
)

type GenerarReporteInventario struct {
	Inventario []domain.Inventario
	PDFPath    string
}

type GenerarReporteInventarioHandler interface {
	Handle(ctx context.Context, cmd *GenerarReporteInventario) error
}

type generarReporteInventarioHandler struct {
	logger *zap.Logger
}

func (h *generarReporteInventarioHandler) Handle(ctx context.Context, cmd *GenerarReporteInventario) error {
	h.logger.Info("Iniciando generación de reporte de inventario")

	baseFolder := "uuid_storage"
	if _, err := os.Stat(baseFolder); os.IsNotExist(err) {
		if err := os.Mkdir(baseFolder, os.ModePerm); err != nil {
			h.logger.Error("No se pudo crear uuid_storage", zap.Error(err))
			return fmt.Errorf("no se pudo crear uuid_storage: %w", err)
		}
	}

	pdfFileName := fmt.Sprintf("%s/reporte_inventario_%d.pdf", baseFolder, time.Now().Unix())

	absPath, err := filepath.Abs(pdfFileName)
	if err != nil {
		h.logger.Error("Error obteniendo ruta absoluta del PDF", zap.Error(err))
		return fmt.Errorf("error al obtener ruta absoluta del PDF: %w", err)
	}

	err = template.ReporteInventario(cmd.Inventario, absPath)
	if err != nil {
		h.logger.Error("Error generando PDF de inventario", zap.Error(err))
		return fmt.Errorf("error al generar el PDF: %w", err)
	}

	h.logger.Info("Reporte de inventario generado correctamente", zap.String("ruta_pdf", absPath))
	cmd.PDFPath = absPath

	go func(filePath string) {
		<-time.After(3 * time.Minute)
		err := os.Remove(filePath)
		if err != nil {
			h.logger.Error("Error eliminando el PDF automáticamente", zap.Error(err))
		} else {
			h.logger.Info("PDF eliminado automáticamente", zap.String("ruta_pdf", filePath))
		}
	}(absPath)

	return nil
}

func NewGenerarReporteInventarioHandler(logger *zap.Logger) GenerarReporteInventarioHandler {
	return &generarReporteInventarioHandler{
		logger: logger,
	}
}
