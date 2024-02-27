package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func Loop() {
	// th := material.NewTheme(gofont.Collection())
	go func() {
		w := app.NewWindow(app.Title("oGio"))

		th := material.NewTheme()
		// awsomeFaces, _ := LoadFontToCollection("assets/Font Awesome 5 Pro-Light-300.otf")
		awsomeFaces, _ := LoadFontToCollection("assets/Consolas Nerd Font.TTF")
		th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(awsomeFaces))
		// th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Regular()))
		// th.Fg = ui.Alpha(colorBlack, 90)
		// th.Bg = colorWhite

		var ops op.Ops
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				// panic(e.Err)
				log.Println("au revoir")
				os.Exit(0)
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				gtx.Metric = unit.Metric{
					// PxPerDp: 1.8,
					// PxPerSp: 1.8,
					PxPerDp: 4,
					PxPerSp: 4,
				}
				// layout
				Layout(gtx, th)
				// render the operation list
				e.Frame(gtx.Ops)
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
