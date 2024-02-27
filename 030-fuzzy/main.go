package main

import (
	"image"
	"image/color"
	"log"
	"strings"

	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type C = layout.Context
type D = layout.Dimensions

var (
	colorWhite = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorGrey  = color.NRGBA{R: 0x55, G: 0x55, B: 0x55, A: 0xff}
	colorBlack = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	colorBlue  = color.NRGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
	colorGreen = color.NRGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
)
var (
	searchInput widget.Editor
	list        = layout.List{Axis: layout.Vertical}
	stringList  = []string{
		"Orders: Buy Limit",
		"Orders: Buy Market",
		"Orders: Sell Limit",
		"Orders: Sell Market",
		"Drawing: Trendline",
		"Drawing: Priceline",
		"Drawing: Volume Profile",
	}
	// stringListFiltered fuzzy.Ranks
	// stringListFiltered = fuzzy.RankFind("", stringList)
	stringListFiltered = stringList
	cursor             int
)

func inputLayout(gtx C, th *material.Theme) D {
	// Wrap the editor in material design
	ed := material.Editor(th, &searchInput, "sec")
	// Define characteristics of the input box
	searchInput.SingleLine = true
	searchInput.Alignment = text.Start
	searchInput.Submit = true
	searchInput.Focus()

	// Define insets ...
	margins := layout.Inset{
		Top:    unit.Dp(5),
		Right:  unit.Dp(5),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(5),
	}

	// layout
	return margins.Layout(gtx,
		func(gtx C) D {
			return ed.Layout(gtx)
		},
	)
}

func listLayout(gtx C, th *material.Theme) D {
	inset := layout.UniformInset(unit.Dp(2))
	return list.Layout(gtx, len(stringListFiltered), func(gtx C, i int) D {
		return inset.Layout(gtx,
			func(gtx C) D {
				prefix := ""
				if cursor == i {
					prefix = "> "
					// FillWithLabel(gtx, th, "", colorWhite, colorBlack)
				}
				return material.Button(th, &widget.Clickable{}, prefix+stringListFiltered[i]).Layout(gtx)
				// return material.Label(th, unit.Sp(14), prefix+stringListFiltered[i]).Layout(gtx)
			},
		)
	})
}

func Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {

	ColorBox(gtx, image.Point{10, 10}, colorGrey)
	// func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	// 	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	// 	paint.ColorOp{Color: color}.Add(gtx.Ops)
	// 	paint.PaintOp{}.Add(gtx.Ops)
	// 	return layout.Dimensions{Size: size}
	// }
	// // handle arrow keys
	tag := &list

	for _, ev := range gtx.Events(tag) {
		switch e := ev.(type) {
		case pointer.Event:
		case key.Event:
			switch e.State {
			case key.Press:
				log.Println(e.Name, "GOT key.Press:", e.Name, e.Modifiers)
				if e.Modifiers.Contain(key.ModCtrl) {
					// ctrl is pressed
				}
				if e.Name == "↓" || e.Modifiers.Contain(key.ModCtrl) && e.Name == "J" {
					cursor = cursor + 1
				}
				if e.Name == "↑" || e.Modifiers.Contain(key.ModCtrl) && e.Name == "K" {
					cursor = cursor - 1
				}
				if cursor < 0 {
					cursor = 0
				}
				if cursor > len(stringListFiltered)-1 {
					cursor = len(stringListFiltered) - 1
				}
				log.Println("cursor", cursor)
			case key.Release:
				//
			}
		}
	}
	key.InputOp{
		Keys: key.NameEscape + key.Set("|Ctrl-J|←|↓|↑|→|"),
		Tag:  tag, // Use the window as the event routing tag. This means we can call gtx.Events(w) and get these events.
	}.Add(gtx.Ops)
	// for _, e := range gtx.Events(tag) {
	// 	if e, ok := e.(pointer.Event); ok {
	// 		switch e.Type {
	// 		case pointer.Press:
	// 			// b.pressed = true
	// 		case pointer.Release:
	// 			// b.pressed = false
	// 		}
	// 	}
	// }

	// Handle keyboard input for the search field
	for _, ke := range searchInput.Events() {
		// log.Printf("%v %+v\n", i, reflect.TypeOf(ke))
		if _, ok := ke.(widget.ChangeEvent); ok {
			// on change - filter
			inputString := searchInput.Text()
			inputString = strings.TrimSpace(inputString)
			stringListFiltered = fuzzy.FindNormalizedFold(inputString, stringList)
			log.Println(stringListFiltered)
			// reset cursor
			cursor = 0
		}
		if ke, ok := ke.(widget.SubmitEvent); ok {
			// Process the submitted search query
			query := strings.TrimSpace(ke.Text)
			log.Println("got submit query:", query)
		}
	}

	// layout everything
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return inputLayout(gtx, th)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			// FillWithLabel(gtx, th, "Black On Grey", colorWhite, colorGrey)
			return listLayout(gtx, th)
		}),
		// layout.Flexed(0.1, func(gtx layout.Context) layout.Dimensions {
		// 	return FillWithLabel(gtx, th, "Black On Green", colorBlack, colorGreen)
		// }),
	)
}

func main() {
	Loop()
}
