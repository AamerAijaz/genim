package genim

import "image/color"

type Vector2 struct {
	X float32
	Y float32
}

type BoundingBox struct {
	Width  float32
	Height float32
}

type Position Vector2
type Size BoundingBox

type ElementId string

type Color color.RGBA
