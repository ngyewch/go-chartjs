package chartjs

type CategoryScaleOptions struct {
	/**
	 * Min: User defined minimum value for the scale, overrides minimum value from data.
	 */
	Min Float64OrString `json:"min,omitempty"`

	/**
	 * Max: User defined maximum value for the scale, overrides maximum value from data.
	 */
	Max Float64OrString `json:"max,omitempty"`

	// @note Not supported: string[][];
	Labels []string `json:"labels,omitempty"`
}
