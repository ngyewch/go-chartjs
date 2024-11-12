package chartjs

import (
	"encoding/json"
	"gopkg.in/go-playground/colors.v1"
	"image/color"
)

type Color struct {
	color color.Color
}

func NewColor(color color.Color) *Color {
	return &Color{
		color: color,
	}
}

func (c *Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(colors.FromStdColor(c.color).String())
}
