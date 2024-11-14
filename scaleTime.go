package chartjs

import "encoding/json"

type TimeScaleTickOptions struct {
	*CartesianTickOptions

	/**
	 * Source: Ticks generation input values:
	 * - 'auto': generates "optimal" ticks based on scale size and time options.
	 * - 'data': generates ticks from data (including labels from data `{t|x|y}` objects).
	 * - 'labels': generates ticks from user given `data.labels` values ONLY.
	 *
	 * @see https://github.com/chartjs/Chart.js/pull/4507
	 * @since 2.7.0
	 * @default 'auto'
	 */
	Source string `json:"source,omitempty"`

	/**
	 * StepSize: The number of units between grid lines.
	 *
	 * @default 1
	 */
	StepSize *float64 `json:"stepSize,omitempty"`
}

type TimeScaleOptions struct {
	*CartesianScaleOptions

	OffsetAfterAutoskip *bool `json:"offsetAfterAutoskip,omitempty"` // default: false
	// Adapters
	// Time
	Ticks *TimeScaleTickOptions `json:"ticks,omitempty"`
}

func (scale *TimeScaleOptions) isScaleOptions() {
}

func (scale *TimeScaleOptions) isCartesianScaleType() {
}

func (scale *TimeScaleOptions) MarshalJSON() ([]byte, error) {
	type Proxy TimeScaleOptions
	return json.Marshal(struct {
		Type string `json:"type"`
		Proxy
	}{
		Type:  "time",
		Proxy: Proxy(*scale),
	})
}
