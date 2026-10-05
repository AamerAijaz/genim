package main

import "github.com/AamerAijaz/genim"

func main() {
	context := genim.
		NewContext(genim.WithBoundingBox(100, 100))

	root := genim.
		NewElement("Root").
		SetLayoutConfig(
			genim.NewLayoutConfig(genim.Size{Width: 100, Height: 100}, genim.Position{X: 0, Y: 0}),
		).
		SetColor(genim.Color{R: 255, G: 0, B: 255, A: 255})

	for {
		context.
			AddElement(root).
			AddRenderCommand(root.GenerateRenderCommand())

		renderCommands := context.EndLayout()
		genim.Draw(renderCommands)
	}
}
