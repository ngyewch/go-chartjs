package chartjs

import "encoding/json"

// TimeUnit

type TimeUnit string

const (
	TimeUnitMillisecond TimeUnit = "millisecond"
	TimeUnitSecond      TimeUnit = "second"
	TimeUnitMinute      TimeUnit = "minute"
	TimeUnitHour        TimeUnit = "hour"
	TimeUnitDay         TimeUnit = "day"
	TimeUnitWeek        TimeUnit = "week"
	TimeUnitMonth       TimeUnit = "month"
	TimeUnitQuarter     TimeUnit = "quarter"
	TimeUnitYear        TimeUnit = "year"
)

// TimeScaleTickOptions

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

// TimeScaleOptions

type TimeScaleOptions struct {
	*CartesianScaleOptions

	/**
	 * OffsetAfterAutoskip: If true, bar chart offsets are computed with skipped tick sizes
	 *
	 * @since 3.8.0
	 * @default false
	 */
	OffsetAfterAutoskip *bool `json:"offsetAfterAutoskip,omitempty"`

	/**
	 * Adapters: options for creating a new adapter instance
	 */
	Adapters *TimeScaleAdaptersOptions `json:"adapters,omitempty"`

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

// TimeScaleAdaptersOptions

type TimeScaleAdaptersOptions struct {
	Date any `json:"date,omitempty"`
}

// TimeScaleTimeOptions

type TimeScaleTimeOptions struct {
	/**
	 * Parser: Custom parser for dates.
	 */
	Parser string `json:"parser,omitempty"`

	/**
	 * Round: If defined, dates will be rounded to the start of this unit. See Time Units below for the allowed units.
	 */
	Round TimeUnit `json:"round,omitempty"`

	/**
	 * ISOWeekDay: If boolean and true and the unit is set to 'week', then the first day of the week will be Monday.
	 * Otherwise, it will be Sunday.
	 * If `number`, the index of the first day of the week (0 - Sunday, 6 - Saturday).
	 *
	 * @note boolean not supported
	 * @default false
	 */
	ISOWeekday *int `json:"isoWeekday,omitempty"`

	/**
	 * DisplayFormats: Sets how different time units are displayed.
	 */
	DisplayFormats map[string]string `json:"displayFormats,omitempty"`

	/**
	 * TooltipFormat: The format string to use for the tooltip.
	 */
	TooltipFormat string `json:"tooltipFormat,omitempty"`

	/**
	 * Unit: If defined, will force the unit to be a certain type. See Time Units section below for details.
	 * @default false
	 */
	Unit TimeUnit `json:"unit,omitempty"`

	/**
	 * MinUnit: The minimum display format to be used for a time unit.
	 *
	 * @default 'millisecond'
	 */
	MinUnit TimeUnit `json:"minUnit,omitempty"`
}
