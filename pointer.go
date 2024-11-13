package chartjs

func PointerOf[T any](v T) *T {
	return &v
}
