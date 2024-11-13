package chartjs

type Align string

const (
	AlignCenter Align = "center"
	AlignStart  Align = "start"
	AlignEnd    Align = "end"
)

func (a Align) Pointer() *Align {
	return &a
}

type TextAlign string

const (
	TextAlignLeft   TextAlign = "left"
	TextAlignCenter TextAlign = "center"
	TextAlignRight  TextAlign = "right"
)

func (ta TextAlign) Pointer() *TextAlign {
	return &ta
}
