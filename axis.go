package chartjs

import (
	"encoding/json"
	"fmt"
)

type Visibility int

const (
	VisibilityHidden Visibility = iota
	VisibilityVisible
	VisibilityAuto
)

func (v Visibility) MarshalJSON() ([]byte, error) {
	switch v {
	case VisibilityHidden:
		return json.Marshal(false)
	case VisibilityVisible:
		return json.Marshal(true)
	case VisibilityAuto:
		return json.Marshal("auto")
	default:
		return nil, fmt.Errorf("unknown visibility: %d", v)
	}
}

type IScaleOptions interface {
	isScaleOptions() bool // TODO
}

// CoreScale

type CoreScaleOptions struct {
	Display         *Visibility `json:"display"` // default: VisibilityVisible
	AlignToPixels   *bool       `json:"alignToPixels,omitempty"`
	BackgroundColor *Color      `json:"backgroundColor,omitempty"`
	Reverse         *bool       `json:"reverse,omitempty"` // default: false
	Clip            *bool       `json:"clip,omitempty"`    // default: true
	Weight          *float64    `json:"weight,omitempty"`
	Min             *float64    `json:"min,omitempty"`
	Max             *float64    `json:"max,omitempty"`
	SuggestedMin    *float64    `json:"suggestedMin,omitempty"`
	SuggestedMax    *float64    `json:"suggestedMax,omitempty"`
}

// CartesianScale

type CartesianScaleOptions struct {
	*CoreScaleOptions

	Bounds string `json:"bounds,omitempty"`
	// Position
	Stack       string   `json:"stack,omitempty"`
	StackWeight *float64 `json:"stackWeight,omitempty"`
	Axis        string   `json:"axis,omitempty"`
	Offset      *bool    `json:"offset,omitempty"` // default: false
	// Grid
	// Border
	Title *CartesianScaleTitle `json:"title,omitempty"`
	// Stacked
	// Ticks
}

type CartesianScaleTitle struct {
	Display *bool  `json:"display,omitempty"`
	Align   string `json:"align,omitempty"`
	Text    string `json:"text,omitempty"`
	Color   *Color `json:"color,omitempty"`
	// Font
	// Padding
}

// LinearScale

type LinearScaleOptions struct {
	*CartesianScaleOptions

	BeginAtZero bool `json:"beginAtZero,omitempty"`
	// Grace
}

func (axis *LinearScaleOptions) isScaleOptions() bool {
	return true
}

func (axis *LinearScaleOptions) MarshalJSON() ([]byte, error) {
	type Proxy LinearScaleOptions
	return json.Marshal(struct {
		Proxy
		Type string `json:"type"`
	}{
		Proxy: Proxy(*axis),
		Type:  "linear",
	})
}

// TimeScale

type TimeScaleOptions struct {
	*CartesianScaleOptions

	OffsetAfterAutoskip *bool `json:"offsetAfterAutoskip,omitempty"` // default: false
	// Adapters
	// Time
	// Ticks
}

func (axis *TimeScaleOptions) isScaleOptions() bool {
	return true
}

func (axis *TimeScaleOptions) MarshalJSON() ([]byte, error) {
	type Proxy TimeScaleOptions
	return json.Marshal(struct {
		Proxy
		Type string `json:"type"`
	}{
		Proxy: Proxy(*axis),
		Type:  "time",
	})
}

//

type AxisPosition struct {
	Position string
	X        float64
	Y        float64
}

func (ap *AxisPosition) MarshalJSON() ([]byte, error) {
	if ap.Position != "" {
		return json.Marshal(ap.Position)
	} else if ap.X != 0 {
		return json.Marshal(struct {
			X float64 `json:"x"`
		}{
			X: ap.X,
		})
	} else if ap.Y != 0 {
		return json.Marshal(struct {
			Y float64 `json:"y"`
		}{
			Y: ap.Y,
		})
	} else {
		return json.Marshal(nil)
	}
}
