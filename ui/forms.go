package ui

import (
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
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
	e.HintColor = FgColorMuted

	return Background(BgColor,
		Wrap(
			// material.Editor(TH, editor, hint).Layout,
			e.Layout,
			border,
			RoundedCorners,
			Inset1,
		),
	)
}

// returns a clickable button with an icon and a title
func ToolbarButton(clickable *Clickable, icon W, title string, onclick func(), ctx layout.Context) W {
	if clickable.Clicked(ctx) {
		onclick()
	}
	// bg := BgColorMuted
	bg := BgColor
	hovered := clickable.Hovered()
	if hovered {
		// bg = BgColorEmphasis
		bg = BgColorMuted
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

	// if hovered {
	// 	btn = Tooltip(btn, desc)
	// }

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
