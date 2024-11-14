package chartjs

type PluginOptions struct {
	// TODO Colors
	// TODO Decimation
	// TODO Filler
	Legend   *LegendOptions `json:"legend,omitempty"`
	Subtitle *TitleOptions  `json:"subtitle,omitempty"`
	Title    *TitleOptions  `json:"title,omitempty"`
	// TODO Tooltip
}
