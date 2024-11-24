package ui

import (
	"image/color"
)

// GithubColors
var (
	// https://primer.style/foundations/primitives/color
	// ThemeDark
	// fg
	// fgColor-default #e6edf3
	// fgColor-muted #848d97
	// fgColor-disabled #8b949e
	// fgColor-link #4493f8
	// bg
	// bgColor-default #0d1117
	// bgColor-accent #1f6feb
	ColorFg         = Color(0xe6edf3)
	ColorFgMuted    = Color(0x848d97)
	ColorFgDisabled = Color(0x8b949e)
	ColorFgLink     = Color(0x4493f8)
	ColorBg         = Color(0x0d1117)
	ColorBgAccent   = Color(0x1f6feb)
)

// COLOR CONVERSIONS

// Color converts a hex color value (like 0xe6edf3) to an NRGBA color.
// The alpha channel is set to fully opaque (255).
// Use this for defining solid colors from hex values.
func Color(hex int) color.NRGBA {
	return color.NRGBA{
		R: uint8(hex >> 16),
		G: uint8(hex >> 8),
		B: uint8(hex),
		A: 255,
	}
}

// Alpha creates a new color with the specified alpha value while preserving
// the original RGB values. Use this to make a color partially transparent.
// alpha ranges from 0 (fully transparent) to 255 (fully opaque).
func Alpha(c color.NRGBA, alpha uint8) color.NRGBA {
	c.A = alpha
	return c
}

// MixColor blends two colors together based on a percentage.
// percent specifies how much of c1 to use (0-100).
// Example: MixColor(red, blue, 60) gives a color that is 60% red and 40% blue.
func MixColor(c1, c2 color.NRGBA, percent int) color.NRGBA {
	p1 := float32(percent) / float32(100.0)
	p2 := 1 - p1
	return color.NRGBA{
		R: uint8(float32(c1.R)*p1 + float32(c2.R)*p2),
		G: uint8(float32(c1.G)*p1 + float32(c2.G)*p2),
		B: uint8(float32(c1.B)*p1 + float32(c2.B)*p2),
		A: uint8(float32(c1.A)*p1 + float32(c2.A)*p2),
	}
}

// RGB converts a 24-bit RGB color value to NRGBA.
// The alpha channel is set to fully opaque (255).
// Example: RGB(0xFF0000) creates a pure red color.
func RGB(c uint32) color.NRGBA {
	return ARGB(0xff000000 | c)
}

// ARGB converts a 32-bit ARGB color value to NRGBA.
// The alpha channel is taken from the most significant byte.
// Example: ARGB(0x80FF0000) creates a 50% transparent red color.
func ARGB(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
