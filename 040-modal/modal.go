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
			// Handle overlay clicks
			if a.overlay.Clicked(gtx) {
				log.Println("overlay clicked")
				a.showModal = false
			}

			if !a.showModal {
				return layout.Dimensions{}
			}

			// Create a modal clickable
			modalClick := &widget.Clickable{}

			// Draw semi-transparent overlay with click handler
			a.overlay.Layout(gtx, func(gtx C) D {
				defer clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops).Pop()
				paint.ColorOp{Color: color.NRGBA{A: 200}}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				return layout.Dimensions{Size: gtx.Constraints.Max}
			})

			// Draw modal on top with its own click handler
			return layout.Center.Layout(gtx, func(gtx C) D {
				size := layout.Dimensions{Size: image.Pt(300, 100)}
				gtx.Constraints.Min = size.Size

				return modalClick.Layout(gtx, func(gtx C) D {
					// Create a separate clip area for modal
					defer clip.Rect{Max: size.Size}.Push(gtx.Ops).Pop()
					// pointer.CursorPointer.Add(gtx.Ops) // Add pointer cursor for modal
					return ui.FillWithLabel(gtx, *th, "Modal", th.Palette.ContrastFg, th.Palette.ContrastBg)
				})
			})

		})

}
