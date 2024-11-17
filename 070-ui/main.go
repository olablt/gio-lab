package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type appState struct {
	columnWidgets []widget.Clickable
	showModal     bool
}

func main() {
	go func() {
		w := new(app.Window)
		if err := runApp(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func runApp(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	state := &appState{
		columnWidgets: make([]widget.Clickable, 6),
		showModal:     false,
	}

	var ops op.Ops
	
	// Set up key filter for the events we want to handle
	keyFilter := key.Filter{
		Required: key.ModCtrl,
		Optional: key.ModShift,
	}

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case key.Event:
			switch {
			case e.State == key.Press && e.Name == key.NameEscape:
				log.Println("ESC pressed")
				state.showModal = false
			case e.State == key.Press && e.Name == "O" && e.Modifiers.Contain(key.ModCtrl):
				log.Println("Ctrl+O pressed")
				state.showModal = true
			}
		case app.FrameEvent:
			// Reset the operations
			ops.Reset()

			gtx := app.NewContext(&ops, e)
			// xyGridLayout(gtx, th, &aState)
			// myLayout(gtx)

			layout := Rows(
				Rigid(
					AlignMiddle(
						FontSize(22)(
							Label("Opapa"), // one line text
						),
					),
				),
				Flexed(1, myLayout(gtx)),
			)

			if state.showModal {
				layout = Stack(
					layout,
					Centered(
						Background(color.NRGBA{A: 200}, // semi-transparent overlay
							Border(
								Inset(unit.Dp(20),
									Rows(
										Rigid(Label("Modal Window")),
										Rigid(Label("Press ESC to close")),
									),
								),
							),
						),
					),
				)
			}

			// Add key.InputOp to handle key events
			keyFilter.Op(gtx.Ops)
			
			layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}

type (
	C       = layout.Context
	D       = layout.Dimensions
	W       = layout.Widget
	P       = image.Point
	DP      = unit.Dp
	SP      = unit.Sp
	Wrapper = func(W) W
	List    = layout.List
)

var (
	Pt            = f32.Pt
	SpaceUnit  DP = 8
	BorderSize DP = 1

	fonts = gofont.Collection()
	// fontShaper = text.NewShaper(fonts)
	fontShaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	// th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	// th         = material.NewTheme(fonts)
)

func myLayout(gtx C) W {

	return Columns(
		Rigid(ColorBoxW(P{100, 100}, red)),
		Flexed(0.5, ColorBoxW(P{}, blue)),
		Rigid(ColorBoxW(P{100, 100}, red)),
		Rigid(WSpacer1),
		// Flexed(0.5, ColorBoxW(P{}, green)),
		Flexed(0.5, Rows(
			Rigid(
				RoundedCorners(
					// Border(
					BorderActive(
						ColorBoxW(P{100, 100}, SILVER_300),
					),
				),
			),
			// Rigid(HSpacer1),
			Flexed(0.7, ColorBoxW(P{}, blue)),

			Rigid(
				Panel("Panel",
					ColorBoxW(P{100, 100}, red),
				),
			),

			Flexed(0.3,
				// Panel("Panel",
				ColorBoxW(P{}, green),
				// ),
			),

			// Flexed(0.3,
			// 		ColorBoxW(P{400, 200}, green),
			// ),
		)),
	)

	// return layout.Flex{}.Layout(gtx,
	// 	layout.Rigid(func(gtx C) D {
	// 		return ColorBox(gtx, image.Pt(100, 100), red)
	// 	}),
	// 	layout.Flexed(0.5, func(gtx C) D {
	// 		return ColorBox(gtx, gtx.Constraints.Min, blue)
	// 	}),
	// 	layout.Rigid(func(gtx C) D {
	// 		return ColorBox(gtx, image.Pt(100, 100), red)
	// 	}),
	// 	layout.Flexed(0.5, func(gtx C) D {
	// 		return ColorBox(gtx, gtx.Constraints.Min, green)
	// 	}),
	// )

	// // Draw rectangles inside of each other, with 30dp padding.
	// return layout.UniformInset(unit.Dp(30)).Layout(gtx, func(gtx C) D {
	// 	return ColorBox(gtx, gtx.Constraints.Max, getColor(1))
	// })
}

func Rows(children ...layout.FlexChild) W {
	return func(c C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(c, children...)
	}
}

func Columns(children ...layout.FlexChild) W {
	return func(c C) D {
		return layout.Flex{Axis: layout.Horizontal}.Layout(c, children...)
	}
}

var (
	Flexed = layout.Flexed
	Rigid  = layout.Rigid
)

func EmptyWidget(c C) D { return D{} }

func Wrap(w W, wrappers ...Wrapper) W {
	for i := len(wrappers) - 1; i >= 0; i-- {
		w = wrappers[i](w)
	}

	return w
}

func LayoutToWidget(r func(C, W) D, w W) W {
	return func(c C) D {
		return r(c, w)
	}
}

func LayoutToWrapper(r func(C, W) D) func(w W) W {
	return func(w W) W {
		return func(c C) D {
			return r(c, w)
		}
	}
}

// determineLayoutConfig returns configurations for rows based on the window width
func getContextWidth(gtx C) int {
	width := gtx.Constraints.Max.X
	switch {
	case width < 600:
		return 1 // Small screen: 1 column layout
	case width < 1200:
		return 2 // Medium screen: 2 column layout
	default:
		return 3 // Large screen: 3 column layout
	}
}

// Test colors.
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}
)

func getColor(i int) color.NRGBA {
	return color.NRGBA{R: uint8(100 + i*20), G: uint8(150 + i*15), B: uint8(200 - i*10), A: 255}
}

// ColorBox creates a widget with the specified dimensions and color.
func ColorBoxW(size image.Point, color color.NRGBA) W {
	return func(c C) D {
		if size.X == 0 {
			size = c.Constraints.Min
		}
		defer clip.Rect{Max: size}.Push(c.Ops).Pop()
		paint.ColorOp{Color: color}.Add(c.Ops)
		paint.PaintOp{}.Add(c.Ops)
		return D{Size: size}
	}
}
func ColorBox(c C, size image.Point, color color.NRGBA) D {
	defer clip.Rect{Max: size}.Push(c.Ops).Pop()
	paint.ColorOp{Color: color}.Add(c.Ops)
	paint.PaintOp{}.Add(c.Ops)
	return D{Size: size}
}

// xyGridLayout implements a layout based on the XY Grid system similar to Foundation ZURB
func xyGridLayout(gtx C, th *material.Theme, aState *appState) D {
	width := gtx.Constraints.Max.X
	layoutConfig := determineLayoutConfig(width)

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			l := material.H3(th, "XY Grid Layout Example")
			l.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			l.Alignment = text.Middle
			return l.Layout(gtx)
		}),
		layout.Flexed(1, func(gtx C) D {
			rows := make([]layout.FlexChild, len(layoutConfig))
			for i, columns := range layoutConfig {
				columns := columns // capture loop variable
				rows[i] = layout.Rigid(func(gtx C) D {
					return layout.Flex{
						Axis: layout.Horizontal,
					}.Layout(gtx, createRow(gtx, th, aState, columns)...)
				})
			}
			return layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx, rows...)
		}),
	)
}

// determineLayoutConfig returns configurations for rows based on the window width
func determineLayoutConfig(width int) [][]int {
	switch {
	case width < 600:
		return [][]int{
			{12}, {6, 6}, {12}, {6, 6}, {12},
		} // Small screen: 1 column layout
	case width < 1200:
		return [][]int{
			{4, 8}, {6, 6}, {12}, {4, 8}, {8, 4},
		} // Medium screen: 2 column layout
	default:
		return [][]int{
			{3, 9}, {4, 4, 4}, {6, 6}, {3, 6, 3}, {8, 4},
		} // Large screen: 3 column layout
	}
}

// createRow creates a Flex row with the specified column widths
func createRow(gtx C, th *material.Theme, aState *appState, columns []int) []layout.FlexChild {
	children := make([]layout.FlexChild, len(columns))

	totalWidth := gtx.Constraints.Max.X
	for i, col := range columns {
		colWidth := totalWidth * col / 12 // Determine column width as a fraction of the total width
		children[i] = layout.Rigid(func(gtx C) D {
			gtx.Constraints.Min.X = colWidth
			gtx.Constraints.Max.X = colWidth
			btn := material.Button(th, &aState.columnWidgets[i%len(aState.columnWidgets)], "Column "+string(rune('A'+i)))
			btn.Background = color.NRGBA{R: uint8(100 + i*20), G: uint8(150 + i*15), B: uint8(200 - i*10), A: 255}
			return btn.Layout(gtx)
		})
	}
	return children
}

func Stack(layers ...layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		dims := layers[0](gtx)
		for _, layer := range layers[1:] {
			layer(gtx)
		}
		return dims
	}
}

func Centered(w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		dims := w(gtx)
		position := layout.FPt(gtx.Constraints.Min).Sub(layout.FPt(dims.Size).Mul(0.5))
		defer op.Offset(position.Round()).Push(gtx.Ops).Pop()
		return dims
	}
}

func Inset(inset unit.Dp, w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(inset).Layout(gtx, w)
	}
}
