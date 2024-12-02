package ui

import (
	"image/color"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui/f32color"
)

type (
	Clickable = widget.Clickable
	Editor    = widget.Editor
)

type Clickables map[string]*Clickable
type Editors map[string]*Editor

func NewClickables() Clickables {
	return map[string]*Clickable{}
}

func (c Clickables) Get(id string) *Clickable {
	if btn, ok := c[id]; ok {
		return btn
	}

	btn := new(Clickable)
	c[id] = btn

	return btn
}

func OnClick(btn *Clickable, w W, onclick func(), ctx layout.Context) W {
	if btn.Clicked(ctx) {
		onclick()
	}

	return func(c C) D {
		return btn.Layout(c, w)
	}

}

func NewEditors() Editors {
	return map[string]*Editor{}
}

func (e Editors) Get(id string) *Editor {
	if editor, ok := e[id]; ok {
		return editor
	}

	// editor := new(Editor)
	editor := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}

	e[id] = editor

	return editor
}

func TextInput(editor *widget.Editor, hint string) W {
	border := Border

	// if editor.Focused() {
	// 	border = BorderActive
	// }

	e := material.Editor(TH, editor, hint)
	e.TextSize = Theme.FontSize
	// e.LineHeight = SP(1)
	// tl.LineHeight = 1
	// e.LineHeightScale = 0.5

	e.HintColor = FgColorMuted

	// w := func(c C) D {
	return RoundedCorners(
		Background(BgColorDisabled,
			Wrap(
				e.Layout,
				border,
				Inset1,
				// Inset05,
			),
		),
	)
}

func DefaultButton(clickable *Clickable, icon W, title string, onclick func(), ctx layout.Context) W {
	// bg := BgColor
	bg := BgColorDisabled
	fg := FgColor
	fgH := FgColor
	// bgH := BgColorDisabled
	bgH := f32color.Hovered(bg)
	return IconButton(clickable, icon, title, onclick, ctx, fg, bg, fgH, bgH)
}

func InvisibleButton(clickable *Clickable, icon W, title string, onclick func(), ctx layout.Context) W {
	bg := BgColor
	fg := FgColor
	fgH := FgColor
	bgH := BgColorDisabled
	// bgH := f32color.Hovered(bg)
	return IconButton(clickable, icon, title, onclick, ctx, fg, bg, fgH, bgH)
}

func PrimaryButton(clickable *Clickable, icon W, title string, onclick func(), ctx layout.Context) W {
	bg := BgColorSuccess
	fg := FgColor
	fgH := FgColor
	// bgH := BgColorSuccess
	bgH := f32color.Hovered(bg)
	return IconButton(clickable, icon, title, onclick, ctx, fg, bg, fgH, bgH)
}

func DangerButton(clickable *Clickable, icon W, title string, onclick func(), ctx layout.Context) W {
	bg := BgColorDisabled
	fg := FgColor
	fgH := FgColor
	bgH := BgColorDanger
	// return IconButton(clickable, EmptyWidget, title, onclick, ctx, fg, bg, fgH, bgH)
	return IconButton(clickable, icon, title, onclick, ctx, fg, bg, fgH, bgH)
}

// returns a clickable button with an icon and a title
func IconButton(clickable *Clickable, icon W, title string, onclick func(), ctx layout.Context, fg, bg, fgH, bgH color.NRGBA) W {
	if clickable.Clicked(ctx) {
		onclick()
	}
	// // bg := BgColorMuted
	// bg := BgColor
	hovered := clickable.Hovered()
	if hovered {
		// // bg = BgColorEmphasis
		// bg = BgColorMuted
		fg = fgH
		bg = bgH
		// pointer.CursorGrab.Add(ctx.Ops) // set mouse cursor
		pointer.CursorPointer.Add(ctx.Ops) // set mouse cursor
	}
	w := func(c C) D {
		return clickable.Layout(c,
			RoundedCorners(
				// ui.Background(ui.BgColorMuted,
				Background(bg,
					Inset1(
						ColumnsVCentered(
							Rigid(icon),
							ColSpacer1,
							Flexed(1, Label(title)),
						),
					),
				),
			),
		)
	}
	return w
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
