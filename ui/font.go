package ui

import (
	"fmt"
	"log"
	"os"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

func LoadFontCollection() []font.FontFace {

	// // load source sans fonts
	// faces := []font.FontFace{}
	// faces = append(faces, LoadFontFace(font.Font{}, "assets/source_sans_pro_regular.otf"))
	// faces = append(faces, LoadFontFace(font.Font{Weight: font.Medium}, "assets/source_sans_pro_semibold.otf"))
	// faces = append(faces, LoadFontFace(font.Font{Weight: font.Bold}, "assets/source_sans_pro_bold.otf"))

	// // load go fonts
	// faces := []font.FontFace{}
	// faces = append(faces, opentypeParse(font.Font{}, goregular.TTF))
	// faces = append(faces, opentypeParse(font.Font{Weight: font.Medium}, gomedium.TTF))
	// faces = append(faces, opentypeParse(font.Font{Weight: font.Bold}, gobold.TTF))

	// load Roboto fonts
	faces := []font.FontFace{}
	// faces = append(faces, LoadFontFace(font.Font{Weight: font.Thin}, "assets/Roboto-Thin.ttf"))
	faces = append(faces, LoadFontFace(font.Font{}, "assets/Roboto-Light.ttf"))
	faces = append(faces, LoadFontFace(font.Font{Weight: font.Medium}, "assets/Roboto-Regular.ttf"))
	faces = append(faces, LoadFontFace(font.Font{Weight: font.Bold}, "assets/Roboto-Medium.ttf"))
	////
	//faces = append(faces, LoadFontFace(font.Font{}, "assets/Roboto-Regular.ttf"))
	//faces = append(faces, LoadFontFace(font.Font{Weight: font.Medium}, "assets/Roboto-Medium.ttf"))
	//faces = append(faces, LoadFontFace(font.Font{Weight: font.Bold}, "assets/Roboto-Bold.ttf"))

	return faces
}

func LoadFontFace(fnt font.Font, filename string) font.FontFace {
	fontData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error loading font file:", err)
	}
	face, err := opentype.Parse(fontData)
	// faces, err := opentype.ParseCollection(fontData)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	fnt.Typeface = "Go"
	fontFace := font.FontFace{Font: fnt, Face: face}
	return fontFace
	// return faces
}

func opentypeParse(fnt font.Font, fontByte []byte) font.FontFace {
	face, err := opentype.Parse(fontByte)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	fnt.Typeface = "Go"
	fontFace := font.FontFace{Font: fnt, Face: face}
	return fontFace
}

// func register(fnt font.Font, fontByte []byte) {
// 	face, err := opentype.Parse(fontByte)
// 	if err != nil {
// 		panic(fmt.Errorf("failed to parse font: %v", err))
// 	}
// 	fnt.Typeface = "Go"
// 	collection = append(collection, font.FontFace{Font: fnt, Face: face})
// }

// func getFontByte(path string) ([]byte, error) {
// 	return content.ReadFile(path)
// }
