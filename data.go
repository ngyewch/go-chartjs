package chartjs

type Dataset struct {
	Data    []Point           `json:"data,omitempty"`
	Label   string            `json:"label,omitempty"`
	Clip    *Clip             `json:"clip,omitempty"`
	Order   float32           `json:"order,omitempty"`
	Stack   string            `json:"stack,omitempty"`
	Parsing map[string]string `json:"parsing,omitempty"`
	Hidden  bool              `json:"hidden,omitempty"`
}

type Clip struct {
	Left   float32 `json:"left"`
	Top    float32 `json:"top"`
	Right  float32 `json:"right"`
	Bottom float32 `json:"bottom"`
}
