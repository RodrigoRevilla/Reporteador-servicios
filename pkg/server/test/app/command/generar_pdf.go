package command

import (
	"context"
	"fmt"
	"path/filepath"
	"reporteador/folder"
	"reporteador/pkg/server/test/domain"
	"reporteador/pkg/server/test/domain/template"

	"go.uber.org/zap"
)

type GenerarPDF struct {
	Data    domain.Template
	PDFPath string
}

type GenerarPDFHandler interface {
	Handle(ctx context.Context, cmd *GenerarPDF) error
}

type generarPDFHandler struct {
	logger *zap.Logger
}

func (h *generarPDFHandler) Handle(ctx context.Context, cmd *GenerarPDF) error {
	h.logger.Info("Iniciando generación de PDF", zap.String("url", cmd.Data.Data.Url))

	qrFileName, pdfFileName, err := folder.GenerateQRCode(cmd.Data.Data.Url)
	if err != nil {
		h.logger.Error("Error generando QR", zap.Error(err))
		return fmt.Errorf("error al generar el código QR: %w", err)
	}

	absQRFileName, err := filepath.Abs(qrFileName)
	if err != nil {
		h.logger.Error("Error obteniendo ruta absoluta del QR", zap.Error(err))
		return fmt.Errorf("error al obtener ruta absoluta del QR: %w", err)
	}

	h.logger.Info("QR generado correctamente", zap.String("ruta_qr", absQRFileName))

	err = template.Template(cmd.Data, absQRFileName, pdfFileName)
	if err != nil {
		h.logger.Error("Error generando PDF", zap.Error(err))
		return fmt.Errorf("error al generar el PDF: %w", err)
	}

	h.logger.Info("PDF generado correctamente", zap.String("ruta_pdf", pdfFileName))
	cmd.PDFPath = pdfFileName
	return nil
}

func NewGenerarPDFHandler(logger *zap.Logger) GenerarPDFHandler {
	return &generarPDFHandler{
		logger: logger,
	}
}
