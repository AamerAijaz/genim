package genim

type RenderCommand struct {
	ID          ElementId
	BoundingBox BoundingBox
	Color       Color
	Position    Position
}

type Element struct {
	id           ElementId
	color        Color
	layoutConfig LayoutConfig
}

func ElementWithLayoutConfig(layoutConfig LayoutConfig) func(*Element) {
	return func(e *Element) {
		e.layoutConfig = layoutConfig
	}
}

func ElementWithColor(color Color) func(*Element) {
	return func(e *Element) {
		e.color = color
	}
}

func NewElement(id string, options ...func(*Element)) *Element {
	element := &Element{
		id:    ElementId(id),
		color: Color{R: 255, G: 0, B: 255, A: 255},
		layoutConfig: LayoutConfig{
			size:     Size{Width: 0, Height: 0},
			position: Position{X: 0, Y: 0},
		},
	}

	for _, opt := range options {
		opt(element)
	}

	return element
}

func (e *Element) SetLayoutConfig(config LayoutConfig) *Element {
	e.layoutConfig = config

	return e
}

func (e *Element) SetColor(color Color) *Element {
	e.color = color
	return e
}

func (e *Element) GenerateRenderCommand() RenderCommand {
	return RenderCommand{
		ID:          e.id,
		Color:       e.color,
		BoundingBox: BoundingBox(e.layoutConfig.size),
		Position:    e.layoutConfig.position,
	}
}
