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

func ReporteVentas(ventas []domain.Venta, pdfFileName string) error {
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
		text.NewCol(12, "Reporte de Ventas", props.Text{
			Align: align.Center,
			Style: fontstyle.Bold,
			Size:  16,
		}),
	)

	m.AddRow(10,
		col.New(3).Add(text.New("Fecha", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
		col.New(3).Add(text.New("Producto", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
		col.New(2).Add(text.New("Cantidad", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
		col.New(2).Add(text.New("Precio", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
		col.New(2).Add(text.New("Total", props.Text{Style: fontstyle.Bold, Align: align.Center, Size: 10})),
	)

	var totalGeneral float64

	for _, v := range ventas {
		total := float64(v.Cantidad) * v.Precio
		totalGeneral += total

		m.AddRow(8,
			col.New(3).Add(text.New(v.Fecha, props.Text{Align: align.Center, Size: 9})),
			col.New(3).Add(text.New(v.Producto, props.Text{Align: align.Center, Size: 9})),
			col.New(2).Add(text.New(fmt.Sprintf("%d", v.Cantidad), props.Text{Align: align.Center, Size: 9})),
			col.New(2).Add(text.New(fmt.Sprintf("$%.2f", v.Precio), props.Text{Align: align.Center, Size: 9})),
			col.New(2).Add(text.New(fmt.Sprintf("$%.2f", total), props.Text{Align: align.Center, Size: 9})),
		)
	}

	m.AddRow(15,
		text.NewCol(12, fmt.Sprintf("TOTAL GENERAL: $%.2f", totalGeneral), props.Text{
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
