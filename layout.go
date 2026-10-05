package genim

type LayoutConfig struct {
	size     Size
	position Position
}

var defaultLayoutConfig = LayoutConfig{
	size:     Size{Width: 0, Height: 0},
	position: Position{X: -1, Y: -1},
}

func NewLayoutConfig(size Size, position Position) LayoutConfig {
	return LayoutConfig{
		size:     size,
		position: position,
	}
}
