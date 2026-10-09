package genim

type LayoutDirection string

const (
	TopToBottom LayoutDirection = "Top_To_Bottom"
	LeftToRight LayoutDirection = "Left_To_Right"
)

type LayoutConfig struct {
	size      Size
	position  Position
	direction LayoutDirection
}

func NewLayoutConfig(size Size, direction LayoutDirection) LayoutConfig {
	return LayoutConfig{
		size:      size,
		direction: direction,
	}
}

func (l LayoutConfig) SetPosition(position Position) LayoutConfig {
	l.position = position
	return l
}
