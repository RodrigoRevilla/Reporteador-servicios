package domain

import (
	"fmt"
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

func Validate(v Data) error {
	if v.Nombre == "" {
		return fmt.Errorf("el campo 'nombre' no puede estar vacío")
	}
	if v.Cargo == "" {
		return fmt.Errorf("el campo 'cargo' no puede estar vacío")
	}
	if v.Ubicacion == "" {
		return fmt.Errorf("el campo 'ubicacion' no puede estar vacío")
	}
	if v.Obra == "" {
		return fmt.Errorf("el campo 'obra' no puede estar vacío")
	}
	if v.Tramite == "" {
		return fmt.Errorf("el campo 'tramite' no puede estar vacío")
	}
	if v.Name == "" {
		return fmt.Errorf("el campo 'name' no puede estar vacío")
	}
	if v.Num_folio == "" {
		return fmt.Errorf("el campo 'num_folio' no puede estar vacío")
	}
	if v.Num_acta == "" {
		return fmt.Errorf("el campo 'num_acta' no puede estar vacío")
	}
	if v.Levantada == "" {
		return fmt.Errorf("el campo 'levantada' no puede estar vacío")
	}
	if v.Dd == "" {
		return fmt.Errorf("el campo 'dd' no puede estar vacío")
	}
	if v.Mmmm == "" {
		return fmt.Errorf("el campo 'mmmm' no puede estar vacío")
	}
	if v.Yyyy == "" {
		return fmt.Errorf("el campo 'yyyy' no puede estar vacío")
	}
	if v.Url == "" {
		return fmt.Errorf("el campo 'url' no puede estar vacío")
	}
	return nil
}
