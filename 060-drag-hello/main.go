// I need to track mouse position inside the image and outside the image.

package main

import (
	"image"
	"image/color"
	"log"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/olablt/gio-lab/qasset"
	"github.com/olablt/gio-lab/ui"
)

var imageOp = paint.NewImageOp(qasset.Neutral)
var imageLocation = f32.Pt(300, 300)
var pointerLocation = f32.Pt(300, 300)

// func main() { qapp.Layout(MyLayout) }
func main() {
	ui.Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
		// ChartLayout(gtx, th)
		MyLayout(gtx)
	})
}

// MyLayout handles rendering and input
func MyLayout(gtx layout.Context) layout.Dimensions {
	winTag := 0x01
	imgTag := 0x02

	// 1.a PROCESS drag EVENTS
	var dragOffset f32.Point
	for {
		event, ok := gtx.Event(
			pointer.Filter{
				Target: imgTag,
				Kinds:  pointer.Press | pointer.Enter | pointer.Leave | pointer.Move | pointer.Drag | pointer.Release,
			},
		)
		if !ok {
			break
		}
		ev, ok := event.(pointer.Event)
		if ok {
			// handle ev
			log.Printf("got pointer event %v", ev.Kind.String())
			switch ev.Kind {
			case pointer.Drag, pointer.Release, pointer.Press:
				dragOffset = ev.Position.Sub(f32.Pt(50, 50))
			}
		}
	}
	// for {
	// 	ev, ok := drag.Update(gtx.Metric, gtx.Source, gesture.Both)
	// 	if !ok {
	// 		break
	// 	}
	// 	log.Printf("got drag event %v %+v", ev.Kind.String(), ev)
	// 	switch ev.Kind {
	// 	case pointer.Drag, pointer.Release, pointer.Press:
	// 		dragOffset = ev.Position.Sub(f32.Pt(50, 50))
	// 		// pointerLocation = ev.Position
	// 	}
	// }

	// 1.b PROCESS pointer EVENTS
	for {
		event, ok := gtx.Event(
			pointer.Filter{
				Target: winTag,
				Kinds:  pointer.Press | pointer.Enter | pointer.Leave | pointer.Move | pointer.Drag | pointer.Release,
			},
		)
		if !ok {
			break
		}
		ev, ok := event.(pointer.Event)
		if ok {
			// handle ev
			log.Printf("got pointer event %v", ev.Kind.String())
			switch ev.Kind {
			case pointer.Move, pointer.Drag:
				pointerLocation = ev.Position
			}
		}
	}

	// 2. LAYOUT
	// register window area
	winArea := clip.Rect(image.Rectangle{Max: gtx.Constraints.Max}).Push(gtx.Ops)
	// retister winTag for input events inside the window
	event.Op(gtx.Ops, winTag)

	// update the offset, must be after drag.Events
	imageLocation = imageLocation.Add(dragOffset)
	affineArea := op.Affine(f32.Affine2D{}.Offset(imageLocation)).Push(gtx.Ops) //.Pop()

	// register image area for input events
	imageArea := clip.Rect{Max: imageOp.Size()}.Push(gtx.Ops) //.Pop()
	// retister imgTag for input events inside the window
	event.Op(gtx.Ops, imgTag)
	pointer.CursorGrab.Add(gtx.Ops) // set mouse cursor
	imageArea.Pop()

	// draw the image
	imageOp.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	affineArea.Pop()

	winArea.Pop()

	// 3. DRAW CROSSING LINES
	pt := pointerLocation.Round()
	w := gtx.Constraints.Max.X
	h := gtx.Constraints.Max.Y
	drawRect(pt.X, 0, 1, h, gtx)
	drawRect(0, pt.Y, w, 1, gtx)

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func drawRect(x, y int, w, h int, gtx layout.Context) {
	defer op.Offset(image.Point{X: x, Y: y}).Push(gtx.Ops).Pop()
	defer clip.Rect{Max: image.Pt(w, h)}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 200}}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
}
