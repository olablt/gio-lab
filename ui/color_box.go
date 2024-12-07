package ui

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
)

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func FillWithLabel(gtx layout.Context, th material.Theme, text string, fg, bg color.NRGBA) layout.Dimensions {
	th.Palette.Fg = fg
	th.Palette.Bg = bg
	ColorBox(gtx, gtx.Constraints.Min, bg)
	// return layout.Center.Layout(gtx, material.Label(&th, unit.Sp(10), text).Layout)
	return layout.Center.Layout(gtx, material.Body1(&th, text).Layout)
}

// FillWithLabelH3 creates a label with the specified text and background color
func FillWithLabelH3(gtx layout.Context, th *material.Theme, text string, backgroundColor color.NRGBA) layout.Dimensions {
	ColorBox(gtx, gtx.Constraints.Max, backgroundColor)
	return layout.Center.Layout(gtx, material.H3(th, text).Layout)
}
