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
	"gioui.org/widget/material"
	// "gioui.org/widget/material"
)

type Area struct {
	Name         string
	PointerPress bool
	KeyPress     bool
	// subscribe to key events
	Keys      []key.Name
	areaStack clip.Stack
	Focus     bool
}

func (a *Area) ProcessEvents(gtx layout.Context) {
	// Declare the tag.
	tag := a

	// Confine the area of interest to a gtx Max
	a.areaStack = clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops)

	// new input op
	event.Op(gtx.Ops, tag)

	// New pointer event reading
	a.PointerPress = false
	for {
		event, ok := gtx.Event(
			pointer.Filter{
				Target: tag,
				Kinds:  pointer.Press | pointer.Enter | pointer.Leave,
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
				if a.Focus {
					gtx.Execute(key.FocusCmd{Tag: tag})
				}
				// log
				log.Printf("[%v] got Pointer.Enter", a.Name)
			case pointer.Leave:
				// log
				log.Printf("[%v] got Pointer.Leave", a.Name)
			}
		}
	}

	event.Op(gtx.Ops, tag)
	// New key event reading
	filters := []event.Filter{}
	if a.Focus {
		filters = append(filters, key.FocusFilter{Target: tag})
	}
	// set key filters
	for _, k := range a.Keys {
		if a.Focus {
			filters = append(filters, key.Filter{Focus: tag, Name: k})
		} else {
			filters = append(filters, key.Filter{Focus: nil, Name: k})
		}
	}

	// New key event reading
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
			case key.NameEscape:
				// log
				log.Printf("[%v] got key.Escape", a.Name)
				a.KeyPress = true
			case key.NameReturn:
				// log
				log.Printf("[%v] got key.Return", a.Name)
				a.KeyPress = true
			default:
				// log
				log.Printf("[%v] got key.%v", a.Name, ev.Name)
				a.KeyPress = true
			}
		}
	}

}

func (a *Area) Pop() {
	a.areaStack.Pop()
}

func main() {
	Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
		ChartLayout(gtx, th)
	})
}

// define event areas
var windowArea = &Area{Name: "Window", Keys: []key.Name{key.NameEscape, "Q"}, Focus: false}
var chartArea = &Area{Name: "Chart", Keys: []key.Name{key.NameEscape}, Focus: true}
var xArea = &Area{Name: "X", Keys: []key.Name{"1", "2"}, Focus: true}
var yArea = &Area{Name: "Y", Keys: []key.Name{"1", "2"}, Focus: true}

func ChartLayout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// whole window events
	windowArea.ProcessEvents(gtx)
	defer windowArea.Pop()

	return layout.Flex{}.Layout(gtx,
		// MAIN CHART & X Axis AREA
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			// return ColorBox(gtx, gtx.Constraints.Min, green)
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				// MAIN CHART AREA
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					chartArea.ProcessEvents(gtx)
					chartArea.Pop()
					return FillWithLabelH3(gtx, th, "Chart", green)
				}),
				// X Axis AREA
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					xArea.ProcessEvents(gtx)
					xArea.Pop()
					return FillWithLabelH3(gtx, th, "X", red)
				}),
			)
		}),
		// Y Axis
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			yArea.ProcessEvents(gtx)
			yArea.Pop()
			return FillWithLabelH3(gtx, th, " Y ", blue)
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
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: alpha}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: alpha}
)

// FillWithLabelH3 creates a label with the specified text and background color
func FillWithLabelH3(gtx layout.Context, th *material.Theme, text string, backgroundColor color.NRGBA) layout.Dimensions {
	ColorBox(gtx, gtx.Constraints.Max, backgroundColor)
	return layout.Center.Layout(gtx, material.H3(th, text).Layout)
}

// ColorBox creates a box with the specified color
func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

// Loop is a helper function that runs the app event loop
func Loop(fn func(win *app.Window, gtx layout.Context, th *material.Theme)) {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	go func() {
		w := app.NewWindow(
			app.Title("oGio"),
			app.Size(unit.Dp(1920/2), unit.Dp(1080/2)),
		)
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
