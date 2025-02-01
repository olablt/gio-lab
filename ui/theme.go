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

// Text:           color.NRGBA{189, 195, 199, 255}, // silver sand

var Theme = ThemeStyle{
	FontSize:      12,
	FontFamily:    FontShaper,
	FontWeight:    font.Normal,
	TextAlignment: text.Start,
	// TextColor:     FgColor,
	TextColor: color.NRGBA{189, 195, 199, 255}, // silver sand
	MaxLines:  0,
}
