package chartjs

import "encoding/json"

type FontStyle string

const (
	FontStyleNormal  FontStyle = "normal"
	FontStyleItalic  FontStyle = "italic"
	FontStyleOblique FontStyle = "oblique"
	FontStyleInitial FontStyle = "initial"
	FontStyleInherit FontStyle = "inherit"
)

type FontWeight int

const (
	FontWeightNormal  FontWeight = -1
	FontWeightBold    FontWeight = -2
	FontWeightLighter FontWeight = -3
	FontWeightBolder  FontWeight = -4
)

func (fw FontWeight) MarshalJSON() ([]byte, error) {
	switch fw {
	case FontWeightNormal:
		return json.Marshal("normal")
	case FontWeightBold:
		return json.Marshal("bold")
	case FontWeightLighter:
		return json.Marshal("lighter")
	case FontWeightBolder:
		return json.Marshal("bolder")
	default:
		return json.Marshal(int(fw))
	}
}

func (fw FontWeight) Pointer() *FontWeight {
	return &fw
}

type FontSpec struct {
	Family     string      `json:"family,omitempty"`
	Size       *float64    `json:"size,omitempty"`  // default: 12
	Style      *FontStyle  `json:"style,omitempty"` // default: FontStyleNormal
	Weight     *FontWeight `json:"weight,omitempty"`
	LineHeight string      `json:"lineHeight,omitempty"` // default: "1.2"
}
