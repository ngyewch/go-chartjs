package chartjs

import "encoding/json"

//

type ISpanGaps interface {
	isSpanGaps()
}

type SpanGaps bool

func (sg SpanGaps) isSpanGaps() {}

type SpanGapsNumeric float64

func (sg SpanGapsNumeric) isSpanGaps() {}

//

type Stepped string

const (
	SteppedBefore Stepped = "before"
	SteppedAfter  Stepped = "after"
	SteppedMiddle         = "middle"
)

func (s Stepped) isIStepped() {}

//

type FillTarget string

const (
	FillTargetStart  FillTarget = "start"
	FillTargetEnd    FillTarget = "end"
	FillTargetOrigin FillTarget = "origin"
	FillTargetStack  FillTarget = "stack"
	FillTargetShape  FillTarget = "shape"
)

func (v FillTarget) isIFillTarget() {}

//

type LineChartConfiguration struct {
	Data    *LineChartData              `json:"data,omitempty"`
	Options *LineControllerChartOptions `json:"options,omitempty"`
}

func (config *LineChartConfiguration) isChartConfiguration() {
}

func (config *LineChartConfiguration) MarshalJSON() ([]byte, error) {
	type Proxy LineChartConfiguration
	return json.Marshal(struct {
		Type string `json:"type"`
		Proxy
	}{
		Type:  "line",
		Proxy: Proxy(*config),
	})
}

//

type LineChartData struct {
	Datasets []*LineChartDataset `json:"datasets,omitempty"`
}

//

type LineOptions struct {
	*CommonElementOptions

	/**
	 * BorderCapStyle: Line cap style. See MDN.
	 *
	 * @default 'butt'
	 */
	BorderCapStyle CanvasLineCap `json:"borderCapStyle,omitempty"`

	/**
	 * BorderDash: Line dash. See MDN.
	 *
	 * @default []
	 */
	BorderDash []float64 `json:"borderDash,omitempty"`

	/**
	 * BorderDashOffset: Line dash offset. See MDN.
	 *
	 * @default 0.0
	 */
	BorderDashOffset *float64 `json:"borderDashOffset,omitempty"`

	/**
	 * BorderJoinStyle: Line join style. See MDN.
	 *
	 * @default 'miter'
	 */
	BorderJoinStyle CanvasLineJoin `json:"borderJoinStyle,omitempty"`

	/**
	 * CapBezierPoints: true to keep Bézier control inside the chart, false for no restriction.
	 *
	 * @default true
	 */
	CapBezierPoints *bool `json:"capBezierPoints,omitempty"`

	/**
	 * CubicInterpolationMode: Interpolation mode to apply.
	 *
	 * @default 'default'
	 */
	CubicInterpolationMode CubicInterpolationMode `json:"cubicInterpolationMode,omitempty"`

	/**
	 * Tension: Bézier curve tension (0 for no Bézier curves).
	 *
	 * @default 0
	 */
	Tension *float64 `json:"tension,omitempty"`

	/**
	 * Stepped: true to show the line as a stepped line (tension will be ignored).
	 * @default false
	 */
	Stepped IStepped `json:"stepped,omitempty"`

	/**
	 * Fill: Both line and radar charts support a fill option on the dataset object which can be used to create area between two datasets or a dataset and a boundary, i.e. the scale origin, start or end
	 */
	Fill IFillTarget `json:"fill,omitempty"`

	/**
	 * SpanGaps controls whether gaps will be spanned.
	 * If true, lines will be drawn between points with no or null data.
	 * If false, points with NaN data will create a break in the line.
	 * Can also be a number specifying the maximum gap length to span. The unit of the value depends on the scale used.
	 *
	 * @default false
	 */
	SpanGaps ISpanGaps `json:"spanGaps,omitempty"`

	// TODO Segment
}

type CubicInterpolationMode string

const (
	CubicInterpolationModeDefault  CubicInterpolationMode = "default"
	CubicInterpolationModeMonotone CubicInterpolationMode = "monotone"
)

type LineHoverOptions struct {
	*CommonHoverOptions

	HoverBorderCapStyle   CanvasLineCap  `json:"hoverBorderCapStyle,omitempty"`
	HoverBorderDash       []float64      `json:"hoverBorderDash,omitempty"`
	HoverBorderDashOffset *float64       `json:"hoverBorderDashOffset,omitempty"`
	HoverBorderJoinStyle  CanvasLineJoin `json:"hoverBorderJoinStyle,omitempty"`
}

type LineChartDataset struct {
	*ControllerDatasetOptions
	*PointPrefixedOptions
	*PointPrefixedHoverOptions
	*LineOptions
	*LineHoverOptions

	Data []Point `json:"data,omitempty"`

	/**
	 * XAxisID is the ID of the x axis to plot this dataset on.
	 */
	XAxisID string `json:"xAxisID,omitempty"`

	/**
	 * YAxisID is the ID of the y axis to plot this dataset on.
	 */
	YAxisID string `json:"yAxisID,omitempty"`

	/**
	 * SpanGaps controls whether gaps will be spanned.
	 * If true, lines will be drawn between points with no or null data.
	 * If false, points with NaN data will create a break in the line.
	 * Can also be a number specifying the maximum gap length to span. The unit of the value depends on the scale used.
	 *
	 * @default false
	 */
	SpanGaps ISpanGaps `json:"spanGaps,omitempty"`

	/**
	 * ShowLine controls whether lines between points are drawn.
	 * If false, the lines between points are not drawn.
	 *
	 * @default true
	 */
	ShowLine *bool `json:"showLine,omitempty"`
}

func (dataset *LineChartDataset) MarshalJSON() ([]byte, error) {
	type Proxy LineChartDataset
	return json.Marshal(struct {
		Type string `json:"type"`
		Proxy
	}{
		Type:  "line",
		Proxy: Proxy(*dataset),
	})

}

type LineControllerChartOptions struct {
	*CoreChartOptions

	// TODO ElementChartOptions

	Plugins *PluginOptions `json:"plugins,omitempty"`

	// TODO DatasetChartOptions

	Scales map[string]ICartesianScaleType `json:"scales,omitempty"`

	/**
	 * SpanGaps controls whether gaps will be spanned.
	 * If true, lines will be drawn between points with no or null data.
	 * If false, points with NaN data will create a break in the line.
	 * Can also be a number specifying the maximum gap length to span. The unit of the value depends on the scale used.
	 *
	 * @default false
	 */
	SpanGaps ISpanGaps `json:"spanGaps,omitempty"`

	/**
	 * ShowLine controls whether lines between points are drawn.
	 * If false, the lines between points are not drawn.
	 *
	 * @default true
	 */
	ShowLine *bool `json:"showLine,omitempty"`
}
