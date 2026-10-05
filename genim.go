package genim

type RenderCommand struct {
	id          ElementId
	boundingBox BoundingBox
	color       Color
	position    Position
}

type Element struct {
	id           ElementId
	color        Color
	layoutConfig LayoutConfig
}

func WithLayoutConfig(layoutConfig LayoutConfig) func(*Element) {
	return func(e *Element) {
		e.layoutConfig = layoutConfig
	}
}

func WithColor(color Color) func(*Element) {
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
		id:          e.id,
		color:       e.color,
		boundingBox: BoundingBox(e.layoutConfig.size),
		position:    e.layoutConfig.position,
	}
}
