package main

import "github.com/olablt/gio-lab/ui"

func WidgetIf(cond bool, w ui.W) ui.W {
	if cond {
		return w
	} else {
		return ui.EmptyWidget
	}
}
