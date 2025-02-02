package ui

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui/f32color"
)

type (
	Clickable = widget.Clickable
	Bool      = widget.Bool
	Editor    = widget.Editor
)

type Clickables map[string]*Clickable
type Bools map[string]*Bool
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

func NewBools() Bools {
	return map[string]*Bool{}
}

func (c Bools) Get(id string) *Bool {
	if btn, ok := c[id]; ok {
		return btn
	}

	btn := new(Bool)
	c[id] = btn

	return btn
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
	e.Color = FgColor
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

func TextInputWithKeys(editor *widget.Editor, hint string) W {
	return func(gtx layout.Context) layout.Dimensions {
		// Create an input area for key events
		// input := key.InputOp{Tag: editor, Keys: key.NameFilterPrintable}.Push(gtx.Ops)
		// defer input.Pop()

		// Layout the editor
		return material.Editor(TH, editor, hint).Layout(gtx)
	}
}

type ButtonStyle struct {
	Bg        color.NRGBA
	Fg        color.NRGBA
	BgH       color.NRGBA
	FgH       color.NRGBA
	Inset     DP
	Alignment layout.Alignment // layout.Start, layout.Middle, layout.End
	Font      font.Font
	Size      unit.Sp
}

func StyledButton(clickable *Clickable, title string, onclick func(), ctx layout.Context, style ButtonStyle) W {
	// log.Printf("styled button ctx %+v", ctx)
	if onclick != nil && clickable.Clicked(ctx) {
		onclick()
	}
	bg := style.Bg
	fg := style.Fg
	hovered := clickable.Hovered()
	if hovered {
		fg = style.FgH
		bg = style.BgH
		pointer.CursorPointer.Add(ctx.Ops) // set mouse cursor
	}

	// columns := func(alignment layout.Alignment, children ...layout.FlexChild) W {
	// 	return func(c C) D {
	// 		return layout.Flex{Axis: layout.Horizontal, Alignment: alignment}.Layout(c, children...)
	// 	}
	// }
	_ = fg
	_ = bg

	inset := LayoutToWrapper(layout.UniformInset(style.Inset).Layout)
	w := func(c C) D {
		// c.Constraints.Min.X = 0 // Allow natural width
		return clickable.Layout(c,
			Background(bg,
				inset(
					// columns(style.Alignment,
					// Font(style.Font)(FontSize(style.Size)(TextColor(fg)(Label(title)))),
					FontSize(style.Size)(
						TextColor(fg)(
							// Label(title),
							func(c C) D {
								// tl := widget.Label{Alignment: Theme.TextAlignment, MaxLines: 1}
								// return tl.Layout(c, Theme.FontFamily, style.Font, style.Size, title, op.CallOp{})

								tl := widget.Label{Alignment: Theme.TextAlignment, MaxLines: Theme.MaxLines}
								// tl.LineHeight = 1
								// tl.LineHeightScale = 1
								paint.ColorOp{Color: Theme.TextColor}.Add(c.Ops)
								return tl.Layout(c, Theme.FontFamily, font.Font{Weight: Theme.FontWeight}, Theme.FontSize, title, op.CallOp{})

							},
						),
					),
				),
			// ),
			),
		)
	}
	return w
}

func StyledLabel(title string, style ButtonStyle, ctx layout.Context) W {
	bg := style.Bg
	fg := style.Fg
	inset := LayoutToWrapper(layout.UniformInset(style.Inset).Layout)
	return func(c C) D {
		c.Constraints.Min.X = 0 // Allow natural width
		return Background(bg,
			inset(
				// FontSize(style.Size)(
				// TextColor(fg)(
				func(c C) D {
					tl := widget.Label{Alignment: Theme.TextAlignment, MaxLines: 1}
					paint.ColorOp{Color: fg}.Add(c.Ops)
					return tl.Layout(c, Theme.FontFamily, style.Font, style.Size, title, op.CallOp{})
				},
				// ),
				// ),
			),
		)(c)
	}
	// w := func(c C) D {
	// 	c.Constraints.Min.X = 0 // Allow natural width
	// 	return clickable.Layout(c,
	// 		Background(bg,
	// 			inset(
	// 				// columns(style.Alignment,
	// 				// Font(style.Font)(FontSize(style.Size)(TextColor(fg)(Label(title)))),
	// 				FontSize(style.Size)(
	// 					TextColor(fg)(
	// 						// Label(title),
	// 						func(c C) D {
	// 							// tl := widget.Label{Alignment: Theme.TextAlignment, MaxLines: 1}
	// 							// return tl.Layout(c, Theme.FontFamily, style.Font, style.Size, title, op.CallOp{})

	// 							tl := widget.Label{Alignment: Theme.TextAlignment, MaxLines: Theme.MaxLines}
	// 							// tl.LineHeight = 1
	// 							// tl.LineHeightScale = 1
	// 							paint.ColorOp{Color: Theme.TextColor}.Add(c.Ops)
	// 							return tl.Layout(c, Theme.FontFamily, font.Font{Weight: Theme.FontWeight}, Theme.FontSize, title, op.CallOp{})

	// 						},
	// 					),
	// 				),
	// 			),
	// 		),
	// 	)
	// }
	// return w
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

	if icon == nil {
		icon = EmptyWidget
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
		c.Constraints.Min.X = 0 // Allow natural width
		// ctx.Constraints.Max.X = 0 // Allow natural width
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
