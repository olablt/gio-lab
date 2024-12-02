package ui

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
)

var Bold = FontWeight(font.Bold)
var AlignStart = TextAlignment(text.Start)
var AlignMiddle = TextAlignment(text.Middle)
var AlignEnd = TextAlignment(text.End)

var OneLine = MaxLines(1)
var IconSize = FontEnlarge(2)

// Label - one line text
func Label(s string) W { return OneLine(Text(s)) }

// H1 - one line text header
func H1(s string) W { h1 := FontEnlarge(2); return h1(OneLine(Text(s))) }

func Text(s string) W {
	return func(c C) D {
		tl := widget.Label{Alignment: Theme.TextAlignment, MaxLines: Theme.MaxLines}
		// tl.LineHeight = 1
		// tl.LineHeightScale = 1
		paint.ColorOp{Color: Theme.TextColor}.Add(c.Ops)
		return tl.Layout(c, Theme.FontFamily, font.Font{Weight: Theme.FontWeight}, Theme.FontSize, s, op.CallOp{})
	}
}

func FontEnlarge(s float32) Wrapper {
	return FontSize(SP(s) * Theme.FontSize)
}

func FontSize(s SP) Wrapper {
	return func(w W) W {
		return func(c C) D {
			old := Theme.FontSize
			Theme.FontSize = s
			d := w(c)
			Theme.FontSize = old
			return d
		}
	}
}

func MaxLines(i int) Wrapper {
	return func(w W) W {
		return func(c C) D {
			old := Theme.MaxLines
			Theme.MaxLines = i
			d := w(c)
			Theme.MaxLines = old
			return d
		}
	}
}

func Font(f *text.Shaper) Wrapper {
	return func(w W) W {
		return func(c C) D {
			old := Theme.FontFamily
			Theme.FontFamily = f
			d := w(c)
			Theme.FontFamily = old
			return d
		}
	}
}

func FontWeight(f font.Weight) Wrapper {
	return func(w W) W {
		return func(c C) D {
			old := Theme.FontWeight
			Theme.FontWeight = f
			d := w(c)
			Theme.FontWeight = old
			return d
		}
	}
}

func TextAlignment(a text.Alignment) Wrapper {
	return func(w W) W {
		return func(c C) D {
			old := Theme.TextAlignment
			Theme.TextAlignment = a
			d := w(c)
			Theme.TextAlignment = old
			return d
		}
	}
}

func TextColor(col color.NRGBA) Wrapper {
	return func(w W) W {
		return func(c C) D {
			old := Theme.TextColor
			Theme.TextColor = col
			d := w(c)
			Theme.TextColor = old
			return d
		}
	}
}
