package main

import (
	"image"

	"gioui.org/font"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui"
)

// type C = layout.Context
// type D = layout.Dimensions

type MyApp struct {
	Inset     layout.Inset
	showModal bool
}

// new MyApp
func NewMyApp() *MyApp {
	a := &MyApp{
		Inset:     layout.UniformInset(12),
		showModal: false,
	}
	return a
}

func (a *MyApp) HandleKeyEvents(gtx layout.Context) {
	// Handle keyboard shortcuts
	for _, e := range gtx.Events(a) {
		if e, ok := e.(key.Event); ok {
			if e.State == key.Press {
				switch e.Name {
				case key.NameEscape:
					a.showModal = false
				case "O":
					if e.Modifiers.Contain(key.ModCtrl) {
						a.showModal = true
					}
				}
			}
		}
	}
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
	// Handle keyboard events
	key.InputOp{Tag: a}.Add(gtx.Ops)
	a.HandleKeyEvents(gtx)

	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return a.LayoutMainWindow(gtx, th)
		},
		func(gtx layout.Context) layout.Dimensions {
			if !a.showModal {
				return layout.Dimensions{}
			}

			size := layout.Dimensions{Size: image.Pt(300, 100)}
			gtx.Constraints.Min = size.Size

			// Center the modal
			return layout.Center.Layout(gtx, func(gtx C) D {
				return ui.FillWithLabel(gtx, *th, "Press ESC to close", ui.ColorFg, ui.ColorBgAccent)
			})
		})

}
