package cpalette

import (
	"image"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/unit"
)

type (
	C       = layout.Context
	D       = layout.Dimensions
	W       = layout.Widget
	P       = image.Point
	DP      = unit.Dp
	SP      = unit.Sp
	Wrapper = func(W) W
	List    = layout.List
)

var (
	Flexed = layout.Flexed
	Rigid  = layout.Rigid
)

var (
	Pt           = f32.Pt
	SpaceUnit DP = 8
	// BorderSize DP = 1
	BorderSize DP = 0.5

	// fonts = gofont.Collection()
	// fontShaper = text.NewShaper(fonts)
	// FontShaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	// TH         = material.NewTheme()
	// th.Palette = Palette{
	// 	Fg:         rgb(0x000000),
	// 	Bg:         rgb(0xffffff),
	// 	ContrastBg: rgb(0x3f51b5),
	// 	ContrastFg: rgb(0xffffff),
	// }
	// th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
)
