package main

import (
	"github.com/AamerAijaz/genim"
	genimrl "github.com/AamerAijaz/genim/renderers/raylib"
)

func main() {
	genimrl.Init(800, 450, "genim")
	defer genimrl.Close()

	context := genim.NewContext(
		genim.ContextWithBoundingBox(100, 100),
		genim.ContextWithColor(genim.White),
	)

	for !genimrl.ShouldClose() {
		w, h := genimrl.ScreenSize()
		context.SetLayoutDimensions(w, h)

		context.
			CreateRoot("root").
			AddElement(
				genim.NewElement(
					"box1",
					genim.ElementWithColor(genim.Red),
					genim.ElementWithLayoutConfig(
						genim.NewLayoutConfig(
							genim.Size{Width: 100, Height: 100},
							genim.Position{X: 10, Y: 10},
						),
					),
				),
			).
			AddElement(
				genim.NewElement(
					"box2",
					genim.ElementWithColor(genim.Blue),
					genim.ElementWithLayoutConfig(
						genim.NewLayoutConfig(
							genim.Size{Width: 100, Height: 100},
							genim.Position{X: 120, Y: 10},
						),
					),
				),
			)

		genimrl.Draw(context.EndLayout())
	}
}
