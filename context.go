package genim

// Context represents the state of the current UI.
// It contains all the data needed to draw the layout for a single frame.
type Context struct {
	elements       []*Element
	renderCommands []RenderCommand
	boundingBox    BoundingBox
}

func WithBoundingBox(width, height float32) func(*Context) {
	return func(c *Context) {
		c.boundingBox = BoundingBox{
			Width:  width,
			Height: height,
		}
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

func (c *Context) AddElement(element *Element) *Context {
	c.elements = append(c.elements, element)
	return c
}

func (c *Context) AddRenderCommand(cmd RenderCommand) *Context {
	c.renderCommands = append(c.renderCommands, cmd)
	return c
}

func (c *Context) EndLayout() []RenderCommand {
	return c.renderCommands
}

func (c *Context) Draw() {
	Draw(c.renderCommands)
}

func (c *Context) SetLayoutDimensions(width, height float32) *Context {
	c.boundingBox = BoundingBox{Width: width, Height: height}
	return c
}
