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

type GenerarReporteUsuariosActivos struct {
	Usuarios []domain.UsuarioActivo
	PDFPath  string
}

type GenerarReporteUsuariosActivosHandler interface {
	Handle(ctx context.Context, cmd *GenerarReporteUsuariosActivos) error
}

type generarReporteUsuariosActivosHandler struct {
	logger *zap.Logger
}

func (h *generarReporteUsuariosActivosHandler) Handle(ctx context.Context, cmd *GenerarReporteUsuariosActivos) error {
	h.logger.Info("Iniciando generación de reporte de usuarios activos")

	baseFolder := "uuid_storage"
	if _, err := os.Stat(baseFolder); os.IsNotExist(err) {
		if err := os.Mkdir(baseFolder, os.ModePerm); err != nil {
			h.logger.Error("No se pudo crear uuid_storage", zap.Error(err))
			return fmt.Errorf("no se pudo crear uuid_storage: %w", err)
		}
	}

	pdfFileName := fmt.Sprintf("%s/reporte_usuarios_activos_%d.pdf", baseFolder, time.Now().Unix())

	absPath, err := filepath.Abs(pdfFileName)
	if err != nil {
		h.logger.Error("Error obteniendo ruta absoluta del PDF", zap.Error(err))
		return fmt.Errorf("error al obtener ruta absoluta del PDF: %w", err)
	}

	err = template.ReporteUsuariosActivos(cmd.Usuarios, absPath)
	if err != nil {
		h.logger.Error("Error generando PDF de usuarios activos", zap.Error(err))
		return fmt.Errorf("error al generar el PDF: %w", err)
	}

	h.logger.Info("Reporte de usuarios activos generado correctamente", zap.String("ruta_pdf", absPath))
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

func NewGenerarReporteUsuariosActivosHandler(logger *zap.Logger) GenerarReporteUsuariosActivosHandler {
	return &generarReporteUsuariosActivosHandler{
		logger: logger,
	}
}
