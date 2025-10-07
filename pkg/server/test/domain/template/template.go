package template

import (
	"encoding/json"
	"fmt"
	"reporteador/pkg/server/test/domain"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

func Template(data domain.Template, qrFileName string, pdfFileName string) error {
	cfg := config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		WithBottomMargin(10).
		WithKeywords("maroto", true).
		Build()

	m := maroto.New(cfg)

	addHeader(m, data.Data)
	addItemList(m, data.Data)
	addFooter(m, data.Data)

	err := m.RegisterHeader(addQR(data.Data, qrFileName))
	if err != nil {
		return fmt.Errorf("error al registrar el header: %w", err)
	}

	document, err := m.Generate()
	if err != nil {
		return fmt.Errorf("error al generar el PDF: %w", err)
	}

	err = document.Save(pdfFileName)
	if err != nil {
		return fmt.Errorf("error al guardar el PDF: %w", err)
	}

	return nil
}

func addHeader(m core.Maroto, data domain.Data) {
	m.AddRow(30,
		image.NewFromFileCol(12, "assets/png.jpg",
			props.Rect{
				Center:  false,
				Percent: 100,
			},
		),
	)
}

func addItemList(m core.Maroto, data domain.Data) {
	m.AddRow(80, text.NewCol(12, "", props.Text{
		Top:   5,
		Style: fontstyle.Bold,
		Align: align.Center,
		Size:  10,
	}))
}

func addFooter(m core.Maroto, data domain.Data) {
	m.AddRow(15,
		text.NewCol(12, "El que suscribe: "+data.Nombre+", "+data.Cargo+" "+data.Ubicacion+"", props.Text{
			Top:             5,
			Style:           fontstyle.Normal,
			Align:           align.Justify,
			VerticalPadding: 2,
			Size:            11,
		}),
	)
	m.AddRow(10,
		text.NewCol(12, "CERTIFICA", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Align: align.Center,
			Size:  12,
		}),
	)
	m.AddRow(25,
		text.NewCol(12, "Que la presente fotocopia del registro de "+data.Tramite+", a nombre de: "+data.Name+" es fiel y exacta reproducción del acervo registral que obra en: "+data.Obra+", con los siguientes datos; Folio: "+data.Num_folio+" Acta: "+data.Num_acta+" Fecha: "+data.Dd+" de "+data.Mmmm+" del "+data.Yyyy+" Levantado en: "+data.Levantada+".", props.Text{
			Top:             5,
			Style:           fontstyle.Normal,
			Align:           align.Justify,
			VerticalPadding: 2,
			Size:            12,
		}),
	)
	m.AddRow(20,
		text.NewCol(12, "A petición de la parte interesada y con fundamento en el artículo 52 del Código Civil para el Estado de Oaxaca, se extiende la presente, a los "+data.Mm+" días del mes de "+data.Mmmm+" de "+data.Yyyy+".", props.Text{
			Top:             12,
			Style:           fontstyle.Normal,
			Align:           align.Justify,
			VerticalPadding: 2,
			Size:            12,
		}),
	)
	m.AddRow(30,
		signature.NewCol(12, "FIRMA", props.Signature{
			FontFamily:  fontfamily.Courier,
			SafePadding: 1.5,
		}),
	)
	m.AddRow(10,
		text.NewCol(12, "Cotejo el encargado de la expedición de documentos:", props.Text{
			Top:   5,
			Style: fontstyle.Normal,
			Align: align.Left,
			Size:  12,
		}),
	)
}

func addQR(data domain.Data, qrFileName string) core.Row {
	return row.New(50).Add(
		col.New(6).Add(
			text.New("    • "+data.Name+"", props.Text{
				Top:   0,
				Style: fontstyle.BoldItalic,
				Size:  12,
				Align: align.Left,
			}),
			text.New("    • Causa derechos conforme al Artículo 43", props.Text{
				Top:   5,
				Style: fontstyle.BoldItalic,
				Size:  12,
				Align: align.Left,
			}),
			text.New("      de la Ley Estatal de Derechos Vigente", props.Text{
				Top:   10,
				Style: fontstyle.BoldItalic,
				Size:  12,
				Align: align.Left,
			}),
		),
		col.New(1),
		image.NewFromFileCol(6, qrFileName, props.Rect{
			Center:  true,
			Percent: 100,
			Left:    50,
		}),
	)
}

type Data struct {
	ID int `json:"id"`
}

func (c Data) ToBytes() []byte {
	b, _ := json.Marshal(c)
	return b
}
