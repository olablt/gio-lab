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
	ui.Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
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
					return ui.FillWithLabelH3(gtx, th, "Chart", green)
				}),
				// X Axis AREA
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					xArea.ProcessEvents(gtx)
					xArea.Pop()
					return ui.FillWithLabelH3(gtx, th, "X", red)
				}),
			)
		}),
		// Y Axis
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			yArea.ProcessEvents(gtx)
			yArea.Pop()
			return ui.FillWithLabelH3(gtx, th, " Y ", blue)
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
