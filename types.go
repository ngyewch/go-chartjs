package chartjs

// Float64OrString

type Float64OrString interface {
	isFloat64OrString()
}

// Float64

type Float64 float64

func (v Float64) isFloat64OrString() {}

// String

type String string

func (v String) isFloat64OrString() {}
