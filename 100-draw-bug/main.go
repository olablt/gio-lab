package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {

	var path clip.Path
	// ui.Layout(func(gtx layout.Context) layout.Dimensions {
	Loop(func(win *app.Window, gtx layout.Context) {
		log.Println("opa")

		// large triangle
		path.Begin(gtx.Ops)
		path.MoveTo(f32.Pt(10, 10))
		path.LineTo(f32.Pt(float32(gtx.Constraints.Max.X)-50, 100))
		path.LineTo(f32.Pt(200, 500))
		path.Close()
		paint.FillShape(gtx.Ops, color.NRGBA{0, 255, 0, 255}, clip.Outline{Path: path.End()}.Op()) // fill polygon

		// small triangle
		path.Begin(gtx.Ops)
		path.MoveTo(f32.Pt(10, 10))
		path.LineTo(f32.Pt(100, 50))
		path.LineTo(f32.Pt(20, 100))
		path.Close()
		paint.FillShape(gtx.Ops, color.NRGBA{255, 0, 0, 255}, clip.Outline{Path: path.End()}.Op()) // fill polygon

		// log.Println(gtx.Constraints.Max)
	})

}

func Loop(fn func(win *app.Window, gtx layout.Context)) {
	go func() {
		w := app.Window{}
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
				// render contents
				fn(&w, gtx)
				// render frame
				e.Frame(gtx.Ops)
			case app.DestroyEvent:
				if e.Err != nil {
					log.Println("got error", e.Err)
					os.Exit(1)
				}
				log.Println("exiting...")
				os.Exit(0)
			// case app.StageEvent:
			case app.ConfigEvent:
				log.Printf("got config event Focused:%v", e.Config.Focused)
			}
		}

	}()
	app.Main()
}
