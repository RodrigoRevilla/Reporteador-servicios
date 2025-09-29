package http

import (
	"encoding/json"
	"log"
	"reporteador/pkg/server/test/domain"
)

func (pdf TemplateRequest) ToDomain() *domain.Template {
	var dataStr string
	if pdf.Data != nil {
		jsonData, err := json.Marshal(pdf.Data)
		if err != nil {
			log.Printf("Error al serializar Data: %v", err)
			dataStr = "{}"
		} else {
			dataStr = string(jsonData)
		}
	}

	var data domain.Data
	err := json.Unmarshal([]byte(dataStr), &data)
	if err != nil {
		log.Printf("Error al deserializar Data: %v", err)
		data = domain.Data{}
	}

	resp := domain.Template{
		Data: data,
	}

	return &resp
}
