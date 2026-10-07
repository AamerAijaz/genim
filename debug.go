package genim

import "fmt"

// DebugDraw prints each RenderCommand to stdout in a human-readable format.
func DebugDraw(cmds []RenderCommand) {
	fmt.Printf("RenderCommands (%d total):\n", len(cmds))
	for i, cmd := range cmds {
		c := cmd.Color
		fmt.Printf(
			"  [%d] id=%-20q  pos=(%6.1f, %6.1f)  size=(%6.1f x %6.1f)  color=rgba(%d,%d,%d,%d)\n",
			i, cmd.ID,
			cmd.Position.X, cmd.Position.Y,
			cmd.BoundingBox.Width, cmd.BoundingBox.Height,
			c.R, c.G, c.B, c.A,
		)
	}
}
