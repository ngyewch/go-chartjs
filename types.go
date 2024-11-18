package chartjs

// Float64OrString

type Float64OrString interface {
	isFloat64OrString()
}

// StringOrStringArray

type StringOrStringArray interface {
	isStringOrStringArray()
}

// IStepped

type IStepped interface {
	isIStepped()
}

type IFillTarget interface {
	isIFillTarget()
}

// Float64

type Float64 float64

func (v Float64) isFloat64OrString() {}

func (v Float64) isIFillTarget() {}

// String

type String string

func (v String) isFloat64OrString() {}

func (v String) isStringOrStringArray() {}

func (v String) isIFillTarget() {}

// StringArray

type StringArray []string

func (v StringArray) isStringOrStringArray() {}

// Bool

type Bool bool

func (v Bool) isIStepped() {}

func (v Bool) isIFillTarget() {}
