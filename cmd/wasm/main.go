package main

import "github.com/AamerAijaz/genim"

func main() {
	context := genim.NewContext(genim.WithBoundingBox(800, 600))

	context.Draw()
	select {}
}
