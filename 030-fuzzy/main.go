package main

import (
	"log"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Command struct {
	Name string
	Func func()
	Key  key.Filter
}

func main() {
	cp := NewCommandPalette()
	commands := []Command{
		{Name: "File: New", Func: nil, Key: key.Filter{Name: "N", Required: key.ModCtrl}},
		{Name: "File: Open", Func: nil, Key: key.Filter{}},
		{Name: "File: Save", Func: nil, Key: key.Filter{}},
		{Name: "File: Save As", Func: nil, Key: key.Filter{}},
		{Name: "Edit: Undo", Func: nil, Key: key.Filter{}},
		{Name: "Edit: Redo", Func: nil, Key: key.Filter{}},
		{Name: "Edit: Cut", Func: nil, Key: key.Filter{}},
		{Name: "Format: Indent", Func: nil, Key: key.Filter{}},
		{Name: "Format: Outdent", Func: nil, Key: key.Filter{}},
	}
	for _, c := range commands {
		cp.RegisterCommand(c.Name, c.Func, c.Key)
	}
	// set "File: New" callback
	cp.SetCallback("File: New", func() {
		log.Println("[FIRE] File: New")
	})

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
