package widgets

import (
	"gioui.org/layout"
	"image"
)

// Widget represents a UI component that can be laid out
type Widget interface {
	Layout(gtx layout.Context) layout.Dimensions
}

// Wrapper modifies a widget's behavior
type Wrapper func(Widget) Widget

// Dimensions aliases layout.Dimensions for convenience
type Dimensions = layout.Dimensions

// Context aliases layout.Context for convenience
type Context = layout.Context

// Point aliases image.Point for convenience
type Point = image.Point
