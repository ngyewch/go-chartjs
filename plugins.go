package chartjs

type IPlugin interface {
	isPlugin() bool
}

type LegendPlugin struct {
	Display   bool   `json:"display"`
	Position  string `json:"position,omitempty"`
	Align     string `json:"align,omitempty"`
	MaxHeight int    `json:"maxHeight,omitempty"`
	MaxWidth  int    `json:"maxWidth,omitempty"`
	FullSize  bool   `json:"fullSize"`
	Reverse   bool   `json:"reverse,omitempty"`
	// Labels
	Rtl           bool         `json:"rtl,omitempty"`
	TextDirection string       `json:"textDirection,omitempty"`
	Title         *LegendTitle `json:"title,omitempty"`
}

type LegendTitle struct {
	Color   string `json:"color,omitempty"`
	Display bool   `json:"display"`
	// Font
	// Padding
	Text string `json:"text,omitempty"`
}

func (plugin *LegendPlugin) isPlugin() bool {
	return true
}

type TitlePlugin struct {
	Align    string `json:"align,omitempty"`
	Color    string `json:"color,omitempty"`
	Display  bool   `json:"display"`
	FullSize bool   `json:"fullSize"`
	Position string `json:"position,omitempty"`
	// Font
	// Padding
	Text string `json:"text,omitempty"`
}

func (plugin *TitlePlugin) isPlugin() bool {
	return true
}
