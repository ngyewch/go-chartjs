package chartjs

type LogarithmicScaleOptions struct {
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

	Ticks *LogarithmicTickOptions `json:"ticks,omitempty"`
}

type LogarithmicTickOptions struct {
	*CartesianTickOptions

	/**
	 * Format: The Intl.NumberFormat options used by the default label formatter
	 */
	Format map[string]any `json:"format,omitempty"`
}
