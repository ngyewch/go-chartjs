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

	/**
	 * OffsetAfterAutoskip: If true, bar chart offsets are computed with skipped tick sizes
	 *
	 * @since 3.8.0
	 * @default false
	 */
	OffsetAfterAutoskip *bool `json:"offsetAfterAutoskip,omitempty"`

	// TODO Adapters

	Time *TimeScaleTimeOptions `json:"time,omitempty"`

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

type TimeScaleTimeOptions struct {
	// TODO parser
	// TODO round
	// TODO isoWeekday

	/**
	 * DisplayFormats: Sets how different time units are displayed.
	 */
	DisplayFormats map[string]string `json:"displayFormats,omitempty"`

	/**
	 * TooltipFormat: The format string to use for the tooltip.
	 */
	TooltipFormat string `json:"tooltipFormat,omitempty"`

	// TODO unit
	// TODO minUnit
}
