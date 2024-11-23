package main

import (
	"image"
	"image/color"
	"log"

	"gioui.org/font"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui"
)

// type C = layout.Context
// type D = layout.Dimensions

type MyApp struct {
	Inset     layout.Inset
	showModal bool
	overlay   widget.Clickable
}

// new MyApp
func NewMyApp() *MyApp {
	a := &MyApp{
		Inset:     layout.UniformInset(12),
		showModal: true,
	}
	return a
}

func (a *MyApp) HandleKeyEvents(gtx layout.Context) {
	// Declare the tag.
	tag := a

	// Confine the area of interest to a gtx Max
	areaStack := clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops)

	// Handle keyboard shortcuts
	event.Op(gtx.Ops, tag)

	// New key event reading
	for {
		// Read the next event
		event, ok := gtx.Event(
			key.Filter{
				Required: key.ModCtrl,
				Name:     "O",
			},
			key.Filter{
				Name: "Q",
			},
			key.Filter{
				Name: "O",
			},
			key.Filter{
				Name: key.NameEscape,
			},
			key.Filter{
				Name: key.NameEnter,
			},
		)
		if !ok {
			break
		}
		// filter key events
		ev, ok := event.(key.Event)
		if !ok {
			continue
		}

		switch ev.Name {
		case key.NameEscape:
			// log
			log.Printf("got key.Escape")
			a.showModal = false
		case key.NameReturn:
			// log
			log.Printf("got key.Return")
			a.showModal = false
		case "O":
			if ev.Modifiers.Contain(key.ModCtrl) {
				log.Println("got Ctrl + key.O")
				a.showModal = true
			} else {
				log.Println("got key.O")
				a.showModal = true
			}
		default:
			// log
			log.Printf("got key.%v", ev.Name)
		}
		// }
	}

	areaStack.Pop()
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
	// key.InputOp{Tag: a}.Add(gtx.Ops)
	a.HandleKeyEvents(gtx)

	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return a.LayoutMainWindow(gtx, th)
		},
		func(gtx layout.Context) layout.Dimensions {
			if !a.showModal {
				return layout.Dimensions{}
			}

			// Handle overlay clicks
			if a.overlay.Clicked(gtx) {
				a.showModal = false
			}

			// Draw clickable overlay
			return a.overlay.Layout(gtx, func(gtx C) D {
				// Draw semi-transparent background
				defer clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops).Pop()
				paint.ColorOp{Color: color.NRGBA{A: 200}}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)

				// Draw centered modal
				return layout.Center.Layout(gtx, func(gtx C) D {
					size := layout.Dimensions{Size: image.Pt(300, 100)}
					gtx.Constraints.Min = size.Size
					return ui.FillWithLabel(gtx, *th, "Modal", th.Palette.ContrastFg, th.Palette.ContrastBg)
				})
			})

			// // First draw semi-transparent overlay covering whole screen
			// defer clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops).Pop()
			// paint.ColorOp{Color: color.NRGBA{A: 245}}.Add(gtx.Ops) // Alpha 200 for semi-transparency
			// paint.PaintOp{}.Add(gtx.Ops)

			// // ui.FillWithLabel(gtx, *th, "Modal", ui.ColorFg, ui.ColorBgAccent)
			// // size := layout.Dimensions{Size: image.Pt(300, 100)}
			// // gtx.Constraints.Min = size.Size
			// // return size

			// // Then draw centered modal
			// return layout.Center.Layout(gtx, func(gtx C) D {
			// 	size := layout.Dimensions{Size: image.Pt(300, 100)}
			// 	gtx.Constraints.Min = size.Size
			// 	return ui.FillWithLabel(gtx, *th, "Modal", th.Palette.ContrastFg, th.Palette.ContrastBg)
			// })

		})

}
