package genim

// Context represents the state of the current UI.
// It contains all the data needed to draw the layout for a single frame.
type Context struct {
	elements        []*Element
	renderCommands  []RenderCommand
	boundingBox     BoundingBox
	backgroundColor Color
}

func ContextWithBoundingBox(width, height float32) func(*Context) {
	return func(c *Context) {
		c.boundingBox = BoundingBox{
			Width:  width,
			Height: height,
		}
	}
}

func ContextWithColor(color Color) func(*Context) {
	return func(ctx *Context) {
		ctx.backgroundColor = color
	}
}

func NewContext(options ...func(*Context)) *Context {
	ctx := &Context{
		boundingBox: BoundingBox{
			Width:  0,
			Height: 0,
		},
		elements: make([]*Element, 0),
	}

	for _, op := range options {
		op(ctx)
	}

	return ctx
}

func (c *Context) CreateRoot(title string) *Context {
	root := NewElement(title).
		SetLayoutConfig(
			NewLayoutConfig(Size{Width: c.boundingBox.Width, Height: c.boundingBox.Height}, Position{X: 0, Y: 0}),
		).
		SetColor(c.backgroundColor)

	c.AddElement(root)

	return c
}

func (c *Context) AddElement(element *Element) *Context {
	c.elements = append(c.elements, element)
	return c
}

func (c *Context) AddRenderCommand(cmd RenderCommand) *Context {
	c.renderCommands = append(c.renderCommands, cmd)
	return c
}

func (c *Context) EndLayout() []RenderCommand {
	for _, e := range c.elements {
		c.AddRenderCommand(e.GenerateRenderCommand())
	}

	return c.renderCommands
}

func (c *Context) SetLayoutDimensions(width, height float32) *Context {
	c.boundingBox = BoundingBox{Width: width, Height: height}
	return c
}
