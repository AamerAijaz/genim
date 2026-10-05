//go:build js && wasm

package genim

import (
	"fmt"
	"syscall/js"
)

const canvasID = "genim-canvas"

func Draw(cmds []RenderCommand) {
	canvas := js.Global().Get("document").Call("getElementById", canvasID)
	if canvas.IsNull() || canvas.IsUndefined() {
		panic("genim: element #" + canvasID + " not found")
	}

	ctx := canvas.Call("getContext", "2d")
	ctx.Call("clearRect", 0, 0, canvas.Get("width").Float(), canvas.Get("height").Float())

	for _, cmd := range cmds {
		c := cmd.color
		ctx.Set("fillStyle", fmt.Sprintf("rgba(%d,%d,%d,%f)", c.R, c.G, c.B, float64(c.A)/255))
		ctx.Call("fillRect", cmd.position.X, cmd.position.Y, cmd.boundingBox.Width, cmd.boundingBox.Height)
	}
}
