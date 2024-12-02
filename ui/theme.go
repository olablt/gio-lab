package ui

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/text"
)

type ThemeStyle struct {
	FontSize      SP
	FontFamily    *text.Shaper
	FontWeight    font.Weight
	TextAlignment text.Alignment
	MaxLines      int
	TextColor     color.NRGBA
}

var Theme = ThemeStyle{
	FontSize:      13,
	FontFamily:    FontShaper,
	FontWeight:    font.Normal,
	TextAlignment: text.Start,
	TextColor:     FgColor,
	MaxLines:      0,
}
