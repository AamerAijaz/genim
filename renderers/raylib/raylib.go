package raylib

import (
	"github.com/AamerAijaz/genim"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init(width, height int32, title string) {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(60)
}

func Close() {
	rl.CloseWindow()
}

func ShouldClose() bool {
	return rl.WindowShouldClose()
}

func ScreenSize() (width, height float32) {
	return float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())
}

func Draw(cmds []genim.RenderCommand) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	for _, cmd := range cmds {
		c := cmd.Color
		rl.DrawRectangle(
			int32(cmd.Position.X),
			int32(cmd.Position.Y),
			int32(cmd.BoundingBox.Width),
			int32(cmd.BoundingBox.Height),
			rl.NewColor(c.R, c.G, c.B, c.A),
		)
	}

	rl.EndDrawing()
}
