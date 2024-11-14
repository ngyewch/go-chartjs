package chartjs

import "encoding/json"

type LinearTickOptions struct {
	*CartesianTickOptions

	/**
	 * Format: The Intl.NumberFormat options used by the default label formatter
	 */
	Format map[string]any `json:"format,omitempty"`

	/**
	 * Precision: if defined and stepSize is not specified, the step size will be rounded to this many decimal places.
	 */
	Precision *int `json:"precision,omitempty"`

	/**
	 * StepSize: User defined fixed step size for the scale
	 */
	StepSize *float64 `json:"stepSize,omitempty"`

	/**
	 * Count: User defined count of ticks
	 */
	Count *int `json:"count,omitempty"`
}

type LinearScaleOptions struct {
	*CartesianScaleOptions

	/**
	 * BeginAtZero: if true, scale will include 0 if it is not already included.
	 *
	 * @default true
	 */
	BeginAtZero *bool `json:"beginAtZero,omitempty"`

	// TODO Grace

	Ticks *LinearTickOptions `json:"ticks,omitempty"`
}

func (scale *LinearScaleOptions) isScaleOptions() {
}

func (scale *LinearScaleOptions) isCartesianScaleType() {
}

func (scale *LinearScaleOptions) MarshalJSON() ([]byte, error) {
	type Proxy LinearScaleOptions
	return json.Marshal(struct {
		Type string `json:"type"`
		Proxy
	}{
		Type:  "linear",
		Proxy: Proxy(*scale),
	})
}
