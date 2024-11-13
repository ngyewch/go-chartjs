package chartjs

import (
	"encoding/json"
	"fmt"
)

// Visibility

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

func (v Visibility) Pointer() *Visibility {
	return &v
}

// GridLineOptions

type GridLineOptions struct {
	Display              *bool     `json:"display,omitempty"`              // default: true
	Circular             *bool     `json:"circular,omitempty"`             // default: false
	Color                string    `json:"color,omitempty"`                // default: 'rgba(0, 0, 0, 0.1)'
	LineWidth            *float64  `json:"lineWidth,omitempty"`            // default: 1
	DrawOnChartArea      *bool     `json:"drawOnChartArea,omitempty"`      // default: true
	DrawTicks            *bool     `json:"drawTicks,omitempty"`            // default: true
	TickBorderDash       []float64 `json:"tickBorderDash,omitempty"`       // default: []
	TickBorderDashOffset *float64  `json:"tickBorderDashOffset,omitempty"` // default: 0
	TickColor            string    `json:"tickColor,omitempty"`            // default: 'rgba(0, 0, 0, 0.1)'
	TickLength           *float64  `json:"tickLength,omitempty"`           // default: 10
	TickWidth            *float64  `json:"tickWidth,omitempty"`            // default: 1
	Offset               *bool     `json:"offset,omitempty"`               // default: false
	Z                    *float64  `json:"z,omitempty"`                    // default: 0
}

// BorderOptions

type BorderOptions struct {
	Display    *bool     `json:"display,omitempty"`    // default: true
	Dash       []float64 `json:"dash,omitempty"`       // default: []
	DashOffset *float64  `json:"dashOffset,omitempty"` // default: 0
	Color      string    `json:"color,omitempty"`
	Width      *float64  `json:"width,omitempty"`
	Z          *float64  `json:"z,omitempty"`
}

// IPadding

type IPadding interface {
	isPadding() bool
}

// SimplePadding

type SimplePadding float64

func (p SimplePadding) isPadding() bool {
	return true
}

// Padding

type Padding ChartArea

func (p Padding) isPadding() bool {
	return true
}

// TickOptions

type TickOptionsMajor struct {
	Enabled *bool `json:"enabled,omitempty"` // default: false
}

type TickOptions struct {
	BackdropColor     string            `json:"backdropColor,omitempty"`   // default: 'rgba(255, 255, 255, 0.75)'
	BackdropPadding   IPadding          `json:"backdropPadding,omitempty"` // default: 2
	Display           *bool             `json:"display,omitempty"`         // default: true
	Color             string            `json:"color,omitempty"`
	Font              *FontSpec         `json:"font,omitempty"`
	Padding           *float64          `json:"padding,omitempty"`
	ShowLabelBackdrop *bool             `json:"showLabelBackdrop,omitempty"` // default: false
	TextStrokeColor   string            `json:"textStrokeColor,omitempty"`
	TextStrokeWidth   *float64          `json:"textStrokeWidth,omitempty"` // default: 0
	Z                 *float64          `json:"z,omitempty"`               // default: 0
	Major             *TickOptionsMajor `json:"major,omitempty"`
}

// IScaleOptions

type IScaleOptions interface {
	isScaleOptions() bool // TODO
}

// CoreScale

type CoreScaleOptions struct {
	Display         *Visibility `json:"display"` // default: VisibilityVisible
	AlignToPixels   *bool       `json:"alignToPixels,omitempty"`
	BackgroundColor string      `json:"backgroundColor,omitempty"`
	Reverse         *bool       `json:"reverse,omitempty"` // default: false
	Clip            *bool       `json:"clip,omitempty"`    // default: true
	Weight          *float64    `json:"weight,omitempty"`
	Min             *float64    `json:"min,omitempty"`
	Max             *float64    `json:"max,omitempty"`
	SuggestedMin    *float64    `json:"suggestedMin,omitempty"`
	SuggestedMax    *float64    `json:"suggestedMax,omitempty"`
}

// CartesianTickOptions

type CartesianTickOptions struct {
	*TickOptions

	SampleSize      int      `json:"sampleSize,omitempty"`      // default: ticks.length
	Align           *Align   `json:"align,omitempty"`           // 'inner' not supported. default: 'center'
	AutoSkip        *bool    `json:"autoSkip,omitempty"`        // default: true
	AutoSkipPadding *float64 `json:"autoSkipPadding,omitempty"` // default: 0
	CrossAlign      string   `json:"crossAlign,omitempty"`      // 'near' | 'center' | 'far'. default: 'near'
	IncludeBounds   *bool    `json:"includeBounds,omitempty"`   //  default: true
	LabelOffset     *float64 `json:"labelOffset,omitempty"`     // default: 0
	MinRotation     *float64 `json:"minRotation,omitempty"`     // default: 0
	MaxRotation     *float64 `json:"maxRotation,omitempty"`     // default: 50
	Mirror          *bool    `json:"mirror,omitempty"`          // default: false
	Padding         *float64 `json:"padding,omitempty"`         // default: 0
	MaxTicksLimit   *int     `json:"maxTicksLimit,omitempty"`   // default: 11
}

// CartesianScale

type CartesianScaleOptions struct {
	*CoreScaleOptions

	Bounds      string                 `json:"bounds,omitempty"` // 'ticks' | 'data', default: 'ticks'
	Position    ICartesianAxisPosition `json:"position,omitempty"`
	Stack       string                 `json:"stack,omitempty"`
	StackWeight *float64               `json:"stackWeight,omitempty"`
	Axis        string                 `json:"axis,omitempty"`   // 'x' | 'y' | 'r'
	Offset      *bool                  `json:"offset,omitempty"` // default: false
	Grid        *GridLineOptions       `json:"grid,omitempty"`
	Border      *BorderOptions         `json:"border,omitempty"`
	Title       *CartesianScaleTitle   `json:"title,omitempty"`
	Stacked     *bool                  `json:"stacked,omitempty"` // 'single' not supported. default: false
	Ticks       *CartesianTickOptions  `json:"ticks,omitempty"`
}

type ICartesianAxisPosition interface {
	isCartesianAxisPosition() bool
}

type CartesianAxisPositionSimple string

const (
	CartesianAxisPositionLeft   CartesianAxisPositionSimple = "left"
	CartesianAxisPositionTop    CartesianAxisPositionSimple = "top"
	CartesianAxisPositionRight  CartesianAxisPositionSimple = "right"
	CartesianAxisPositionBottom CartesianAxisPositionSimple = "bottom"
	CartesianAxisPositionCenter CartesianAxisPositionSimple = "center"
)

func (p CartesianAxisPositionSimple) isCartesianAxisPosition() bool {
	return true
}

type CartesianAxisPositionRelative map[string]float64

func (p CartesianAxisPositionRelative) isCartesianAxisPosition() bool {
	return true
}

type CartesianScaleTitle struct {
	Display *bool                       `json:"display,omitempty"`
	Align   *Align                      `json:"align,omitempty"`
	Text    string                      `json:"text,omitempty"`
	Color   string                      `json:"color,omitempty"`
	Font    *FontSpec                   `json:"font,omitempty"`
	Padding *CartesianScaleTitlePadding `json:"padding,omitempty"`
}

type CartesianScaleTitlePadding struct {
	Top    *float64 `json:"top,omitempty"`
	Bottom *float64 `json:"bottom,omitempty"`
	Y      *float64 `json:"y,omitempty"`
}

// LinearScale

type LinearTickOptions struct {
	*CartesianTickOptions

	Format    map[string]any `json:"format,omitempty"`
	Precision *int           `json:"precision,omitempty"`
	StepSize  *float64       `json:"stepSize,omitempty"`
	Count     *int           `json:"count,omitempty"`
}

type LinearScaleOptions struct {
	*CartesianScaleOptions

	BeginAtZero *bool `json:"beginAtZero,omitempty"`
	// Grace
	Ticks *LinearTickOptions `json:"ticks,omitempty"`
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

type TimeScaleTickOptions struct {
	*CartesianTickOptions

	Source   string   `json:"source,omitempty"`   // 'labels' | 'auto' | 'data'. default: 'auto'
	StepSize *float64 `json:"stepSize,omitempty"` // default: 1
}

type TimeScaleOptions struct {
	*CartesianScaleOptions

	OffsetAfterAutoskip *bool `json:"offsetAfterAutoskip,omitempty"` // default: false
	// Adapters
	// Time
	Ticks *TimeScaleTickOptions `json:"ticks,omitempty"`
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
