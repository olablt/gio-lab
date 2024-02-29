package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

var (
	colorWhite = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorGrey  = color.NRGBA{R: 0x55, G: 0x55, B: 0x55, A: 0xff}
	colorBlack = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	colorBlue  = color.NRGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
	colorGreen = color.NRGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
)

// Loop is a helper function that runs the app event loop
func Loop(fn func(win *app.Window, gtx layout.Context, th *material.Theme)) {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Collection()))
	// awsomeFaces, _ := LoadFontToCollection("assets/Consolas Nerd Font.TTF")
	// th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(awsomeFaces))
	go func() {
		w := app.NewWindow(
			app.Title("oGio"),
			app.Size(unit.Dp(1920/4), unit.Dp(1080/2)),
			// app.Size(unit.Dp(1920/2), unit.Dp(1080/2)),
		)
		// w.Option(
		// 	app.Title("Gio"),
		// 	app.Size(unit.Dp(400), unit.Dp(600)),
		// )

		// ops will be used to encode different operations.
		var ops op.Ops

		// new event queue
		for {
			switch e := w.NextEvent().(type) {
			case app.FrameEvent:
				// gtx is used to pass around rendering and event information.
				gtx := app.NewContext(&ops, e)
				// render contents
				fn(w, gtx, th)
				// render frame
				e.Frame(gtx.Ops)
			case app.DestroyEvent:
				if e.Err != nil {
					log.Println("got error", e.Err)
					os.Exit(1)
				}
				log.Println("exiting...")
				os.Exit(0)
			case app.StageEvent:
				log.Printf("got stage event %#+v", e.Stage.String())
			}
		}

	}()
	app.Main()
}

// ColorBox creates a widget with the specified dimensions and color.
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func FillWithLabel(gtx layout.Context, th *material.Theme, text string, fg, bg color.NRGBA) layout.Dimensions {
	th.Palette.Fg = fg
	th.Palette.Bg = bg
	ColorBox(gtx, gtx.Constraints.Max, bg)
	return layout.Center.Layout(gtx, material.Label(th, unit.Sp(10), text).Layout)
}

func LoadFontToCollection(filename string) ([]font.FontFace, error) {
	// materialTheme := material.NewTheme(gofont.Collection())
	// load Awesome font
	fontData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error loading font file:", err)
	}
	awsomeFaces, err := opentype.ParseCollection(fontData)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	return awsomeFaces, nil
	// // merge go font and awsome font
	// faces := []font.FontFace{}
	// faces = append(gofont.Collection(), awsomeFaces...)
	// // for i, face := range faces {
	// // 	log.Println(i, "face", face)
	// // }
	// materialTheme := material.NewTheme()
	// materialTheme.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(faces))
	// // materialTheme.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Collection()))

	// // set the theme
	// // c.Theme = materialTheme
}
