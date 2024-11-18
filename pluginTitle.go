package chartjs

type TitlePosition string

const (
	TitlePositionLeft   TitlePosition = "left"
	TitlePositionTop    TitlePosition = "top"
	TitlePositionBottom TitlePosition = "bottom"
	TitlePositionRight  TitlePosition = "right"
)

type TitleOptions struct {
	/**
	 * Align, alignment of the title.
	 *
	 * @default 'center'
	 */
	Align Align `json:"align,omitempty"`

	/**
	 * Display, is the title shown?
	 *
	 * @default false
	 */
	Display *bool `json:"display,omitempty"`

	/**
	 * Position of title
	 *
	 * @default 'top'
	 */
	Position TitlePosition `json:"position,omitempty"`

	/**
	 * Color of text
	 */
	Color string `json:"color,omitempty"`

	Font *FontSpec `json:"font,omitempty"`

	/**
	 * FullSize, marks that this box should take the full width/height of the canvas (moving other boxes).
	 * If set to `false`, places the box above/beside the chart area
	 *
	 * @default true
	 */
	FullSize *bool `json:"fullSize,omitempty"`

	/**
	 * Padding, adds padding above and below the title text if a single number is specified.
	 * It is also possible to change top and bottom padding separately.
	 */
	Padding ITitlePadding `json:"padding,omitempty"`

	/**
	 * Title text to display. If specified as an array, text is rendered on multiple lines.
	 */
	Text StringOrStringArray `json:"text,omitempty"`
}

type ITitlePadding interface {
	isTitlePadding()
}

type TitlePadding struct {
	/** Top, padding on the (relative) top side of this axis label. */
	Top *float64 `json:"top,omitempty"`

	/** Bottom, padding on the (relative) bottom side of this axis label. */
	Bottom *float64 `json:"bottom,omitempty"`

	/** Y is a shorthand for defining top/bottom to the same values. */
	Y *float64 `json:"y,omitempty"`
}

func (p TitlePadding) isTitlePadding() {}
