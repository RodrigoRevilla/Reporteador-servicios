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

func (req ReporteVentasRequest) ToDomain() *domain.ReporteVentas {
	reporte := domain.ReporteVentas{
		Ventas: make([]domain.Venta, 0),
	}

	if req.Ventas != nil {
		for _, v := range *req.Ventas {
			reporte.Ventas = append(reporte.Ventas, domain.Venta{
				Fecha:    derefString(v.Fecha),
				Producto: derefString(v.Producto),
				Cantidad: derefInt(v.Cantidad),
				Precio:   derefFloat32(v.Precio),
			})
		}
	}

	return &reporte
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func derefInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

func derefFloat32(f *float32) float64 {
	if f != nil {
		return float64(*f)
	}
	return 0.0
}
