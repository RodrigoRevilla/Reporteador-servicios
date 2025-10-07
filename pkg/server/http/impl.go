package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reporteador/pkg/server/test/app"
	"reporteador/pkg/server/test/app/command"
	"reporteador/pkg/server/test/domain"
	"reporteador/pkg/server/test/domain/template"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/render"
)

type HttpServer struct {
	app app.Application
}

type MessageResponse struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
}

// (POST /generate)
func (h HttpServer) GeneratePDF(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("Cuerpo recibido:", string(body))

	pdf := &TemplateRequest{}
	if err := json.Unmarshal(body, pdf); err != nil {
		http.Error(w, "Error al deserializar JSON", http.StatusBadRequest)
		return
	}

	if err := domain.Validate(pdf.ToDomain().Data); err != nil {
		http.Error(w, fmt.Sprintf("Validación fallida: %v", err), http.StatusBadRequest)
		return
	}

	GenerarPDFCmd := command.GenerarPDF{Data: *pdf.ToDomain()}
	err = h.app.Commands.GenerarPDF.Handle(r.Context(), &GenerarPDFCmd)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al generar el PDF: %v", err), http.StatusInternalServerError)
		return
	}

	if GenerarPDFCmd.PDFPath == "" {
		http.Error(w, "No se generó ningún PDF", http.StatusInternalServerError)
		return
	}

	download := r.URL.Query().Get("download")

	w.Header().Set("Content-Type", "application/pdf")
	if download == "true" {
		w.Header().Set("Content-Disposition", "attachment; filename=documento.pdf")
	} else {
		w.Header().Set("Content-Disposition", "inline; filename=documento.pdf")
	}

	file, err := os.Open(GenerarPDFCmd.PDFPath)
	if err != nil {
		http.Error(w, "Error al abrir el PDF", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, _ = io.Copy(w, file)
}

// (POST /ventas)
func (h HttpServer) GenerateReporteVentas(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var reporte domain.ReporteVentas
	if err := json.Unmarshal(body, &reporte); err != nil {
		http.Error(w, "Error al deserializar JSON", http.StatusBadRequest)
		return
	}

	if len(reporte.Ventas) == 0 {
		http.Error(w, "No se recibieron ventas para generar el reporte", http.StatusBadRequest)
		return
	}

	pdfFile := fmt.Sprintf("uuid_storage/reporte_ventas_%d.pdf", time.Now().Unix())

	log.Printf("Iniciando generación de PDF de ventas")

	err = template.ReporteVentas(reporte.Ventas, pdfFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al generar el PDF: %v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("PDF de ventas generado correctamente en: %s", pdfFile)

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=reporte_ventas.pdf")

	file, err := os.Open(pdfFile)
	if err != nil {
		http.Error(w, "Error al abrir el PDF", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error al enviar el PDF", http.StatusInternalServerError)
		return
	}
}

// (POST /inventario)
func (h HttpServer) GenerateReporteInventario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var reporte domain.ReporteInventario
	if err := json.Unmarshal(body, &reporte); err != nil {
		http.Error(w, "Error al deserializar JSON", http.StatusBadRequest)
		return
	}

	if len(reporte.Inventario) == 0 {
		http.Error(w, "No se recibieron items de inventario para generar el reporte", http.StatusBadRequest)
		return
	}

	pdfFile := fmt.Sprintf("uuid_storage/reporte_inventario_%d.pdf", time.Now().Unix())

	log.Printf("Iniciando generación de PDF de inventario")

	err = template.ReporteInventario(reporte.Inventario, pdfFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al generar el PDF: %v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("PDF de inventario generado correctamente en: %s", pdfFile)

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=reporte_inventario.pdf")

	file, err := os.Open(pdfFile)
	if err != nil {
		http.Error(w, "Error al abrir el PDF", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error al enviar el PDF", http.StatusInternalServerError)
		return
	}
}

// (POST /usuarios-activos)
func (h HttpServer) GenerateReporteUsuariosActivos(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	req := ReporteUsuariosActivosRequest{}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Error al deserializar JSON (request): "+err.Error(), http.StatusBadRequest)
		return
	}

	reporteDomain := req.ToDomain()

	activos := make([]domain.UsuarioActivo, 0, len(reporteDomain.UsuariosActivos))
	for _, u := range reporteDomain.UsuariosActivos {
		if u.Activo {
			activos = append(activos, u)
		}
	}

	if len(activos) == 0 {
		http.Error(w, "No se recibieron usuarios activos para generar el reporte", http.StatusBadRequest)
		return
	}

	pdfFile := fmt.Sprintf("uuid_storage/reporte_usuarios_activos_%d.pdf", time.Now().Unix())

	log.Printf("Iniciando generación de PDF de usuarios activos")
	if err := template.ReporteUsuariosActivos(activos, pdfFile); err != nil {
		http.Error(w, fmt.Sprintf("Error al generar el PDF: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("PDF de usuarios activos generado correctamente en: %s", pdfFile)

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=reporte_usuarios_activos.pdf")

	file, err := os.Open(pdfFile)
	if err != nil {
		http.Error(w, "Error al abrir el PDF", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error al enviar el PDF", http.StatusInternalServerError)
		return
	}
}

// (GET /view)
func (h HttpServer) ViewPDF(w http.ResponseWriter, r *http.Request) {
	spew.Dump("ViewPDF ejecutado")
	render.Respond(w, r, MessageResponse{Message: "PDF obtenido correctamente"})
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}
