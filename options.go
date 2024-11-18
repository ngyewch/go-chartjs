package chartjs

import "encoding/json"

type CoreChartOptions struct {
	// TODO datasets

	/**
	 * IndexAxis is the base axis of the chart.
	 * 'x' for vertical charts and 'y' for horizontal charts.
	 *
	 * @default 'x'
	 */
	IndexAxis string `json:"indexAxis,omitempty"`

	/**
	 * Clip controls how to clip relative to chartArea.
	 * Positive value allows overflow, negative value clips that many pixels inside chartArea. 0 = clip at chartArea.
	 * Clipping can also be configured per side: `clip: {left: 5, top: false, right: -2, bottom: 0}`
	 */
	Clip IClip `json:"clip,omitempty"`

	/**
	 * Color is the base color.
	 */
	Color string `json:"color,omitempty"`

	/**
	 * BackgroundColor is the base background color.
	 */
	BackgroundColor string `json:"backgroundColor,omitempty"`

	/**
	 * HoverBackgroundColor is the base hover background color.
	 */
	HoverBackgroundColor string `json:"hoverBackgroundColor,omitempty"`

	/**
	 * BorderColor is the base border color.
	 */
	BorderColor string `json:"borderColor,omitempty"`

	/**
	 * HoverBorderColor is the base hover border color.
	 */
	HoverBorderColor string `json:"hoverBorderColor,omitempty"`

	/**
	 * Font is the base font.
	 */
	Font *FontSpec `json:"font,omitempty"`

	/**
	 * Responsive controls responsive resize.
	 * If true, resizes the chart canvas when its container does (important note...).
	 *
	 * @default true
	 */
	Responsive *bool `json:"responsive,omitempty"`

	/**
	 * MaintainAspectRatio, if true maintain the original canvas aspect ratio (width / height) when resizing.
	 * For this option to work properly the chart must be in its own dedicated container.
	 *
	 * @default true
	 */
	MaintainAspectRatio *bool `json:"maintainAspectRatio,omitempty"`

	/**
	 * ResizeDelay, delay the resize update by give amount of milliseconds.
	 * This can ease the resize process by debouncing update of the elements.
	 *
	 * @default 0
	 */
	ResizeDelay *int `json:"resizeDelay,omitempty"`

	/**
	 * AspectRatio, canvas aspect ratio (i.e. width / height, a value of 1 representing a square canvas).
	 * Note that this option is ignored if the height is explicitly defined either as attribute or via the style.
	 *
	 * @default 2
	 */
	AspectRatio *float64 `json:"aspectRatio,omitempty"`

	/**
	 * Locale used for number formatting (using `Intl.NumberFormat`).
	 */
	Locale string `json:"locale,omitempty"`

	/**
	 * DevicePixelRatio overrides the window's default devicePixelRatio.
	 */
	DevicePixelRatio *float64 `json:"devicePixelRatio,omitempty"`

	Interaction *CoreInteractionOptions `json:"interaction,omitempty"`

	Hover *CoreInteractionOptions `json:"hover,omitempty"`

	Layout *CoreChartLayout `json:"layout,omitempty"`
}

type CoreChartLayout struct {
	AutoPadding *bool    `json:"autoPadding,omitempty"`
	Padding     IPadding `json:"padding,omitempty"`
}

type IClip interface {
	isClip()
}

type Clip float64

func (c Clip) isClip() {}

type ClipChartArea ChartArea

func (c ClipChartArea) isClip() {}

type ControllerDatasetOptions struct {
	/**
	 * The base axis of the chart.
	 * 'x' for vertical charts and 'y' for horizontal charts.
	 *
	 * @default 'x'
	 */
	IndexAxis string `json:"indexAxis,omitempty"`

	/**
	 * Clip controls how to clip relative to chartArea.
	 * Positive value allows overflow, negative value clips that many pixels inside chartArea. 0 = clip at chartArea.
	 * Clipping can also be configured per side: `clip: {left: 5, top: false, right: -2, bottom: 0}`
	 */
	Clip IClip `json:"clip,omitempty"`

	/**
	 * Label is the label for the dataset which appears in the legend and tooltips.
	 */
	Label string `json:"label,omitempty"`

	/**
	 * Order is the drawing order of dataset.
	 * Also affects order for stacking, tooltip and legend.
	 */
	Order *float64 `json:"order,omitempty"`

	/**
	 * Stack is the ID of the group to which this dataset belongs to (when stacked, each group will be a separate stack).
	 */
	Stack string `json:"stack,omitempty"`

	/**
	 * Hidden configures the visibility state of the dataset.
	 * Set it to true, to hide the dataset from the chart.
	 *
	 * @default false
	 */
	Hidden *bool `json:"hidden,omitempty"`
}

type CoreInteractionOptions struct {
	/**
	 * Mode sets which elements appear in the tooltip. See Interaction Modes for details.
	 *
	 * @default 'nearest'
	 */
	Mode InteractionMode `json:"mode,omitempty"`

	/**
	 * Intersect if true, the hover mode only applies when the mouse position intersects an item on the chart.
	 *
	 * @default true
	 */
	Intersect *bool `json:"intersect,omitempty"`

	/**
	 * Axis defines which directions are used in calculating distances.
	 * Defaults to 'x' for 'index' mode and 'xy' in dataset and 'nearest' modes.
	 */
	Axis InteractionAxis `json:"axis,omitempty"`

	/**
	 * IncludeInvisible if true, the invisible points that are outside of the chart area will also be included when evaluating interactions.
	 *
	 * @default false
	 */
	IncludeInvisible *bool `json:"includeInvisible,omitempty"`
}

type InteractionMode string

const (
	InteractionModeIndex   InteractionMode = "index"
	InteractionModeDataset InteractionMode = "dataset"
	InteractionModePoint   InteractionMode = "point"
	InteractionModeNearest InteractionMode = "nearest"
	InteractionModeX       InteractionMode = "x"
	InteractionModeY       InteractionMode = "y"
)

type InteractionAxis string

const (
	InteractionAxisX  InteractionAxis = "x"
	InteractionAxisY  InteractionAxis = "y"
	InteractionAxisXY InteractionAxis = "xy"
	InteractionAxisR  InteractionAxis = "r"
)

type CommonElementOptions struct {
	BorderWidth     *float64 `json:"borderWidth,omitempty"`
	BorderColor     string   `json:"borderColor,omitempty"`
	BackgroundColor string   `json:"backgroundColor,omitempty"`
}

type CommonHoverOptions struct {
	HoverBorderWidth     *float64 `json:"hoverBorderWidth,omitempty"`
	HoverBorderColor     string   `json:"hoverBorderColor,omitempty"`
	HoverBackgroundColor string   `json:"hoverBackgroundColor,omitempty"`
}

type CanvasLineCap string

const (
	CanvasLineCapRound  CanvasLineCap = "round"
	CanvasLineCapButt   CanvasLineCap = "butt"
	CanvasLineCapSquare CanvasLineCap = "square"
)

type CanvasLineJoin string

const (
	CanvasLineJoinBevel CanvasLineJoin = "bevel"
	CanvasLineJoinMiter CanvasLineJoin = "miter"
	CanvasLineJoinRound CanvasLineJoin = "round"
)

type PointPrefixedOptions struct {
	/**
	 * PointBackgroundColor: The fill color for points.
	 */
	PointBackgroundColor string `json:"pointBackgroundColor,omitempty"`

	/**
	 * PointBorderColor: The border color for points.
	 */
	PointBorderColor string `json:"pointBorderColor,omitempty"`

	/**
	 * PointBorderWidth: The width of the point border in pixels.
	 */
	PointBorderWidth *float64 `json:"pointBorderWidth,omitempty"`

	/**
	 * PointHitRadius: The pixel size of the non-displayed point that reacts to mouse events.
	 */
	PointHitRadius *float64 `json:"pointHitRadius,omitempty"`

	/**
	 * PointRadius: The radius of the point shape. If set to 0, the point is not rendered.
	 */
	PointRadius *float64 `json:"pointRadius,omitempty"`

	/**
	 * PointRotation: The rotation of the point in degrees.
	 */
	PointRotation *float64 `json:"pointRotation,omitempty"`

	/**
	 * PointStyle: Style of the point.
	 */
	PointStyle PointStyle `json:"pointStyle,omitempty"`
}

type PointStyle string

const (
	PointStyleCircle      PointStyle = "circle"
	PointStyleCross       PointStyle = "cross"
	PointStyleCrossRot    PointStyle = "crossRot"
	PointStyleDash        PointStyle = "dash"
	PointStyleLine        PointStyle = "line"
	PointStyleRect        PointStyle = "rect"
	PointStyleRectRounded PointStyle = "rectRounded"
	PointStyleRectRot     PointStyle = "rectRot"
	PointStyleStar        PointStyle = "star"
	PointStyleTriangle    PointStyle = "triangle"
	PointStyleFalse       PointStyle = "false"
)

func (ps PointStyle) MarshalJSON() ([]byte, error) {
	if ps == PointStyleFalse {
		return json.Marshal(false)
	} else {
		return json.Marshal(string(ps))
	}
}

type PointPrefixedHoverOptions struct {
	/**
	 * PointHoverBackgroundColor: Point background color when hovered.
	 */
	PointHoverBackgroundColor string `json:"pointHoverBackgroundColor,omitempty"`

	/**
	 * PointHoverBorderColor: Point border color when hovered.
	 */
	PointHoverBorderColor string `json:"pointHoverBorderColor,omitempty"`

	/**
	 * PointHoverBorderWidth: Border width of point when hovered.
	 */
	PointHoverBorderWidth *float64 `json:"pointHoverBorderWidth,omitempty"`

	/**
	 * PointHoverRadius: The radius of the point when hovered.
	 */
	PointHoverRadius *float64 `json:"pointHoverRadius,omitempty"`
}
