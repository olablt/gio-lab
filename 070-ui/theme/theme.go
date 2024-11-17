package theme

import (
	"image/color"
	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
)

// Theme defines the visual properties of the UI
type Theme struct {
	Colors     ColorScheme
	Typography TypographyScheme
	Spacing    SpacingScheme
}

type ColorScheme struct {
	Primary   color.NRGBA
	Secondary color.NRGBA
	Text      color.NRGBA
	Border    color.NRGBA
	Background color.NRGBA
}

type TypographyScheme struct {
	Default     TextStyle
	Title       TextStyle
	Subtitle    TextStyle
	Body        TextStyle
}

type TextStyle struct {
	Size        unit.SP
	Font        font.Font
	Shaper      *text.Shaper
	Color       color.NRGBA
	Alignment   text.Alignment
	MaxLines    int
}

type SpacingScheme struct {
	Small  unit.Dp
	Medium unit.Dp
	Large  unit.Dp
}

// DefaultTheme returns a new theme with default values
func DefaultTheme(shaper *text.Shaper) *Theme {
	return &Theme{
		Colors: ColorScheme{
			Primary:    color.NRGBA{R: 98, G: 0, B: 238, A: 255},
			Secondary:  color.NRGBA{R: 3, G: 218, B: 198, A: 255},
			Text:      color.NRGBA{R: 0, G: 0, B: 0, A: 255},
			Border:    color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			Background: color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		},
		Typography: TypographyScheme{
			Default: TextStyle{
				Size:      unit.Sp(16),
				Font:      font.Font{Weight: font.Normal},
				Shaper:    shaper,
				Color:     color.NRGBA{R: 0, G: 0, B: 0, A: 255},
				Alignment: text.Start,
				MaxLines:  0,
			},
			Title: TextStyle{
				Size:      unit.Sp(24),
				Font:      font.Font{Weight: font.Bold},
				Shaper:    shaper,
				Color:     color.NRGBA{R: 0, G: 0, B: 0, A: 255},
				Alignment: text.Start,
				MaxLines:  1,
			},
		},
		Spacing: SpacingScheme{
			Small:  unit.Dp(8),
			Medium: unit.Dp(16),
			Large:  unit.Dp(24),
		},
	}
}
