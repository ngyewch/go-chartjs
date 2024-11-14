package chartjs

type LegendOptions struct {
	/**
	 * Display, Is the legend shown?
	 *
	 * @default true
	 */
	Display *bool `json:"display,omitempty"`

	/**
	 * Position of the legend.
	 *
	 * @default 'top'
	 */
	Position ILayoutPosition `json:"position,omitempty"`

	/**
	 * Align is the alignment of the legend.
	 *
	 * @default 'center'
	 */
	Align Align `json:"align,omitempty"`

	/**
	 * MaxHeight, maximum height of the legend, in pixels
	 */
	MaxHeight *float64 `json:"maxHeight,omitempty"`

	/**
	 * MaxWidth, maximum width of the legend, in pixels
	 */
	MaxWidth *float64 `json:"maxWidth,omitempty"`

	/**
	 * FullSize marks that this box should take the full width/height of the canvas (moving other boxes). This is unlikely to need to be changed in day-to-day use.
	 *
	 * @default true
	 */
	FullSize *bool `json:"fullSize,omitempty"`

	/**
	 * Reverse, legend will show datasets in reverse order.
	 *
	 * @default false
	 */
	Reverse *bool `json:"reverse,omitempty"`

	// TODO Labels

	/**
	 * Rtl, true for rendering the legends from right to left.
	 */
	Rtl *bool `json:"rtl,omitempty"`

	/**
	 * TextDirection, this will force the text direction 'rtl' or 'ltr' on the canvas for rendering the legend, regardless of the css specified on the canvas
	 */
	TextDirection string `json:"textDirection,omitempty"`

	Title *LegendTitleOptions `json:"title,omitempty"`
}

type LegendTitleOptions struct {
	/**
	 * Display, is the legend title displayed.
	 *
	 * @default false
	 */
	Display *bool `json:"display,omitempty"`

	/**
	 * Color of title
	 */
	Color string `json:"color,omitempty"`

	Font *FontSpec `json:"font,omitempty"`

	Position LegendTitlePosition `json:"position,omitempty"`

	Padding IPadding `json:"padding,omitempty"`

	/**
	 * Text, the string title.
	 */
	Text string `json:"text,omitempty"`
}

type LegendTitlePosition string

const (
	LegendTitlePositionCenter LegendTitlePosition = "center"
	LegendTitlePositionStart  LegendTitlePosition = "start"
	LegendTitlePositionEnd    LegendTitlePosition = "end"
)
