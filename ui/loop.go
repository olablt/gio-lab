package ui

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
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

func Loop(fn func(win *app.Window, gtx layout.Context, th *material.Theme)) {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(LoadFontCollection()))
	// th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Collection()))
	// set Github theme
	th.Palette.Fg = ColorFg
	th.Palette.Bg = ColorBg
	th.Palette.ContrastFg = ColorFg
	th.Palette.ContrastBg = ColorBgAccent

	go func() {
		w := &app.Window{}
		w.Option(
			app.Title("oGio"),
			app.Size(unit.Dp(1920/4), unit.Dp(1080/2)),
		)

		// ops will be used to encode different operations.
		var ops op.Ops

		// new event queue
		for {
			switch e := w.Event().(type) {
			case app.FrameEvent:
				// gtx is used to pass around rendering and event information.
				gtx := app.NewContext(&ops, e)
				// fill the entire window with the background color
				defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
				paint.Fill(gtx.Ops, th.Palette.Bg)
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
			case app.ConfigEvent:
				log.Printf("got config event Focused:%v", e.Config.Focused)
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

func FillWithLabel(gtx layout.Context, th material.Theme, text string, fg, bg color.NRGBA) layout.Dimensions {
	th.Palette.Fg = fg
	th.Palette.Bg = bg
	ColorBox(gtx, gtx.Constraints.Min, bg)
	// return layout.Center.Layout(gtx, material.Label(&th, unit.Sp(10), text).Layout)
	return layout.Center.Layout(gtx, material.Body1(&th, text).Layout)
}

// FillWithLabelH3 creates a label with the specified text and background color
func FillWithLabelH3(gtx layout.Context, th *material.Theme, text string, backgroundColor color.NRGBA) layout.Dimensions {
	ColorBox(gtx, gtx.Constraints.Max, backgroundColor)
	return layout.Center.Layout(gtx, material.H3(th, text).Layout)
}
