package legend

import (
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/icpac-igad/legend-image-generator/internal/conf"
)

type LegendConfig struct {
	Type        string       `json:"legend_type"`
	Items       []LegendItem `json:"items"`
	Transparent bool         `json:"transparent"`
	Width       int          `json:"width"`
	LabelSize   int          `json:"label_size"`
	LabelSpace  int          `json:"label_space"`
	ItemHeight  int          `json:"item_height"`
}

type LegendItem struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

func GetLegendImg(config LegendConfig) (image.Image, error) {

	itemsLen := len(config.Items)

	if itemsLen < 3 {
		return nil, fmt.Errorf("minimum required number of legend items is 3")
	}

	width := 500

	if config.Width > 100 {
		width = config.Width
	}

	height := 50
	itemHeight := 20

	if config.ItemHeight > 20 {
		itemHeight = config.ItemHeight
		height = itemHeight + 30
	}

	dc := gg.NewContext(width, height)
	dc.DrawRectangle(0, 0, float64(width), float64(height))

	if !config.Transparent {
		dc.SetColor(color.White)
		dc.Fill()
	} else {
		dc.SetColor(color.Transparent)
		dc.Fill()
	}

	if config.LabelSize > 12 {
		err := dc.LoadFontFace(conf.Configuration.Legend.FontPath, float64(config.LabelSize))
		if err != nil {
			return nil, fmt.Errorf("error loading font file: %s", conf.Configuration.Legend.FontPath)
		}
	} else {
		err := dc.LoadFontFace(conf.Configuration.Legend.FontPath, 12)
		if err != nil {
			return nil, fmt.Errorf("error loading font file: %s", conf.Configuration.Legend.FontPath)
		}
	}

	itemWidth := width / itemsLen

	xPosition := 0

	textTopPadding := 5

	if config.LabelSpace > 5 {
		textTopPadding = config.LabelSpace
	}

	textYPosition := (height - itemHeight) + textTopPadding

	strokeColor, _ := hexToColor("#bcc2be")

	for i, item := range config.Items {

		itemColor, err := hexToColor(item.Color)

		if err != nil {
			return nil, fmt.Errorf("error converting hex color to RGBA: %s", item.Color)
		}

		// if first one
		if i == 0 {
			// Draw left triangle
			dc.MoveTo(0, float64(itemHeight)/2)
			dc.LineTo(float64(xPosition)+float64(itemWidth), 0)
			dc.LineTo(float64(xPosition)+float64(itemWidth), float64(itemHeight))

			dc.MoveTo(float64(xPosition)+float64(itemWidth), float64(itemHeight))
			dc.LineTo(float64(xPosition), float64(itemHeight)/2)

			dc.SetColor(itemColor)
			dc.Fill()

			dc.MoveTo(0, float64(itemHeight)/2)
			dc.LineTo(float64(xPosition)+float64(itemWidth), 0)
			dc.LineTo(float64(xPosition)+float64(itemWidth), float64(itemHeight))

			dc.MoveTo(float64(xPosition)+float64(itemWidth), float64(itemHeight))
			dc.LineTo(float64(xPosition), float64(itemHeight)/2)

			dc.SetColor(strokeColor)
			dc.Stroke()

		} else if i == itemsLen-1 {
			// last one. Draw right triangle
			dc.MoveTo(float64(xPosition), 0)
			dc.LineTo(float64(xPosition)+float64(itemWidth), float64(itemHeight)/2)
			dc.LineTo(float64(xPosition), float64(itemHeight))

			dc.MoveTo(float64(xPosition)+float64(itemWidth), float64(itemHeight)/2)
			dc.LineTo(float64(xPosition), float64(itemHeight))

			dc.SetColor(itemColor)
			dc.Fill()

			dc.SetColor(strokeColor)
			dc.MoveTo(float64(xPosition), 0)
			dc.LineTo(float64(xPosition)+float64(itemWidth), float64(itemHeight)/2)
			dc.LineTo(float64(xPosition), float64(itemHeight))

			dc.MoveTo(float64(xPosition)+float64(itemWidth), float64(itemHeight)/2)
			dc.LineTo(float64(xPosition), float64(itemHeight))

			dc.Stroke()
		} else {
			// draw item rectangle bar
			dc.DrawRectangle(float64(xPosition), 0, float64(itemWidth), float64(itemHeight))
			dc.SetColor(itemColor)
			dc.Fill()

			dc.DrawRectangle(float64(xPosition), 0, float64(itemWidth), float64(itemHeight))
			dc.SetColor(strokeColor)
			dc.Stroke()

		}

		if i != (itemsLen - 1) {

			if item.Value != "" {
				// draw label
				pTextWidth, _ := dc.MeasureString(item.Value)
				labelMid := pTextWidth / 2

				labelXPosition := (xPosition + itemWidth) - int(labelMid)

				dc.SetColor(color.Black)
				dc.DrawString(item.Value, float64(labelXPosition), float64(textYPosition))
			}
		}

		// increment xPosition
		xPosition += itemWidth
	}

	return dc.Image(), nil
}
