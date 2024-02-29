package ui

import (
	"image/color"
)

// Test Colors.
var (
	ColorRed   = color.NRGBA{R: 0xC0, G: 0x40, B: 0x40, A: 0xFF}
	ColorGreen = color.NRGBA{R: 0x40, G: 0xC0, B: 0x40, A: 0xFF}
	ColorBlue  = color.NRGBA{R: 0x40, G: 0x40, B: 0xC0, A: 0xFF}

	// WB
	ColorWhite = color.NRGBA{255, 255, 255, 255}
	ColorBlack = color.NRGBA{0, 0, 0, 255}
	// BLUE https://www.flatuiColorpicker.com/blue-rgb-color-model/
	ColorNeonBlue      = color.NRGBA{45, 85, 255, 255}
	ColorAzureRadiance = color.NRGBA{3, 138, 255, 255}
	ColorTealBlue      = color.NRGBA{4, 59, 92, 255}
	ColorBlackPearl    = color.NRGBA{8, 14, 44, 255}
	ColorShark         = color.NRGBA{36, 37, 42, 255}
	// Color = color.NRGBA{ , 255}

	// GREY https://www.flatuiColorpicker.com/grey-rgb-color-model/
	ColorLynch      = color.NRGBA{108, 122, 137, 255}
	ColorSilverSand = color.NRGBA{189, 195, 199, 255}
	ColorAthensGray = color.NRGBA{239, 239, 240, 255}

	// YELLOW https://www.flatuiColorpicker.com/color/yellow/
	ColorWitchHase  = color.NRGBA{255, 246, 143, 255}
	ColorYellow     = color.NRGBA{240, 255, 0, 255}
	ColorGinFizz    = color.NRGBA{255, 249, 222, 255}
	ColorLaserLemon = color.NRGBA{230, 255, 110, 255}
)

// COLOR CONVERSIONS

func Color(hex int) color.NRGBA {
	return color.NRGBA{
		R: uint8(hex >> 16),
		G: uint8(hex >> 8),
		B: uint8(hex),
		A: 255,
	}
}

func Alpha(c color.NRGBA, alpha uint8) color.NRGBA {
	c.A = alpha
	return c
}

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

func RGB(c uint32) color.NRGBA {
	return ARGB(0xff000000 | c)
}

func ARGB(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
