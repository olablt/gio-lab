package main

import (
	"image"
	"log"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui"
)

type MyApp struct {
	Inset layout.Inset
}

// new MyApp
func NewMyApp() *MyApp {
	a := &MyApp{
		Inset: layout.UniformInset(12),
	}
	return a
}

// LayoutMainWindow
func (a *MyApp) LayoutMainWindow(gtx C, th *material.Theme) layout.Dimensions {
	return a.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical, WeightSum: 3}.Layout(gtx,
			// row 1
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.H1(th, "Hello, Gio")
				return l.Layout(gtx)
			}),
			// row
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.Label(th, th.TextSize*14.0/16.0, "Thin weight label")
				l.Font.Weight = font.Thin
				return l.Layout(gtx)
			}),
			// row
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.Label(th, th.TextSize*14.0/16.0, "Light weight label")
				l.Font.Weight = font.Light
				return l.Layout(gtx)
			}),
			// row
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.Label(th, th.TextSize*14.0/16.0, "Normal weight label")
				l.Font.Weight = font.Normal
				return l.Layout(gtx)
			}),
			// row
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.Label(th, th.TextSize*14.0/16.0, "Medium weight label")
				l.Font.Weight = font.Medium
				return l.Layout(gtx)
			}),
			// row
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.Label(th, th.TextSize*14.0/16.0, "Bold weight label")
				l.Font.Weight = font.Bold
				return l.Layout(gtx)
			}),
		)
	})
}

// Layout
func (a *MyApp) Layout(gtx C, th *material.Theme) layout.Dimensions {
	// return widget.Label{}.Layout(gtx, "Hello, Gio!")
	// l := material.H1(th, "Hello, Gio")
	// return l.Layout(gtx)
	log.Printf("main gtx %+v", gtx.Constraints)

	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return a.LayoutMainWindow(gtx, th)
			// defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
			// paint.Fill(gtx.Ops, ui.ColorBg)
			// return layout.Dimensions{Size: gtx.Constraints.Min}
		}, func(gtx layout.Context) layout.Dimensions {
			log.Printf("modal gtx %+v", gtx.Constraints)
			// return ui.ColorBox(gtx, image.Pt(500, 300), ui.ColorBgAccent)
			// func FillWithLabel(gtx layout.Context, th *material.Theme, text string, fg, bg color.NRGBA) layout.Dimensions {
			size := layout.Dimensions{Size: image.Pt(300, 100)}
			gtx.Constraints.Min = size.Size
			// gtx.Constraints.Max = size.Size
			ui.FillWithLabel(gtx, *th, "Modal", ui.ColorFg, ui.ColorBgAccent)
			return size
			// return layout.Dimensions{Size: image.Pt(300, 100)}
		})

}
