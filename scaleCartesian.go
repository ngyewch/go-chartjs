package chartjs

type CartesianScaleOptions struct {
	*CoreScaleOptions

	/**
	 * Bounds: Scale boundary strategy (bypassed by min/max time options)
	 * - `data`: make sure data are fully visible, ticks outside are removed
	 * - `ticks`: make sure ticks are fully visible, data outside are truncated
	 *
	 * @since 2.7.0
	 * @default 'ticks'
	 */
	Bounds string `json:"bounds,omitempty"`

	/**
	 * Position of the axis.
	 */
	Position ICartesianScalePosition `json:"position,omitempty"`

	/**
	 * Stack group. Axes at the same `position` with same `stack` are stacked.
	 */
	Stack string `json:"stack,omitempty"`

	/**
	 * StackWeight: Weight of the scale in stack group.
	 * Used to determine the amount of allocated space for the scale within the group.
	 *
	 * @default 1
	 */
	StackWeight *float64 `json:"stackWeight,omitempty"`

	/**
	 * Axis: Which type of axis this is.
	 * Possible values are: 'x', 'y', 'r'.
	 * If not set, this is inferred from the first character of the ID which should be 'x', 'y' or 'r'.
	 */
	Axis string `json:"axis,omitempty"`

	/**
	 * Offset: If true, extra space is added to the both edges and the axis is scaled to fit into the chart area.
	 * This is set to true for a bar chart by default.
	 *
	 * @default false
	 */
	Offset *bool `json:"offset,omitempty"`

	Grid *GridLineOptions `json:"grid,omitempty"`

	Border *BorderOptions `json:"border,omitempty"`

	/**
	 * Title: Options for the scale title.
	 */
	Title *CartesianScaleTitle `json:"title,omitempty"`

	/**
	 * Stacked: If true, data will be comprised between datasets of data
	 *
	 * @note 'single' not supported.
	 * @default false
	 */
	Stacked *bool `json:"stacked,omitempty"`
}

type ICartesianScalePosition interface {
	isCartesianAxisPosition() bool
}

type CartesianScalePosition string

const (
	CartesianAxisPositionLeft   CartesianScalePosition = "left"
	CartesianAxisPositionTop    CartesianScalePosition = "top"
	CartesianAxisPositionRight  CartesianScalePosition = "right"
	CartesianAxisPositionBottom CartesianScalePosition = "bottom"
	CartesianAxisPositionCenter CartesianScalePosition = "center"
)

func (p CartesianScalePosition) isCartesianAxisPosition() bool {
	return true
}

type CartesianScalePositionRelative map[string]float64

func (p CartesianScalePositionRelative) isCartesianAxisPosition() bool {
	return true
}

type CartesianScaleTitle struct {
	/** Display: If true, displays the axis title. */
	Display *bool `json:"display,omitempty"`

	/** Align: Alignment of the axis title. */
	Align Align `json:"align,omitempty"`

	/** Text: The text for the title, e.g. "# of People" or "Response Choices". */
	Text string `json:"text,omitempty"`

	/** Color of the axis label. */
	Color string `json:"color,omitempty"`

	/** Font: Information about the axis title font. */
	Font *FontSpec `json:"font,omitempty"`

	/** Padding to apply around scale labels. */
	Padding ICartesianScaleTitlePadding `json:"padding,omitempty"`
}

type ICartesianScaleTitlePadding interface {
	isCartesianScaleTitlePadding()
}

type CartesianScaleTitlePadding float64

func (p CartesianScaleTitlePadding) isCartesianScaleTitlePadding() {}

type CartesianScaleTitlePaddingBasic struct {
	Top    *float64 `json:"top,omitempty"`
	Bottom *float64 `json:"bottom,omitempty"`
	Y      *float64 `json:"y,omitempty"`
}

func (p *CartesianScaleTitlePaddingBasic) isCartesianScaleTitlePadding() {}

type CartesianTickOptions struct {
	*TickOptions

	/**
	 * SampleSize: The number of ticks to examine when deciding how many labels will fit. Setting a smaller value will be faster, but may be less accurate when there is large variability in label length.
	 *
	 * @default ticks.length
	 */
	SampleSize *int `json:"sampleSize,omitempty"`

	/**
	 * Align: The label alignment
	 *
	 * @note 'inner' not supported
	 * @default 'center'
	 */
	Align Align `json:"align,omitempty"`

	/**
	 * AutoSkip: If true, automatically calculates how many labels can be shown and hides labels accordingly. Labels will be rotated up to maxRotation before skipping any.
	 * Turn autoSkip off to show all labels no matter what.
	 *
	 * @default true
	 */
	AutoSkip *bool `json:"autoSkip,omitempty"`

	/**
	 * AutoSkipPadding: Padding between the ticks on the horizontal axis when autoSkip is enabled.
	 *
	 * @default 0
	 */
	AutoSkipPadding *float64 `json:"autoSkipPadding,omitempty"`

	/**
	 * CrossAlign: How is the label positioned perpendicular to the axis direction.
	 * This only applies when the rotation is 0 and the axis position is one of "top", "left", "right", or "bottom"
	 *
	 * @default 'near'
	 */
	CrossAlign string `json:"crossAlign,omitempty"`

	/**
	 * IncludeBounds: Should the defined `min` and `max` values be presented as ticks even if they are not "nice".
	 *
	 * @default: true
	 */
	IncludeBounds *bool `json:"includeBounds,omitempty"`

	/**
	 * LabelOffset: Distance in pixels to offset the label from the centre point of the tick (in the x direction for the x axis, and the y direction for the y axis).
	 * Note: this can cause labels at the edges to be cropped by the edge of the canvas
	 *
	 * @default 0
	 */
	LabelOffset *float64 `json:"labelOffset,omitempty"`

	/**
	 * MinRotation: Minimum rotation for tick labels. Note: Only applicable to horizontal scales.
	 *
	 * @default 0
	 */
	MinRotation *float64 `json:"minRotation,omitempty"`

	/**
	 * MaxRotation: Maximum rotation for tick labels when rotating to condense labels. Note: Rotation doesn't occur until necessary. Note: Only applicable to horizontal scales.
	 *
	 * @default 50
	 */
	MaxRotation *float64 `json:"maxRotation,omitempty"`

	/**
	 * Mirror: Flips tick labels around axis, displaying the labels inside the chart instead of outside.
	 * Note: Only applicable to vertical scales.
	 *
	 * @default false
	 */
	Mirror *bool `json:"mirror,omitempty"`

	/**
	 * Padding between the tick label and the axis. When set on a vertical axis, this applies in the horizontal (X) direction. When set on a horizontal axis, this applies in the vertical (Y) direction.
	 *
	 * @default 0
	 */
	Padding *float64 `json:"padding,omitempty"`

	/**
	 * MaxTicksLimit: Maximum number of ticks and gridlines to show.
	 *
	 * @default 11
	 */
	MaxTicksLimit *int `json:"maxTicksLimit,omitempty"`
}
