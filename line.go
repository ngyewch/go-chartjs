package chartjs

type LineChartDataset struct {
	*Dataset
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

type LineChartData struct {
	Datasets []*LineChartDataset `json:"datasets,omitempty"`
}

type LineChartConfig struct {
	Type    string         `json:"type,omitempty"`
	Data    *LineChartData `json:"data,omitempty"`
	Options *Options       `json:"options,omitempty"`
}
