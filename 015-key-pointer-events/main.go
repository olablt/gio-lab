package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui"
	// "gioui.org/widget/material"
)

type Area struct {
	Name         string
	PointerPress bool
	KeyPress     bool
	Key          key.Name
	// subscribe to key events
	Keys                   []key.Name
	areaStack              clip.Stack
	CaptureKeysWhenInFocus bool // should this area be focused for key events when mouse is not over it?
	StatusFocused          bool
}

func (a *Area) Update(gtx layout.Context) {
	// Declare the tag.
	tag := a

	// New pointer event reading
	a.PointerPress = false
	for {
		event, ok := gtx.Event(
			pointer.Filter{
				Target: tag,
				Kinds:  pointer.Press | pointer.Enter | pointer.Leave | pointer.Cancel,
			},
		)
		if !ok {
			break
		}
		ev, ok := event.(pointer.Event)
		if ok {
			// handle ev
			// log.Printf("got pointer event %#+v", ev)
			switch ev.Kind {
			case pointer.Press:
				a.PointerPress = true
				// log
				log.Printf("[%v] got Pointer.Press", a.Name)
			case pointer.Enter:
				if a.CaptureKeysWhenInFocus {
					gtx.Execute(key.FocusCmd{Tag: tag})
				}
				a.StatusFocused = true
				// log
				log.Printf("[%v] got Pointer.Enter", a.Name)
			// case pointer.Move:
			// 	log.Printf("[%v] got Pointer.Move %+v", a.Name, ev.Position)
			case pointer.Leave:
				// log
				log.Printf("[%v] got Pointer.Leave", a.Name)
				a.StatusFocused = false
			case pointer.Cancel:
				// Handle the same way as Leave
				log.Printf("[%v] got Pointer.Cancel (window leave)", a.Name)
				a.StatusFocused = false
			}
		}
	}

	// New key event reading
	filters := []event.Filter{}
	if a.CaptureKeysWhenInFocus {
		filters = append(filters, key.FocusFilter{Target: tag})
	}
	// set key filters
	for _, k := range a.Keys {
		if a.CaptureKeysWhenInFocus {
			filters = append(filters, key.Filter{Focus: tag, Name: k})
		} else {
			filters = append(filters, key.Filter{Focus: nil, Name: k})
		}
	}

	// New key event reading
	a.Key = ""
	a.KeyPress = false
	for {
		event, ok := gtx.Event(filters...)
		if !ok {
			break
		}
		ev, ok := event.(key.Event)
		if !ok {
			continue
		}
		// handle ev
		if ev.State == key.Press {
			// log.Printf("[%v] got key event %#+v", a.Name, ev)
			switch ev.Name {
			// case key.NameEscape:
			// 	// log
			// 	log.Printf("[%v] got key.Escape", a.Name)
			// 	a.KeyPress = true
			// 	a.Key = ev.Name
			// case key.NameReturn:
			// 	// log
			// 	log.Printf("[%v] got key.Return", a.Name)
			// 	a.KeyPress = true
			// 	a.Key = ev.Name
			default:
				// log
				log.Printf("[%v] got key.%v", a.Name, ev.Name)
				a.KeyPress = true
				a.Key = ev.Name
			}
		}
	}

}

func (a *Area) Layout(gtx layout.Context) {
	// Confine the area of interest to a gtx Max
	a.areaStack = clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops)
	event.Op(gtx.Ops, a)
}

// stop collecting events for this area
func (a *Area) Pop() {
	a.areaStack.Pop()
}

func main() {
	clickable := widget.Clickable{}
	editor := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}

	modalVisible := false

	Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
		ed := material.Editor(th, editor, "hint")
		ed.TextSize = unit.Sp(30)
		ed.HintColor = black

		// process events from previous frame
		modalArea.Update(gtx)
		chartArea.Update(gtx)
		barsArea.Update(gtx)
		xArea.Update(gtx)
		yArea.Update(gtx)
		if chartArea.Key == "P" {
			if !modalVisible {
				modalVisible = true
				// gtx.Execute(key.FocusCmd{Tag: modalArea})
				editor.SetText("")
				gtx.Execute(key.FocusCmd{Tag: editor})
			}
		}
		if (chartArea.Key == key.NameEscape || modalArea.Key == key.NameEscape || clickable.Clicked(gtx)) && modalVisible {
			log.Println("GOT ESC MODAL")
			modalVisible = false
		}

		layout.Flex{Axis: 1}.Layout(gtx,
			// HEADER
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				th.Fg = color.NRGBA{0, 0, 0, 255}
				return material.Label(th, unit.Sp(30), "Header").Layout(gtx)
			}),
			// CHART
			layout.Flexed(5, func(gtx layout.Context) layout.Dimensions {
				inset := layout.UniformInset(unit.Dp(0))
				return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return ChartLayout(gtx, th)
				})
			}),
		)

		// MODAL
		if modalVisible {

			layout.Background{}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					// white transparent background
					return clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						// log.Printf("%+v", gtx.Constraints)
						dims := ui.ColorBox(gtx, gtx.Constraints.Min, ui.Alpha(white, 150))
						return dims
					})
				},
				func(gtx layout.Context) layout.Dimensions {
					w := gtx.Dp(200)
					h := gtx.Dp(300)
					gtx.Constraints.Max = image.Pt(w, h)
					gtx.Constraints.Min = image.Pt(w, h)
					modalArea.Layout(gtx)
					defer modalArea.Pop()
					dims := ui.FillWithLabelH3(gtx, th, " Modal ", darkenColor(blue, modalArea.StatusFocused))
					ed.Layout(gtx)
					return dims
				})

		}

	})

	// Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
	// 	ChartLayout(gtx, th)
	// })
}

// define event areas
var chartArea = &Area{Name: "Chart", Keys: []key.Name{key.NameEscape, "Q", "P"}, CaptureKeysWhenInFocus: false}
var barsArea = &Area{Name: "Bars", Keys: []key.Name{key.NameEscape}, CaptureKeysWhenInFocus: true}
var xArea = &Area{Name: "X", Keys: []key.Name{"1", "2"}, CaptureKeysWhenInFocus: true}
var yArea = &Area{Name: "Y", Keys: []key.Name{"1", "2"}, CaptureKeysWhenInFocus: true}
var modalArea = &Area{Name: "Modal", Keys: []key.Name{key.NameEscape, "Q", "1", "2"}, CaptureKeysWhenInFocus: true}

func ChartLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// whole window events
	chartArea.Layout(gtx)
	defer chartArea.Pop()

	return layout.Flex{}.Layout(gtx,
		// MAIN CHART & X Axis AREA
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			// return ColorBox(gtx, gtx.Constraints.Min, green)
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// MAIN CHART AREA
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					barsArea.Layout(gtx)
					barsArea.Pop()
					return ui.FillWithLabelH3(gtx, th, "Bars", darkenColor(green, barsArea.StatusFocused))
				}),
				// X Axis AREA
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					xArea.Layout(gtx)
					xArea.Pop()
					return ui.FillWithLabelH3(gtx, th, "X", darkenColor(red, xArea.StatusFocused))
				}),
			)
		}),
		// Y Axis
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			yArea.Layout(gtx)
			yArea.Pop()
			return ui.FillWithLabelH3(gtx, th, " Y ", darkenColor(blue, yArea.StatusFocused))
		}),
	)
}

// -------------------
// Helper functions

// Test colors.
var (
	background = color.NRGBA{R: 0xC0, G: 0xC0, B: 0xC0, A: 0xFF}
	alpha      = uint8(255)
	red        = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: alpha}
	white      = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: alpha}
	black      = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: alpha}
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: alpha}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: alpha}
)

// darkenColor reduces the brightness of a color if focused
func darkenColor(c color.NRGBA, focused bool) color.NRGBA {
	if !focused {
		return c
	}
	// Reduce each color component by ~25% for focused state
	return ui.Alpha(c, 200)
	// return color.NRGBA{
	// 	R: uint8(float64(c.R) * 0.75),
	// 	G: uint8(float64(c.G) * 0.75),
	// 	B: uint8(float64(c.B) * 0.75),
	// 	A: c.A,
	// }
}

func Loop(fn func(win *app.Window, gtx layout.Context, th *material.Theme)) {
	th := material.NewTheme()
	// th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(LoadFontCollection()))
	th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Collection()))
	// set Github theme
	th.Palette.Fg = color.NRGBA{R: 0xe6, G: 0xed, B: 0xf3, A: 0xff}

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
