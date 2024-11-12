package chartjs

type Options struct {
	AspectRatio         float32                  `json:"aspectRatio,omitempty"`
	MaintainAspectRatio bool                     `json:"maintainAspectRatio"`
	ResizeDelay         int                      `json:"resizeDelay,omitempty"`
	Responsive          bool                     `json:"responsive"`
	Scales              map[string]IScaleOptions `json:"scales,omitempty"`
	Plugins             map[string]IPlugin       `json:"plugins,omitempty"`
}
