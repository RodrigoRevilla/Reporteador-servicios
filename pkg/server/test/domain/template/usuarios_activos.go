package template

import (
	"fmt"

	"reporteador/pkg/server/test/domain"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type UsuarioActivo struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Email       string `json:"email"`
	Activo      bool   `json:"activo"`
	ActivoDesde string `json:"activo_desde"`
}

func ReporteUsuariosActivos(usuarios []domain.UsuarioActivo, pdfFileName string) error {
	cfg := config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		WithBottomMargin(10).
		Build()

	m := maroto.New(cfg)

	m.AddRow(20,
		text.NewCol(12, "Reporte de Usuarios Activos", props.Text{
			Align: align.Center,
			Style: fontstyle.Bold,
			Size:  16,
		}),
	)

	m.AddRow(10,
		col.New(2).Add(text.New("ID", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
		col.New(2).Add(text.New("Nombre", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
		col.New(3).Add(text.New("Email", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
		col.New(3).Add(text.New("Activo", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
		col.New(2).Add(text.New("Activo Desde", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
	)

	for _, user := range usuarios {
		estado := "No"

		if user.Activo {
			estado = "Sí"
		}

		m.AddRow(8,
			col.New(2).Add(text.New(user.ID, props.Text{Align: align.Center, Size: 9})),
			col.New(2).Add(text.New(user.Nombre, props.Text{Align: align.Center, Size: 9})),
			col.New(3).Add(text.New(user.Email, props.Text{Align: align.Center, Size: 9})),
			col.New(3).Add(text.New(estado, props.Text{Align: align.Center, Size: 9})),
			col.New(2).Add(text.New(user.ActivoDesde, props.Text{Align: align.Center, Size: 9})),
		)
	}

	m.AddRow(15,
		text.NewCol(12, fmt.Sprintf("Total de usuarios activos: %d", len(usuarios)), props.Text{
			Align: align.Right,
			Style: fontstyle.Bold,
			Size:  12,
		}),
	)

	doc, err := m.Generate()
	if err != nil {
		return fmt.Errorf("error al generar el PDF: %w", err)
	}

	err = doc.Save(pdfFileName)
	if err != nil {
		return fmt.Errorf("error al guardar el PDF: %w", err)
	}

	return nil
}
