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
		context.CreateRoot("root")

		context.
			OpenElement(
				genim.NewElement(
					"outer-container",
					genim.ElementWithColor(genim.White),
					genim.ElementWithLayoutConfig(
						genim.NewLayoutConfig(
							genim.Size{Width: 400, Height: 400},
							genim.TopToBottom,
						),
					),
				),
			).
			OpenElement(
				genim.NewElement(
					"header",
					genim.ElementWithColor(genim.Blue),
					genim.ElementWithLayoutConfig(
						genim.NewLayoutConfig(
							genim.Size{Width: 400, Height: 100},
							genim.LeftToRight,
						),
					),
				),
			).
			CloseElement().
			OpenElement(
				genim.NewElement(
					"main-content",
					genim.ElementWithLayoutConfig(
						genim.NewLayoutConfig(
							genim.Fit,
							genim.LeftToRight,
						),
					),
				),
			).
			OpenElement(
				genim.NewElement(
					"sidebar",
					genim.ElementWithColor(genim.Red),
					genim.ElementWithLayoutConfig(
						genim.NewLayoutConfig(
							genim.Size{Width: 100, Height: 300},
							genim.TopToBottom,
						),
					),
				),
			).
			CloseElement().
			OpenElement(
				genim.NewElement(
					"content",
					genim.ElementWithColor(genim.Black),
					genim.ElementWithLayoutConfig(
						genim.NewLayoutConfig(
							genim.Size{Width: 300, Height: 300},
							genim.LeftToRight,
						),
					),
				),
			).
			CloseElement().
			CloseElement().
			CloseElement()

		genimrl.Draw(context.EndLayout())
	}
}
