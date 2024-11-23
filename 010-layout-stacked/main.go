package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op/clip"
	"gioui.org/op/paint"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type MyApp struct {
	window *app.Window
}

func newMyApp() *MyApp {
	a := &MyApp{}
	a.createWindow()
	return a
}

func (a *MyApp) createWindow() {
	// w := app.NewWindow()
	w := new(app.Window)
	w.Option(
		app.Title("Gio"),
		app.Size(unit.Dp(600), unit.Dp(400)),
	)
	a.window = w
}

func (a *MyApp) loop(w *app.Window) error {
	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			log.Println("[INFO] DestroyEvent")
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			a.layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

func (a *MyApp) layout(gtx layout.Context) {
	// ColorBox(gtx, gtx.Constraints.Max, background)
	// ColorBox(gtx, gtx.Constraints.Min, background)
	// inset(gtx)
	stacked(gtx)
}

func inset(gtx layout.Context) layout.Dimensions {
	// Draw rectangles inside of each other, with 30dp padding.
	return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return ColorBox(gtx, gtx.Constraints.Max, red)
	})
}

func stacked(gtx layout.Context) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		// Force widget to the same size as the second.
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			// This will have a minimum constraint of 100x100.
			return ColorBox(gtx, gtx.Constraints.Min, red)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(100, 30), green)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return ColorBox(gtx, image.Pt(30, 100), blue)
		}),
	)
}

func main() {
	myapp := newMyApp()
	if err := myapp.loop(myapp.window); err != nil {
		log.Println("[ERROR] run failed.", err)
		os.Exit(1)
	} else {
		log.Println("[INFO] run success.")
	}

	os.Exit(0)
}

// helpers

// Test colors.
var (
	background = color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0xFF}
	// background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red   = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue  = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}
