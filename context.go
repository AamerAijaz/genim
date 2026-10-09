package genim

import "fmt"

// Context represents the state of the current UI.
// It contains all the data needed to draw the layout for a single frame.
type Context struct {
	elements        []*Element
	openElements    []*Element
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
			NewLayoutConfig(Size{Width: c.boundingBox.Width, Height: c.boundingBox.Height}, LeftToRight),
		).
		SetPosition(Position{X: 0, Y: 0}).
		SetColor(c.backgroundColor)

	c.OpenElement(root)

	return c
}

func (c *Context) OpenElement(element *Element) *Context {
	c.elements = append(c.elements, element)

	if len(c.openElements) > 0 {
		parent := c.openElements[len(c.openElements)-1]
		element.parent = parent
		parent.children = append(parent.children, element)

		x := parent.layoutConfig.position.X
		y := parent.layoutConfig.position.Y

		for i := 0; i < len(parent.children)-1; i++ {
			child := parent.children[i]
			if parent.layoutConfig.direction == TopToBottom {
				y += child.layoutConfig.position.Y + child.layoutConfig.size.Height
			} else {
				x += child.layoutConfig.position.X + child.layoutConfig.size.Width
			}
		}

		element.layoutConfig.position.X = x
		element.layoutConfig.position.Y = y
	}

	c.openElements = append(c.openElements, element)

	return c
}

func (c *Context) CloseElement() *Context {
	c.openElements = c.openElements[0 : len(c.openElements)-1]

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

	c.debugRenderCommands()

	renderCommands := make([]RenderCommand, len(c.renderCommands))
	copy(renderCommands, c.renderCommands)

	c.elements = c.elements[0:0]
	c.openElements = c.openElements[0:0]
	c.renderCommands = c.renderCommands[0:0]

	return renderCommands
}

func (c *Context) debugRenderCommands() {
	fmt.Printf("=== Render Commands (%d) ===\n", len(c.renderCommands))
	for i, cmd := range c.renderCommands {
		fmt.Printf(
			"  [%d] ID: %-20s | Pos: (%.1f, %.1f) | Size: %.1f x %.1f | Color: RGBA(%d, %d, %d, %d)\n",
			i, cmd.ID,
			cmd.Position.X, cmd.Position.Y,
			cmd.BoundingBox.Width, cmd.BoundingBox.Height,
			cmd.Color.R, cmd.Color.G, cmd.Color.B, cmd.Color.A,
		)
	}
	fmt.Println("===========================")
}

func (c *Context) SetLayoutDimensions(width, height float32) *Context {
	c.boundingBox = BoundingBox{Width: width, Height: height}
	return c
}
