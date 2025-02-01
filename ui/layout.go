package ui

import (
	"image"

	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
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

	fonts = gofont.Collection()
	// fontShaper = text.NewShaper(fonts)
	FontShaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	TH         = material.NewTheme()
	// th.Palette = Palette{
	// 	Fg:         rgb(0x000000),
	// 	Bg:         rgb(0xffffff),
	// 	ContrastBg: rgb(0x3f51b5),
	// 	ContrastFg: rgb(0xffffff),
	// }
	// th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
)

func Rows(children ...layout.FlexChild) W {
	return func(c C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(c, children...)
	}
}

func Columns(children ...layout.FlexChild) W {
	return func(c C) D {
		return layout.Flex{Axis: layout.Horizontal}.Layout(c, children...)
	}
}

func ColumnsVCentered(children ...layout.FlexChild) W {
	return func(c C) D {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(c, children...)
	}
}

// The widget is called with the context X constraints minimum cleared.
func E(w W) W {
	return func(c C) D {
		return layout.E.Layout(c, w)
	}
}

func EmptyWidget(c C) D { return D{} }

func Wrap(w W, wrappers ...Wrapper) W {
	for i := len(wrappers) - 1; i >= 0; i-- {
		w = wrappers[i](w)
	}

	return w
}

func LayoutToWidget(r func(C, W) D, w W) W {
	return func(c C) D {
		return r(c, w)
	}
}

func LayoutToWrapper(r func(C, W) D) func(w W) W {
	return func(w W) W {
		return func(c C) D {
			return r(c, w)
		}
	}
}

func Centered(w W) W {
	return func(c C) D {
		v := layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceAround}
		h := layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceAround}
		return h.Layout(c, Rigid(func(c C) D {
			return v.Layout(c, Rigid(w))
		}))
	}
}

func Constraint(width, height int, w W) W {
	return func(c C) D {
		wdp := c.Metric.Dp(DP(width))
		hdp := c.Metric.Dp(DP(height))
		c.Constraints.Max = P{wdp, hdp}
		return w(c)
	}
}

func ConstraintW(width int, w W) W {
	return func(c C) D {
		c.Constraints.Min.X = c.Metric.Dp(DP(width))
		c.Constraints.Max.X = c.Metric.Dp(DP(width))
		return w(c)
	}
}
