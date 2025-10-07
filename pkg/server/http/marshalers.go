package http

import (
	"encoding/json"
	"log"
	"reporteador/pkg/server/test/domain"

    openapi_types "github.com/oapi-codegen/runtime/types"
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

func (req ReporteInventarioRequest) ToDomain() *domain.ReporteInventario {
	reporte := domain.ReporteInventario{
		Inventario: make([]domain.Inventario, 0),
	}

	if req.Inventario != nil {
		for _, v := range *req.Inventario {
			reporte.Inventario = append(reporte.Inventario, domain.Inventario{
				ID:        derefString(v.Id),
				Producto:  derefString(v.Producto),
				Categoria: derefString(v.Categoria),
				Cantidad:  derefInt(v.Cantidad),
				Precio:    derefFloat32(v.Precio),
			})
		}
	}

	return &reporte
}

func (req ReporteUsuariosActivosRequest) ToDomain() *domain.ReporteUsuariosActivos {
	reporte := domain.ReporteUsuariosActivos{
		UsuariosActivos: make([]domain.UsuarioActivo, 0),
	}

	if req.Usuarios != nil {
		for _, u := range *req.Usuarios {
			reporte.UsuariosActivos = append(reporte.UsuariosActivos, domain.UsuarioActivo{
				ID:          derefString(u.Id),
				Nombre:      derefString(u.Nombre),
				Email:       derefString(u.Email),
				ActivoDesde: derefDate(u.ActivoDesde),
				Activo:      derefBool(u.Activo),
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


func derefDate(d *openapi_types.Date) string {
	if d != nil {
		return d.String()
	}
	return ""
}

func derefBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

