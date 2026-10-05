//go:build js && wasm

package web

import "syscall/js"

func GetWindowHeight() float32 {
	height := js.Global().Get("innerHeight").Int()
	return float32(height)
}

func GetWindowWidth() float32 {
	width := js.Global().Get("innerWidth").Int()
	return float32(width)
}

func Document() js.Value {
	return js.Global().Get("document")
}

func GetRootElement() js.Value {
	document := Document()
	root := document.Call("getElementById", "root")
	if root.IsUndefined() || root.IsNull() {
		panic("Invalid index.html, root not present")
	}

	return root
}
