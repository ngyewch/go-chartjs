package chartjs

type ChartArea struct {
	Top    *float64 `json:"top,omitempty"`
	Left   *float64 `json:"left,omitempty"`
	Right  *float64 `json:"right,omitempty"`
	Bottom *float64 `json:"bottom,omitempty"`
	Width  *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
}

type Point struct {
	X float64  `json:"x"`
	Y *float64 `json:"y"`
}
