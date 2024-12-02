// adapted material.Button to ActionListItem
package cpalette

import (
	"image"
	"image/color"
	"math"

	"gioui.org/font"
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui/f32color"
)

type ActionListItemStyle struct {
	Text  string
	RText string
	// Color is the text color.
	Color          color.NRGBA
	Font           font.Font
	TextSize       unit.Sp
	Background     color.NRGBA
	CornerRadius   unit.Dp
	Inset          layout.Inset
	ActionListItem *widget.Clickable
	shaper         *text.Shaper
}

type ActionListItemLayoutStyle struct {
	Background     color.NRGBA
	CornerRadius   unit.Dp
	ActionListItem *widget.Clickable
}

type IconActionListItemStyle struct {
	Background color.NRGBA
	// Color is the icon color.
	Color color.NRGBA
	Icon  *widget.Icon
	// Size is the icon size.
	Size           unit.Dp
	Inset          layout.Inset
	ActionListItem *widget.Clickable
	Description    string
}

func ActionListItem(th *material.Theme, button *widget.Clickable, txt, rtxt string) ActionListItemStyle {
	b := ActionListItemStyle{
		Text:         txt,
		RText:        rtxt,
		Color:        th.Palette.Fg,
		CornerRadius: 4,
		Background:   th.Palette.Bg,
		TextSize:     th.TextSize * 14.0 / 16.0,
		Inset: layout.Inset{
			Top: 5, Bottom: 5,
			Left: 12, Right: 12,
		},
		ActionListItem: button,
		shaper:         th.Shaper,
	}
	b.Font.Typeface = th.Face
	return b
}

func ActionListItemLayout(th *material.Theme, button *widget.Clickable) ActionListItemLayoutStyle {
	return ActionListItemLayoutStyle{
		ActionListItem: button,
		Background:     th.Palette.Bg,
		CornerRadius:   4,
	}
}

func IconActionListItem(th *material.Theme, button *widget.Clickable, icon *widget.Icon, description string) IconActionListItemStyle {
	return IconActionListItemStyle{
		Background:     th.Palette.Bg,
		Color:          th.Palette.Fg,
		Icon:           icon,
		Size:           24,
		Inset:          layout.UniformInset(12),
		ActionListItem: button,
		Description:    description,
	}
}

// Clickable lays out a rectangular clickable widget without further
// decoration.
func Clickable(gtx layout.Context, button *widget.Clickable, w layout.Widget) layout.Dimensions {
	return button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
				if button.Hovered() || gtx.Focused(button) {
					paint.Fill(gtx.Ops, f32color.Hovered(color.NRGBA{}))
				}
				for _, c := range button.History() {
					drawInk(gtx, c)
				}
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			w,
		)
	})
}

func (b ActionListItemStyle) Layout(gtx layout.Context) layout.Dimensions {
	return ActionListItemLayoutStyle{
		Background:     b.Background,
		CornerRadius:   b.CornerRadius,
		ActionListItem: b.ActionListItem,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		weight := font.Normal
		return b.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{WeightSum: 2}.Layout(gtx,
				// command
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					colMacro := op.Record(gtx.Ops)
					paint.ColorOp{Color: b.Color}.Add(gtx.Ops)
					b.Font.Weight = weight
					return widget.Label{Alignment: text.Start}.Layout(gtx, b.shaper, b.Font, b.TextSize, b.Text, colMacro.Stop())
				}),
				// shortcut
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					colMacro := op.Record(gtx.Ops)
					paint.ColorOp{Color: b.Color}.Add(gtx.Ops)
					b.Font.Weight = weight
					return widget.Label{Alignment: text.End}.Layout(gtx, b.shaper, b.Font, b.TextSize, b.RText, colMacro.Stop())
				}),
			)
		})

	})
}

func (b ActionListItemLayoutStyle) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	min := gtx.Constraints.Min
	return b.ActionListItem.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rr := gtx.Dp(b.CornerRadius)
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, rr).Push(gtx.Ops).Pop()
				background := b.Background
				switch {
				case !gtx.Enabled():
					background = f32color.Disabled(b.Background)
				case b.ActionListItem.Hovered() || gtx.Focused(b.ActionListItem):
					background = f32color.Hovered(b.Background)
				}
				paint.Fill(gtx.Ops, background)
				for _, c := range b.ActionListItem.History() {
					drawInk(gtx, c)
				}
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min = min
				return layout.W.Layout(gtx, w)
			},
		)
	})
}

func (b IconActionListItemStyle) Layout(gtx layout.Context) layout.Dimensions {
	m := op.Record(gtx.Ops)
	dims := b.ActionListItem.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		if d := b.Description; d != "" {
			semantic.DescriptionOp(b.Description).Add(gtx.Ops)
		}
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rr := (gtx.Constraints.Min.X + gtx.Constraints.Min.Y) / 4
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, rr).Push(gtx.Ops).Pop()
				background := b.Background
				switch {
				case !gtx.Enabled():
					background = f32color.Disabled(b.Background)
				case b.ActionListItem.Hovered() || gtx.Focused(b.ActionListItem):
					background = f32color.Hovered(b.Background)
				}
				paint.Fill(gtx.Ops, background)
				for _, c := range b.ActionListItem.History() {
					drawInk(gtx, c)
				}
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			func(gtx layout.Context) layout.Dimensions {
				return b.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					size := gtx.Dp(b.Size)
					if b.Icon != nil {
						gtx.Constraints.Min = image.Point{X: size}
						b.Icon.Layout(gtx, b.Color)
					}
					return layout.Dimensions{
						Size: image.Point{X: size, Y: size},
					}
				})
			},
		)
	})
	c := m.Stop()
	bounds := image.Rectangle{Max: dims.Size}
	defer clip.Ellipse(bounds).Push(gtx.Ops).Pop()
	c.Add(gtx.Ops)
	return dims
}

func drawInk(gtx layout.Context, c widget.Press) {
	// duration is the number of seconds for the
	// completed animation: expand while fading in, then
	// out.
	const (
		expandDuration = float32(0.5)
		fadeDuration   = float32(0.9)
	)

	now := gtx.Now

	t := float32(now.Sub(c.Start).Seconds())

	end := c.End
	if end.IsZero() {
		// If the press hasn't ended, don't fade-out.
		end = now
	}

	endt := float32(end.Sub(c.Start).Seconds())

	// Compute the fade-in/out position in [0;1].
	var alphat float32
	{
		var haste float32
		if c.Cancelled {
			// If the press was cancelled before the inkwell
			// was fully faded in, fast forward the animation
			// to match the fade-out.
			if h := 0.5 - endt/fadeDuration; h > 0 {
				haste = h
			}
		}
		// Fade in.
		half1 := t/fadeDuration + haste
		if half1 > 0.5 {
			half1 = 0.5
		}

		// Fade out.
		half2 := float32(now.Sub(end).Seconds())
		half2 /= fadeDuration
		half2 += haste
		if half2 > 0.5 {
			// Too old.
			return
		}

		alphat = half1 + half2
	}

	// Compute the expand position in [0;1].
	sizet := t
	if c.Cancelled {
		// Freeze expansion of cancelled presses.
		sizet = endt
	}
	sizet /= expandDuration

	// Animate only ended presses, and presses that are fading in.
	if !c.End.IsZero() || sizet <= 1.0 {
		gtx.Execute(op.InvalidateCmd{})
	}

	if sizet > 1.0 {
		sizet = 1.0
	}

	if alphat > .5 {
		// Start fadeout after half the animation.
		alphat = 1.0 - alphat
	}
	// Twice the speed to attain fully faded in at 0.5.
	t2 := alphat * 2
	// Beziér ease-in curve.
	alphaBezier := t2 * t2 * (3.0 - 2.0*t2)
	sizeBezier := sizet * sizet * (3.0 - 2.0*sizet)
	size := gtx.Constraints.Min.X
	if h := gtx.Constraints.Min.Y; h > size {
		size = h
	}
	// Cover the entire constraints min rectangle and
	// apply curve values to size and color.
	size = int(float32(size) * 2 * float32(math.Sqrt(2)) * sizeBezier)
	alpha := 0.7 * alphaBezier
	const col = 0.8
	ba, bc := byte(alpha*0xff), byte(col*0xff)
	rgba := f32color.MulAlpha(color.NRGBA{A: 0xff, R: bc, G: bc, B: bc}, ba)
	ink := paint.ColorOp{Color: rgba}
	ink.Add(gtx.Ops)
	rr := size / 2
	defer op.Offset(c.Position.Add(image.Point{
		X: -rr,
		Y: -rr,
	})).Push(gtx.Ops).Pop()
	defer clip.UniformRRect(image.Rectangle{Max: image.Pt(size, size)}, rr).Push(gtx.Ops).Pop()
	paint.PaintOp{}.Add(gtx.Ops)
}
