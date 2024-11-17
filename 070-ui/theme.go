package main

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
	TextColor     color.NRGBA
	MaxLines      int
}

var Theme = ThemeStyle{
	FontSize:      13,
	FontFamily:    fontShaper,
	FontWeight:    font.Normal,
	TextAlignment: text.Start,
	TextColor:     BLACK_500,
	MaxLines:      0,
}
