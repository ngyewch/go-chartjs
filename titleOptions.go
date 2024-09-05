package chartjs

type TitlePosition string

const (
	TitlePositionLeft   TitlePosition = "left"
	TitlePositionTop    TitlePosition = "top"
	TitlePositionBottom TitlePosition = "bottom"
	TitlePositionRight  TitlePosition = "right"
)

type TitlePadding struct {
	Bottom float64 `json:"bottom,omitempty"`
	Top    float64 `json:"top,omitempty"`
}

type TitleOptions struct {
	Align    Align               `json:"align,omitempty"` // default: 'center'
	Color    string              `json:"color,omitempty"`
	Display  bool                `json:"display,omitempty"` // default: false
	Font     *FontSpec           `json:"font,omitempty"`
	FullSize bool                `json:"fullSize,omitempty"` // default: true [TODO]
	Padding  *TitlePadding       `json:"padding,omitempty"`
	Position TitlePosition       `json:"position,omitempty"` // default: 'top'
	Text     StringOrStringArray `json:"text,omitempty"`
}
