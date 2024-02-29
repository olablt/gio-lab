package main

import (
	"gioui-lab/ui"
	"log"
	"strings"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type CommandPalette struct {
	SearchInput        *widget.Editor
	List               *widget.List
	StringList         []string
	StringListFiltered []string
	Cursor             int
	cache              struct {
		inputString string
		inputLen    int
	}
}

func NewCommandPalette() *CommandPalette {
	cp := &CommandPalette{
		SearchInput:        &widget.Editor{SingleLine: true, Submit: true, Alignment: text.Start},
		List:               &widget.List{List: layout.List{Axis: layout.Vertical}},
		StringList:         []string{},
		StringListFiltered: []string{},
		Cursor:             0,
	}
	cp.StringList = []string{
		"File: New",
		"File: Open",
		"File: Save",
		"File: Save As",
		"Edit: Undo",
		"Edit: Redo",
		"Edit: Cut",
		"Format: Indent",
		"Format: Outdent",
	}
	cp.StringListFiltered = cp.StringList
	return cp
}

func (cp *CommandPalette) InputLayout(gtx C, th *material.Theme) D {
	// Wrap the editor in material design
	ed := material.Editor(th, cp.SearchInput, "search phrase")

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

func (cp *CommandPalette) ListLayout(gtx C, th *material.Theme) D {
	// Define insets for the list items
	in := layout.Inset{
		Top:    unit.Dp(0),
		Right:  unit.Dp(0),
		Bottom: unit.Dp(5),
		Left:   unit.Dp(5),
	}
	// layout the list
	return material.List(th, cp.List).Layout(gtx, len(cp.StringListFiltered), func(gtx C, i int) D {
		return in.Layout(gtx,
			func(gtx C) D {
				th2 := *th
				if i != cp.Cursor {
					th2.Palette.ContrastBg = ui.Alpha(th2.Palette.ContrastBg, 100)
					th2.Palette.ContrastFg = th.Palette.Fg
				}
				return material.Button(&th2, &widget.Clickable{}, cp.StringListFiltered[i]).Layout(gtx)
			},
		)
	})
}

func (cp *CommandPalette) submit(i int) {
	log.Printf("[DEBUG] SUBMIT '%v'", cp.StringListFiltered[cp.Cursor])
}

func (cp *CommandPalette) Update(gtx layout.Context) {
	// ColorBox(gtx, image.Point{10, 10}, colorGrey)
	// handle arrow keys
	// tag := &cp.List
	tag := &cp.SearchInput
	event.Op(gtx.Ops, tag)

	filters := []event.Filter{
		key.Filter{Name: "↑"},
		key.Filter{Name: "↓"},
		key.Filter{Name: "J", Required: key.ModCtrl},
		key.Filter{Name: "K", Required: key.ModCtrl},
		key.Filter{Name: key.NameReturn},
		// key.FocusFilter{Target: tag},
		// key.Filter{Focus: tag, Name: "↑"},
		// key.Filter{Focus: tag, Name: "↓"},
		// key.Filter{Focus: tag, Name: "J", Required: key.ModCtrl},
		// key.Filter{Focus: tag, Name: "K", Required: key.ModCtrl},
	}
	//	Keys: key.NameEscape + key.Set("|Ctrl-J|←|↓|↑|→|"),
	// New key event reading
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
		// log.Printf("[DEBUG] got key.%v", ev.Name)
		if ev.State == key.Press {
			// log.Printf("[DEBUG] got key.%v", ev.Name)
			// handle enter
			if ev.Name == key.NameReturn {
				cp.submit(cp.Cursor)
			}
			// cursor movement
			if ev.Name == "↓" || ev.Name == "J" {
				cp.Cursor = cp.Cursor + 1
			}
			if ev.Name == "↑" || ev.Name == "K" {
				cp.Cursor = cp.Cursor - 1
			}
			if cp.Cursor < 0 {
				cp.Cursor = 0
			}
			if cp.Cursor > len(cp.StringListFiltered)-1 {
				cp.Cursor = len(cp.StringListFiltered) - 1
			}
		}
	}

	cp.SearchInput.Update(gtx)
	if cp.SearchInput.Text() != cp.cache.inputString {
		cp.cache.inputString = cp.SearchInput.Text()
		cp.cache.inputString = strings.TrimSpace(cp.cache.inputString)
		cp.StringListFiltered = fuzzy.FindNormalizedFold(cp.cache.inputString, cp.StringList)
		cp.Cursor = 0
		log.Println("input changed!", cp.cache.inputString)
	}
	// if cp.SearchInput.Len() != cp.cache.inputLen {
	// 	cp.cache.inputLen = cp.SearchInput.Len()
	// 	log.Println("input changed! len", cp.cache.inputLen)
	// }
}

func (cp *CommandPalette) Layout(gtx layout.Context, th *material.Theme) D {
	gtx.Execute(key.FocusCmd{Tag: cp.SearchInput})
	// process events
	cp.Update(gtx)

	// layout everything
	d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return cp.InputLayout(gtx, th)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			// FillWithLabel(gtx, th, "Black On Grey", colorWhite, colorGrey)
			return cp.ListLayout(gtx, th)
		}),
		// layout.Flexed(0.1, func(gtx layout.Context) layout.Dimensions {
		// 	return FillWithLabel(gtx, th, "Black On Green", colorBlack, colorGreen)
		// }),
	)

	return d
}

func main() {
	cp := NewCommandPalette()
	Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
		gtx.Metric = unit.Metric{
			PxPerDp: 1.8,
			PxPerSp: 1.8,
			// PxPerDp: 4,
			// PxPerSp: 4,
		}
		// layout
		cp.Layout(gtx, th)
	})
}
