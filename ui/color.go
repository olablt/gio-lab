package ui

import (
	"image/color"
)

// GithubColors
var (
	// https://primer.style/foundations/primitives/color
	// webpack://./dist/internalCss/dark.css

	// Foreground colors
	// --tooltip-bgColor: #3d444d;
	// --tooltip-fgColor: #ffffff;
	FgColor           = Color(0xf0f6fc)
	FgColorMuted      = Color(0x9198a1)
	FgColorOnEmphasis = Color(0xffffff)
	FgColorOnInverse  = Color(0x010409)
	FgColorWhite      = Color(0xffffff)
	FgColorBlack      = Color(0x010409)
	FgColorDisabled   = Color(0x656c7699)
	FgColorLink       = Color(0x4493f8)
	FgColorNeutral    = Color(0x9198a1)
	FgColorAccent     = Color(0x4493f8)
	FgColorSuccess    = Color(0x3fb950)
	FgColorOpen       = Color(0x3fb950)
	FgColorAttention  = Color(0xd29922)
	FgColorSevere     = Color(0xdb6d28)
	FgColorDanger     = Color(0xf85149)
	FgColorClosed     = Color(0xf85149)
	FgColorDone       = Color(0xab7df8)

	// Background colors
	BgColor               = Color(0x0d1117)
	BgColorMuted          = Color(0x151b23)
	BgColorInset          = Color(0x010409)
	BgColorEmphasis       = Color(0x3d444d)
	BgColorInverse        = Color(0xffffff)
	BgColorWhite          = Color(0xffffff)
	BgColorBlack          = Color(0x010409)
	BgColorDisabled       = Color(0x212830)
	BgColorTransparent    = Color(0x00000000)
	BgColorNeutralMuted   = Color(0x656c7633)
	BgColorNeutral        = Color(0x656c76)
	BgColorAccentMuted    = Color(0x388bfd1a)
	BgColorAccent         = Color(0x1f6feb)
	BgColorSuccessMuted   = Color(0x2ea04326)
	BgColorSuccess        = Color(0x238636)
	BgColorOpenMuted      = Color(0x2ea04326)
	BgColorOpen           = Color(0x238636)
	BgColorAttentionMuted = Color(0xbb800926)
	BgColorAttention      = Color(0x9e6a03)
	BgColorSevereMuted    = Color(0xdb6d281a)
	BgColorSevere         = Color(0xbd561d)
	BgColorDangerMuted    = Color(0xf851491a)
	BgColorDanger         = Color(0xda3633)
	BgColorClosedMuted    = Color(0xf851491a)
	BgColorClosed         = Color(0xda3633)
	BgColorDoneMuted      = Color(0xab7df826)
	BgColorDone           = Color(0x8957e5)

	// Border colors
	BorderColor            = Color(0x3d444d)
	BorderColorMuted       = Color(0x3d444db3)
	BorderColorEmphasis    = Color(0x656c76)
	BorderColorDisabled    = Color(0x656c761a)
	BorderColorTransparent = Color(0x00000000)
	BorderColorNeutral     = Color(0x3d444db3)
	BorderColorAccent      = Color(0x388bfd66)
	BorderColorSuccess     = Color(0x2ea04366)
	BorderColorOpen        = Color(0x2ea04366)
	BorderColorAttention   = Color(0xbb800966)
	BorderColorSevere      = Color(0xdb6d2866)
	BorderColorDanger      = Color(0xf8514966)
	BorderColorClosed      = Color(0xf8514966)
	BorderColorDone        = Color(0xab7df866)

	// Header colors
	HeaderFgColor            = Color(0xffffffb3)
	HeaderFgColorLogo        = Color(0xf0f6fc)
	HeaderBgColor            = Color(0x151b23f2)
	HeaderBorderColorDivider = Color(0x656c76)
	HeaderSearchBgColor      = Color(0x0d1117)
	HeaderSearchBorderColor  = Color(0x2a313c)

	// Data colors
	DataBlueColor        = Color(0x0576ff)
	DataBlueColorMuted   = Color(0x001a47)
	DataAuburnColor      = Color(0xa86f6b)
	DataAuburnColorMuted = Color(0x271817)
	DataOrangeColor      = Color(0x984b10)
	DataOrangeColorMuted = Color(0x311708)
	DataYellowColor      = Color(0x895906)
	DataYellowColorMuted = Color(0x2e1a00)
	DataGreenColor       = Color(0x2f6f37)
	DataGreenColorMuted  = Color(0x122117)
	DataTealColor        = Color(0x106c70)
	DataTealColorMuted   = Color(0x041f25)
	DataPurpleColor      = Color(0x975bf1)
	DataPurpleColorMuted = Color(0x211047)
	DataPinkColor        = Color(0xd34591)
	DataPinkColorMuted   = Color(0x2d1524)
	DataRedColor         = Color(0xeb3342)
	DataRedColorMuted    = Color(0x3c0614)
	DataGrayColor        = Color(0x576270)
	DataGrayColorMuted   = Color(0x1c1c1c)
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
