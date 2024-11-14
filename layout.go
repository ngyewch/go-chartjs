package chartjs

type ILayoutPosition interface {
	isLayoutPosition()
}

type LayoutPosition string

const (
	LayoutPositionLeft      LayoutPosition = "left"
	LayoutPositionTop       LayoutPosition = "top"
	LayoutPositionRight     LayoutPosition = "right"
	LayoutPositionBottom    LayoutPosition = "bottom"
	LayoutPositionCenter    LayoutPosition = "center"
	LayoutPositionChartArea LayoutPosition = "chartArea"
)

func (p LayoutPosition) isLayoutPosition() {}

type LayoutPositionRelative map[string]float64

func (p LayoutPositionRelative) isLayoutPosition() {}
