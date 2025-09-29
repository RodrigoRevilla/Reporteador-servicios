package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reporteador/pkg/server/test/app"
	"reporteador/pkg/server/test/app/command"
	"reporteador/pkg/server/test/domain"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/render"
)

// ServerInterface represents all server handlers.
type HttpServer struct {
	app app.Application
}

type MessageResponse struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
}

// (POST /PDF)
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
		fmt.Println("Error al generar el PDF")
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

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error al enviar el PDF", http.StatusInternalServerError)
		return
	}
}

func (h HttpServer) ViewPDF(w http.ResponseWriter, r *http.Request) {
	fmt.Println("............")
	spew.Dump("")
	render.Respond(w, r, MessageResponse{Message: "PDF	 obtenido correctamente"})
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}
