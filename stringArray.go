package chartjs

type StringOrStringArray interface {
	isStringOrStringArray() bool
}

type StringArray []string

func (StringArray) isStringOrStringArray() bool {
	return true
}

type String string

func (String) isStringOrStringArray() bool {
	return true
}

func Strings(str ...string) StringOrStringArray {
	if len(str) == 0 {
		return nil
	} else if len(str) == 1 {
		return String(str[0])
	} else {
		return StringArray(str)
	}
}
