package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type MyApp struct {
	window *app.Window
	theme  *material.Theme
	list   *widget.List
	image  *widget.Image
}

func newMyApp() *MyApp {
	a := &MyApp{}
	a.createWindow()
	a.list = &widget.List{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}
	// theme
	a.theme = material.NewTheme()
	a.theme.Palette = material.Palette{
		Fg:         rgb(0xffffff),
		Bg:         rgb(0x000000),
		ContrastFg: rgb(0x3f51b5),
		ContrastBg: rgb(0xffffff),
	}
	a.theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	// laod image from jpg file "image1.jpg"
	a.loadImage("assets/image1.jpg")

	return a
}

func (a *MyApp) loadImage(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}
	a.image = &widget.Image{
		Fit: widget.Contain,
		Src: paint.NewImageOp(img),
	}
}

func (a *MyApp) createWindow() {
	w := new(app.Window)
	w.Option(
		app.Title("Gio"),
		app.Size(unit.Dp(400), unit.Dp(600)),
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
	// black background
	ColorBox(gtx, gtx.Constraints.Max, background)
	// widgets list
	a.listing(gtx)
}

type (
	D = layout.Dimensions
	C = layout.Context
)

func (a *MyApp) listing(gtx layout.Context) layout.Dimensions {

	widgets := []layout.Widget{
		// Header - first widget
		func(gtx C) D {
			l := material.H4(a.theme, "Hello, Gio!")
			// l.State = topLabelState
			return l.Layout(gtx)
		},
		// Image - second widget
		func(gtx C) D {
			recAlpha := op.Record(gtx.Ops)
			imDims := a.image.Layout(gtx) // store the image dimensions
			alpha := recAlpha.Stop()

			// clip the image to rounded corners
			defer clip.UniformRRect(image.Rectangle{Max: imDims.Size}, 25).Push(gtx.Ops).Pop()
			// overlay the image
			alpha.Add(gtx.Ops)
			// overlay the semi transparent rectangle at the bottom
			rectHeight := int(float64(imDims.Size.X) * 0.1)
			offset := op.Offset(image.Pt(0, imDims.Size.Y-rectHeight)).Push(gtx.Ops)
			ColorBox(gtx, image.Pt(imDims.Size.X, rectHeight), color.NRGBA{R: 0x20, G: 0x20, B: 0x20, A: 200})

			// add the text
			// Add text inside the rectangle
			textInset := layout.Inset{
				Top:   unit.Dp(float32(rectHeight) * 0.1),
				Left:  unit.Dp(10),
				Right: unit.Dp(10),
			}
			textSize := unit.Sp(float32(rectHeight) * 0.3) // Text size proportional to rectangle height
			label := material.Label(a.theme, textSize, "Vartai")
			label.Color = color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF} // White text
			textInset.Layout(gtx, func(gtx C) D {
				return label.Layout(gtx)
			})

			offset.Pop()
			return layout.Dimensions{Size: imDims.Size}

			// radius := unit.Dp(4) // Set the radius for the rounded corners. Adjust this value as needed.
			// rect := image.Rectangle{Max: imDims.Size}
			// rr := int(radius) // Convert radius to pixels.
			// // Define the clip area with rounded corners.
			// clip.RRect{
			// 	Rect: rect,
			// 	NE:   rr, NW: rr, SE: rr, SW: rr, // Apply the same radius to all corners.
			// }.Op(gtx.Ops)
			// // Draw the image within the clipped area.
			// paint.PaintOp{}.Add(gtx.Ops)
			// // Return the original dimensions of the image.
			// // The image will be clipped to rounded corners, but retains its original size.
			// return layout.Dimensions{Size: imDims.Size}
		},
		// third widget
		func(gtx C) D {
			col := color.NRGBA{R: byte((1 + 5) * 20), G: 0x20, B: 0x20, A: 0xFF}
			size := image.Pt(gtx.Constraints.Max.X, 300)
			defer clip.UniformRRect(image.Rectangle{Max: size}, 25).Push(gtx.Ops).Pop()
			return ColorBox(gtx, size, col)
		},
		func(gtx C) D {
			col := color.NRGBA{R: byte((1 + 5) * 20), G: 0x20, B: 0x20, A: 0xFF}
			return ColorBox(gtx, image.Pt(gtx.Constraints.Max.X, 300), col)
		},
		func(gtx C) D {
			col := color.NRGBA{R: byte((1 + 5) * 20), G: 0x20, B: 0x20, A: 0xFF}
			return ColorBox(gtx, image.Pt(gtx.Constraints.Max.X, 300), col)
		},
	}

	in := layout.Inset{
		Top:    unit.Dp(0),
		Right:  unit.Dp(0),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(5),
	}
	return material.List(a.theme, a.list).Layout(gtx, len(widgets), func(gtx C, i int) D {
		return in.Layout(gtx, widgets[i])
	})

	// return material.List(a.theme, a.list).Layout(gtx, 10, func(gtx C, i int) D {
	// // demo random widgets
	// col := color.NRGBA{R: byte((i + 5) * 20), G: 0x20, B: 0x20, A: 0xFF}
	// return in.Layout(gtx, func(gtx C) D {
	// 	return ColorBox(gtx, image.Pt(gtx.Constraints.Max.X, 300), col)
	// })
	// })
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
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	// defer clip.UniformRRect(image.Rectangle{Max: size}, int(unit.Dp(15))).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}
