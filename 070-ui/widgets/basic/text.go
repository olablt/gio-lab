package basic

import (
	"gioui.org/widget"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/yourusername/070-ui/theme"
	"github.com/yourusername/070-ui/widgets"
)

type Text struct {
	text    string
	style   theme.TextStyle
	label   widget.Label
}

func NewText(text string, style theme.TextStyle) *Text {
	return &Text{
		text:  text,
		style: style,
		label: widget.Label{
			Alignment: style.Alignment,
			MaxLines:  style.MaxLines,
		},
	}
}

func (t *Text) Layout(gtx widgets.Context) widgets.Dimensions {
	paint.ColorOp{Color: t.style.Color}.Add(gtx.Ops)
	return t.label.Layout(
		gtx,
		t.style.Shaper,
		t.style.Font,
		t.style.Size,
		t.text,
		op.CallOp{},
	)
}
