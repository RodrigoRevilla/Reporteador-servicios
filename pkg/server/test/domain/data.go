package domain

import (
	"fmt"
	"reflect"
)

type Data struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Cargo     string `json:"cargo"`
	Ubicacion string `json:"ubicacion"`
	Obra      string `json:"obra"`
	Tramite   string `json:"tramite"`
	Name      string `json:"name"`
	Num_folio string `json:"num_folio"`
	Num_acta  string `json:"num_acta"`
	Levantada string `json:"levantada"`
	Dd        string `json:"dd"`
	Mmmm      string `json:"mmmm"`
	Yyyy      string `json:"yyyy"`
	Mm        string `json:"mm"`
	Url       string `json:"url"`
}

type Template struct {
	ID         int  `json:"id"`
	Data       Data `json:"data"`
	Template   int  `json:"template"`
	Limit      int  `json:"limit"`
	Page       int  `json:"page"`
	TotalItems int  `json:"totalItems"`
	TotalPages int  `json:"totalPages"`
}

// ESTRUCTURA PARA VENTAS
type Venta struct {
	Fecha    string  `json:"fecha"`
	Producto string  `json:"producto"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

type ReporteVentas struct {
	ID     int     `json:"id"`
	Ventas []Venta `json:"ventas"`
}

// ESTRUCTURA PARA INVENTARIO
type Inventario struct {
	ID        string  `json:"id"`
	Producto  string  `json:"producto"`
	Categoria string  `json:"categoria"`
	Cantidad  int     `json:"cantidad"`
	Precio    float64 `json:"precio"`
}

type ReporteInventario struct {
	ID         int          `json:"id"`
	Inventario []Inventario `json:"inventario"`
}

// ESTRUCTURA PARA USUARIOS ACTIVOS
type UsuarioActivo struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	ActivoDesde string `json:"activo_desde"`
	Email       string `json:"email"`
	Activo      bool   `json:"activo"`
}

type ReporteUsuariosActivos struct {
	ID              int             `json:"id"`
	UsuariosActivos []UsuarioActivo `json:"usuarios"`
	Limit           int             `json:"limit"`
	Page            int             `json:"page"`
	TotalItems      int             `json:"totalItems"`
	TotalPages      int             `json:"totalPages"`
}

func Validate(d Data) error {
	val := reflect.ValueOf(d)
	typ := reflect.TypeOf(d)

	for i := 0; i < val.NumField(); i++ {
		campo := typ.Field(i).Tag.Get("json")
		valor := val.Field(i).String()

		if valor == "" {
			return fmt.Errorf("el campo '%s' no puede estar vacío", campo)
		}
	}
	return nil
}
