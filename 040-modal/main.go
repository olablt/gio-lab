package main

import (
	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/olablt/gioui-lab/ui"
)

type Command struct {
	Name string
	Func func()
	Key  key.Filter
	Icon *widget.Icon
}

func main() {
	myApp := NewMyApp()

	ui.Loop(func(win *app.Window, gtx layout.Context, th *material.Theme) {
		gtx.Metric = unit.Metric{
			PxPerDp: 1.5,
			PxPerSp: 1.5,
			// PxPerDp: 1.8,
			// PxPerSp: 1.8,
			// PxPerDp: 4,
			// PxPerSp: 4,
		}
		// layout
		// cp.Layout(gtx, th)
		myApp.Layout(gtx, th)
	})
}
