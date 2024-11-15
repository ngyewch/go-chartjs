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
	/**
	 * Display
	 *
	 * @default true
	 */
	Display *bool `json:"display,omitempty"`

	/**
	 * Circular
	 *
	 * @default false
	 */
	Circular *bool `json:"circular,omitempty"`

	/**
	 * Color
	 *
	 * @default 'rgba(0, 0, 0, 0.1)'
	 */
	Color string `json:"color,omitempty"`

	/**
	 * LineWidth
	 *
	 * @default 1
	 */
	LineWidth *float64 `json:"lineWidth,omitempty"`

	/**
	 * DrawOnChartArea
	 *
	 * @default true
	 */
	DrawOnChartArea *bool `json:"drawOnChartArea,omitempty"`

	/**
	 * DrawTicks
	 *
	 * @default true
	 */
	DrawTicks *bool `json:"drawTicks,omitempty"`

	/**
	 * TickBorderDash
	 *
	 * @default []
	 */
	TickBorderDash []float64 `json:"tickBorderDash,omitempty"`

	/**
	 * TickBorderDashOffset
	 *
	 * @default 0
	 */
	TickBorderDashOffset *float64 `json:"tickBorderDashOffset,omitempty"`

	/**
	 * TickColor
	 *
	 * @default 'rgba(0, 0, 0, 0.1)'
	 */
	TickColor string `json:"tickColor,omitempty"`

	/**
	 * TickLength
	 *
	 * @default 10
	 */
	TickLength *float64 `json:"tickLength,omitempty"`

	/**
	 * TickWidth
	 *
	 * @default 1
	 */
	TickWidth *float64 `json:"tickWidth,omitempty"`

	/**
	 * Offset
	 *
	 * @default false
	 */
	Offset *bool `json:"offset,omitempty"`

	/**
	 * Z
	 *
	 * @default 0
	 */
	Z *float64 `json:"z,omitempty"`
}

// BorderOptions

type BorderOptions struct {
	/**
	 * Display
	 *
	 * @default true
	 */
	Display *bool `json:"display,omitempty"`

	/**
	 * Dash
	 *
	 * @default []
	 */
	Dash []float64 `json:"dash,omitempty"`

	/**
	 * DashOffset
	 *
	 * @default 0
	 */
	DashOffset *float64 `json:"dashOffset,omitempty"`

	Color string `json:"color,omitempty"`

	Width *float64 `json:"width,omitempty"`

	Z *float64 `json:"z,omitempty"`
}

// TickOptions

type TickOptionsMajor struct {
	/**
	 * Enabled: If true, major ticks are generated. A major tick will affect autoskipping and major will be defined on ticks in the scriptable options context.
	 *
	 * @default false
	 */
	Enabled *bool `json:"enabled,omitempty"` // default: false
}

type TickOptions struct {
	/**
	 * BackdropColor: Color of label backdrops.
	 *
	 * @default 'rgba(255, 255, 255, 0.75)'
	 */
	BackdropColor string `json:"backdropColor,omitempty"`

	/**
	 * BackdropPadding: Padding of tick backdrop.
	 *
	 * @default 2
	 */
	BackdropPadding IPadding `json:"backdropPadding,omitempty"`

	/**
	 * Display: If true, show tick labels.
	 *
	 * @default true
	 */
	Display *bool `json:"display,omitempty"`

	/**
	 * Color of tick
	 *
	 * @see Defaults.color
	 */
	Color string `json:"color,omitempty"`

	/**
	 * Font: see Fonts
	 */
	Font *FontSpec `json:"font,omitempty"`

	/**
	 * Padding: Sets the offset of the tick labels from the axis
	 */
	Padding *float64 `json:"padding,omitempty"`

	/**
	 * ShowLabelBackdrop: If true, draw a background behind the tick labels.
	 *
	 * @default false
	 */
	ShowLabelBackdrop *bool `json:"showLabelBackdrop,omitempty"`

	/**
	 * TextStrokeColor: The color of the stroke around the text.
	 *
	 * @default undefined
	 */
	TextStrokeColor string `json:"textStrokeColor,omitempty"`

	/**
	 * TextStrokeWidth: Stroke width around the text.
	 *
	 * @default 0
	 */
	TextStrokeWidth *float64 `json:"textStrokeWidth,omitempty"`

	/**
	 * Z: z-index of tick layer.
	 * Useful when ticks are drawn on chart area.
	 * Values <= 0 are drawn under datasets, > 0 on top.
	 *
	 * @default 0
	 */
	Z *float64 `json:"z,omitempty"`

	Major *TickOptionsMajor `json:"major,omitempty"`
}

// IScaleOptions

type IScaleOptions interface {
	isScaleOptions()
}

// ICartesianScaleType

type ICartesianScaleType interface {
	isCartesianScaleType()
}

// IRadialScaleType

type IRadialScaleType interface {
	isRadialScaleType()
}

// CoreScale

type CoreScaleOptions struct {
	/**
	 * Display: Controls the axis global visibility (visible when true, hidden when false).
	 * When display: 'auto', the axis is visible only if at least one associated dataset is visible.
	 *
	 * @default true
	 */
	Display *Visibility `json:"display"`

	/**
	 * AlignToPixels: Align pixel values to device pixels
	 */
	AlignToPixels *bool `json:"alignToPixels,omitempty"`

	/**
	 * BackgroundColor: Background color of the scale area.
	 */
	BackgroundColor string `json:"backgroundColor,omitempty"`

	/**
	 * Reverse the scale.
	 *
	 * @default false
	 */
	Reverse *bool `json:"reverse,omitempty"`

	/**
	 * Clip the dataset drawing against the size of the scale instead of chart area.
	 *
	 * @default true
	 */
	Clip *bool `json:"clip,omitempty"`

	/**
	 * Weight: The weight used to sort the axis.
	 * Higher weights are further away from the chart area.
	 *
	 * @default true
	 */
	Weight *float64 `json:"weight,omitempty"`

	/**
	 * Min: User defined minimum value for the scale, overrides minimum value from data.
	 */
	Min *float64 `json:"min,omitempty"`

	/**
	 * Max: User defined maximum value for the scale, overrides maximum value from data.
	 */
	Max *float64 `json:"max,omitempty"`

	/**
	 * SuggestedMin: Adjustment used when calculating the maximum data value.
	 */
	SuggestedMin *float64 `json:"suggestedMin,omitempty"`

	/**
	 * SuggestedMax: Adjustment used when calculating the minimum data value.
	 */
	SuggestedMax *float64 `json:"suggestedMax,omitempty"`
}
