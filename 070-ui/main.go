package main

import (
	"image"
	"image/color"
	"log"
	"os"

	// "gio.tools/icons"

	// "golang.org/x/exp/shiny/materialdesign/icons"
	"golang.org/x/exp/shiny/materialdesign/icons"

	// "gio.tools/icons"
	"gioui.org/app"
	gio "gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui"
)

type App struct {
	// columnWidgets []widget.Clickable
	Clickables ui.Clickables
	Editors    ui.Editors
}

func main() {
	myApp := &App{
		Clickables: ui.NewClickables(),
		Editors:    ui.NewEditors(),
	}
	// myApp.Option
	ui.TH.Palette = material.Palette{
		Fg:         ui.FgColor,
		Bg:         ui.BgColor,
		ContrastBg: ui.BgColorAccent,
		ContrastFg: ui.FgColor,
	}

	go func() {
		w := new(gio.Window)
		w.Option(
			// app.Title("oGio"),
			app.Size(unit.Dp(1920)/1.5, unit.Dp(1080/2)),
		)
		if err := myApp.RunApp(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func (a *App) RunApp(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	// aState := appState{
	// 	columnWidgets: make([]widget.Clickable, 6),
	// }

	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// Reset the operations
			ops.Reset()

			ctx := app.NewContext(&ops, e)

			cl := clip.Rect{Max: ctx.Constraints.Max}.Push(ctx.Ops)
			paint.Fill(ctx.Ops, ui.BgColor)
			cl.Pop()

			// APP Layout
			ui.Rows(
				// Header
				ui.Rigid(a.AppHeader()),

				// ui.Rigid(ui.Inset1(ui.HR(1))),
				// ui.RowSpacer1,
				ui.Rigid(ui.HR(1)),
				ui.RowSpacer1,

				// Body
				ui.Flexed(1, a.AppBody(ctx)),

				// // Footer
				// ui.Rigid(ui.Label("Footer")),
			)(ctx)

			e.Frame(ctx.Ops)
		}
	}
}

func (a *App) AppHeader() ui.W {
	return ui.Inset1(ui.H1("Gio UI Demo"))
	// return ui.AlignMiddle(
	// 	ui.H1("Gio UI Demo"),
	// 	// ui.FontSize(22)(
	// 	// 	ui.Label("Gio UI Demo"), // one line text
	// 	// ),
	// )
}

func (a *App) AppBody(gtx ui.C) ui.W {
	_ = gtx
	del := func() { log.Println("Delete") }
	IconSettings := ui.IconSize(ui.Icon(icons.ActionSettings, ui.FgColorMuted))
	IconDelete := ui.IconSize(ui.Icon(icons.ContentClear, ui.FgColorDanger))
	IconDeveloperMode := ui.IconSize(ui.Icon(icons.DeviceDeveloperMode, ui.FgColorMuted))

	// gtx.Constraints.Min = image.Point{X: 200, Y: 200}
	// gtx.Constraints.Max = image.Point{X: 200, Y: 200}

	return ui.Columns(
		// Left Column - links
		ui.ColSpacer1,
		ui.Rigid(
			// ui.Flexed(1,
			// ColorBoxW(ui.P{500, 100}, ui.BgColor),

			ui.ConstraintW(250,
				ui.Rows(
					// ui.Flexed(1,
					ui.Rigid(
						ui.InvisibleButton(a.Clickables.Get("GenSettingsButton"), IconSettings, "General Settings", del, gtx),
						// ui.Wrap(
						// 	ui.Text("Left Panel"),
						// 	ui.AlignStart,
						// 	ui.TextColor(ui.FgColor),
						// 	ui.MaxLines(1),
						// ),
					),
					// ui.RowSpacer1,
					ui.Rigid(
						ui.InvisibleButton(a.Clickables.Get("DevSettingsButton"), IconDeveloperMode, "Developer Settings", del, gtx),
					),
					// Removable Item
					ui.Rigid(
						ui.RoundedCorners(
							// ui.Background(ui.BgColorMuted,
							ui.Background(ui.BgColor,
								ui.Inset1(
									ui.Columns(
										// ui.Rigid(ui.OnClick(a.Clickables.Get("GenSettings"), IconSettings, del, gtx)),
										// ui.Rigid(IconSettings),
										ui.RowSpacer1,
										ui.Flexed(1, ui.Label("Removable Item")),
										ui.Rigid(ui.OnClick(a.Clickables.Get("RemovableItemDelete"), IconDelete, del, gtx)),
									),
								),
							),
						),
					),
					// end Removable Item
				),
			),
		),

		ui.ColSpacer1,

		// // Separator
		// ui.Rigid(ui.Inset1(ui.VR(1))),

		// Right Column - content
		ui.Flexed(3,

			ui.Background(ui.BgColorBlack,
				ui.Inset1(
					ui.Rows(

						// row header
						ui.Rigid(
							ui.Inset1(
								ui.Wrap(
									ui.H1("Right Panel - Content"),
									ui.AlignStart,
								),
							),
						),

						ui.Rigid(ui.Columns(
							// row text input
							ui.Rigid(
								// TODO: add a constraint to the text input, this is not working
								ui.ConstraintW(150,
									ui.TextInput(a.Editors.Get("Text input"), "Text Input"),
								),
							),
							// col spacer
							ui.ColSpacer1,
							// col button
							ui.Rigid(
								ui.ConstraintW(150,
									ui.DefaultButton(a.Clickables.Get("Default Button"), IconSettings, "Default Button", del, gtx),
								),
							),
							// col spacer
							ui.ColSpacer1,
							// col button
							ui.Rigid(
								ui.ConstraintW(150,
									ui.InvisibleButton(a.Clickables.Get("Invisible Button"), IconSettings, "Invisible Button", del, gtx),
								),
							),
							// col spacer
							ui.ColSpacer1,
							// col button
							ui.Rigid(
								ui.ConstraintW(150,
									ui.PrimaryButton(a.Clickables.Get("Primary Button"), IconSettings, "Primary Button", del, gtx),
								),
							),
							// col spacer
							ui.ColSpacer1,
							// col button
							ui.Rigid(
								ui.ConstraintW(150,
									ui.DangerButton(a.Clickables.Get("Danger Button"), IconSettings, "Danger Button", del, gtx),
								),
							),
						)),

						// spacer
						ui.RowSpacer1,
						// row text input
						ui.Rigid(
							ui.ConstraintW(150,
								ui.TextInput(a.Editors.Get("Text input2"), "Text Input2"),
							),
						),
					),

					// ColorBoxW(ui.P{}, ui.BgColorTransparent),
				),
			),
		),
	)

	// return ui.Columns(
	// 	ui.Rigid(ColorBoxW(ui.P{100, 100}, ui.BgColorAccent)),
	// 	ui.Flexed(0.5, ColorBoxW(ui.P{}, ui.BgColorDanger)),
	// 	ui.Rigid(ColorBoxW(ui.P{100, 100}, ui.BgColorSevere)),
	// 	ui.Rigid(ui.WSpacer1),
	// 	// Flexed(0.5, ColorBoxW(P{}, green)),
	// 	ui.Flexed(0.5, ui.Rows(
	// 		ui.Rigid(
	// 			ui.RoundedCorners(
	// 				// Border(
	// 				ui.BorderActive(
	// 					ColorBoxW(ui.P{100, 100}, ui.FgColor),
	// 				),
	// 			),
	// 		),
	// 		// ui.Rigid(HSpacer1),
	// 		ui.Flexed(0.7, ColorBoxW(ui.P{}, getColor(1))),

	// 		ui.Rigid(
	// 			ui.Panel("Panel",
	// 				ColorBoxW(ui.P{100, 100}, getColor(2)),
	// 			),
	// 		),

	// 		ui.Flexed(0.3,
	// 			ColorBoxW(ui.P{}, getColor(3)),
	// 		),
	// 	)),
	// )

}

func getColor(i int) color.NRGBA {
	return color.NRGBA{R: uint8(100 + i*20), G: uint8(150 + i*15), B: uint8(200 - i*10), A: 255}
}

// ColorBox creates a widget with the specified dimensions and color.
func ColorBoxW(size image.Point, color color.NRGBA) ui.W {
	return func(c ui.C) ui.D {
		if size.X == 0 {
			size = c.Constraints.Min
		}
		defer clip.Rect{Max: size}.Push(c.Ops).Pop()
		paint.ColorOp{Color: color}.Add(c.Ops)
		paint.PaintOp{}.Add(c.Ops)
		return ui.D{Size: size}
	}
}
func ColorBox(c ui.C, size image.Point, color color.NRGBA) ui.D {
	defer clip.Rect{Max: size}.Push(c.Ops).Pop()
	paint.ColorOp{Color: color}.Add(c.Ops)
	paint.PaintOp{}.Add(c.Ops)
	return ui.D{Size: size}
}
