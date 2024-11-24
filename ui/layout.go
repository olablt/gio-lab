package ui

import "gioui.org/layout"

func Rows(children ...layout.FlexChild) W {
	return func(c C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(c, children...)
	}
}

func Columns(children ...layout.FlexChild) W {
	return func(c C) D {
		return layout.Flex{Axis: layout.Horizontal}.Layout(c, children...)
	}
}

func ColumnsVCentered(children ...layout.FlexChild) W {
	return func(c C) D {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(c, children...)
	}
}

func EmptyWidget(c C) D { return D{} }

func Wrap(w W, wrappers ...Wrapper) W {
	for i := len(wrappers) - 1; i >= 0; i-- {
		w = wrappers[i](w)
	}

	return w
}

func LayoutToWidget(r func(C, W) D, w W) W {
	return func(c C) D {
		return r(c, w)
	}
}

func LayoutToWrapper(r func(C, W) D) func(w W) W {
	return func(w W) W {
		return func(c C) D {
			return r(c, w)
		}
	}
}
