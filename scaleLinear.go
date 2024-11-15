package chartjs

import "encoding/json"

// LinearTickOptions

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

// LinearScaleOptions

type LinearScaleOptions struct {
	*CartesianScaleOptions

	/**
	 * Min: User defined minimum value for the scale, overrides minimum value from data.
	 */
	Min *float64 `json:"min,omitempty"`

	/**
	 * Max: User defined maximum value for the scale, overrides maximum value from data.
	 */
	Max *float64 `json:"max,omitempty"`

	/**
	 * SuggestedMin: Adjustment used when calculating the maximum data value.
	 */
	SuggestedMin *float64 `json:"suggestedMin,omitempty"`

	/**
	 * SuggestedMax: Adjustment used when calculating the minimum data value.
	 */
	SuggestedMax *float64 `json:"suggestedMax,omitempty"`

	/**
	 * BeginAtZero: if true, scale will include 0 if it is not already included.
	 *
	 * @default true
	 */
	BeginAtZero *bool `json:"beginAtZero,omitempty"`

	/**
	 * Grace: Percentage (string ending with %) or amount (number) for added room in the scale range above and below data.
	 */
	Grace Float64OrString `json:"grace,omitempty"`

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
