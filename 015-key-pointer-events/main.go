package main

import (
	"image/color"
	"log"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/ui"
	// "gioui.org/widget/material"
)

type Area struct {
	Name         string
	PointerPress bool
	KeyPress     bool
	// subscribe to key events
	Keys                   []key.Name
	areaStack              clip.Stack
	CaptureKeysWhenInFocus bool // should this area be focused for key events when mouse is not over it?
	StatusFocused          bool
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
				if a.CaptureKeysWhenInFocus {
					gtx.Execute(key.FocusCmd{Tag: tag})
				}
				a.StatusFocused = true
				// log
				log.Printf("[%v] got Pointer.Enter", a.Name)
			case pointer.Leave:
				// log
				log.Printf("[%v] got Pointer.Leave", a.Name)
				a.StatusFocused = false
			}
		}
	}

	event.Op(gtx.Ops, tag)
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
	ui.Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
		ChartLayout(gtx, th)
	})
}

// define event areas
var windowArea = &Area{Name: "Window", Keys: []key.Name{key.NameEscape, "Q"}, CaptureKeysWhenInFocus: false}
var chartArea = &Area{Name: "Chart", Keys: []key.Name{key.NameEscape}, CaptureKeysWhenInFocus: true}
var xArea = &Area{Name: "X", Keys: []key.Name{"1", "2"}, CaptureKeysWhenInFocus: true}
var yArea = &Area{Name: "Y", Keys: []key.Name{"1", "2"}, CaptureKeysWhenInFocus: true}

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
					return ui.FillWithLabelH3(gtx, th, "Chart", darkenColor(green, chartArea.StatusFocused))
				}),
				// X Axis AREA
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					xArea.ProcessEvents(gtx)
					xArea.Pop()
					return ui.FillWithLabelH3(gtx, th, "X", darkenColor(red, xArea.StatusFocused))
				}),
			)
		}),
		// Y Axis
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			yArea.ProcessEvents(gtx)
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
	green      = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: alpha}
	blue       = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: alpha}
)

// darkenColor reduces the brightness of a color if focused
func darkenColor(c color.NRGBA, focused bool) color.NRGBA {
	if !focused {
		return c
	}
	// Reduce each color component by ~25% for focused state
	return color.NRGBA{
		R: uint8(float64(c.R) * 0.75),
		G: uint8(float64(c.G) * 0.75),
		B: uint8(float64(c.B) * 0.75),
		A: c.A,
	}
}
