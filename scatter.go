package chartjs

type ScatterChartDataset struct {
	Data []Point `json:"data,omitempty"`

	BackgroundColor  string    `json:"backgroundColor,omitempty"`
	BorderCapStyle   string    `json:"borderCapStyle,omitempty"`
	BorderColor      string    `json:"borderColor,omitempty"`
	BorderDash       []float32 `json:"borderDash,omitempty"`
	BorderDashOffset float32   `json:"borderDashOffset,omitempty"`
	BorderJoinStyle  string    `json:"borderJoinStyle,omitempty"`
	BorderWidth      float32   `json:"borderWidth,omitempty"`
	// Clip
	CubicInterpolationMode string `json:"cubicInterpolationMode,omitempty"`
	// TODO
	SpanGaps bool   `json:"spanGaps"` // TODO
	Stack    string `json:"stack,omitempty"`
	// Stepped
	Tension float32 `json:"tension,omitempty"`
	XAxisID string  `json:"xAxisID,omitempty"`
	YAxisID string  `json:"yAxisID,omitempty"`
}

type ScatterChartData struct {
	Datasets []*ScatterChartDataset `json:"datasets,omitempty"`
}

type ScatterChartConfig struct {
	Type    string            `json:"type,omitempty"`
	Data    *ScatterChartData `json:"data,omitempty"`
	Options *CoreChartOptions `json:"options,omitempty"`
}
