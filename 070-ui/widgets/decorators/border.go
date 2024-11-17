package decorators

import (
	"gioui.org/widget"
	"github.com/yourusername/070-ui/theme"
	"github.com/yourusername/070-ui/widgets"
)

type Border struct {
	widget   widgets.Widget
	color    color.NRGBA
	width    float32
	radius   float32
}

func NewBorder(w widgets.Widget, theme *theme.Theme) *Border {
	return &Border{
		widget: w,
		color:  theme.Colors.Border,
		width:  1,
		radius: 0,
	}
}

func (b *Border) WithColor(c color.NRGBA) *Border {
	b.color = c
	return b
}

func (b *Border) WithWidth(w float32) *Border {
	b.width = w
	return b
}

func (b *Border) WithRadius(r float32) *Border {
	b.radius = r
	return b
}

func (b *Border) Layout(gtx widgets.Context) widgets.Dimensions {
	return widget.Border{
		Color:        b.color,
		Width:        b.width,
		CornerRadius: b.radius,
	}.Layout(gtx, b.widget.Layout)
}
