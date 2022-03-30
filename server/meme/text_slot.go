package meme

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type HorizontalAlignment int

const (
	Left HorizontalAlignment = iota
	Center
	Right
)

type VerticalAlignment int

const (
	Top VerticalAlignment = iota
	Middle
	Bottom
)

const DEBUG = false
const lineSpacing = 0.8

type TextSlot struct {
	Font                *opentype.Font
	TextColor           color.Color
	OutlineColor        color.Color
	HorizontalAlignment HorizontalAlignment
	VerticalAlignment   VerticalAlignment
	OutlineWidth        int
	AllUppercase        bool
	Bounds              image.Rectangle
	MaxFontSize         float64
	Rotation            float64
}

func (s *TextSlot) Render(dc *gg.Context, text string, debug bool) {
	if s.AllUppercase {
		text = strings.ToUpper(text)
	}
	dc.Push()
	// Compute font size by taking measurement of a text
	face, _, textHeight := faceForSlot(text, s.Font, s.MaxFontSize, s.Bounds.Dx(), s.Bounds.Dy())
	dc.SetFontFace(face)

	rotCenterX := float64((s.Bounds.Min.X + s.Bounds.Dx()) / 2)
	rotCenterY := float64((s.Bounds.Min.Y + s.Bounds.Dy()) / 2)

	dc.RotateAbout(gg.Radians(s.Rotation),
		rotCenterX, rotCenterY)
	dc.SetColor(s.OutlineColor)

	outlineWidth := s.OutlineWidth // "stroke" size
	if outlineWidth == 0 {
		outlineWidth = 8
	}

	textColor := s.TextColor
	if textColor == nil {
		textColor = color.Black
	}

	xStart := float64(s.Bounds.Min.X)
	xAlign := gg.Align(s.HorizontalAlignment)

	// Bottom padding for the text, because the string measurements account for top-padding but not bottom
	yAlign := 0.1

	yStart := float64(s.Bounds.Min.Y)
	switch s.VerticalAlignment {
	case Top:
		yStart = float64(s.Bounds.Min.Y)
	case Middle:
		yStart = (float64(s.Bounds.Dy())-textHeight)/2.0 +
			float64(s.Bounds.Min.Y)
	case Bottom:
		yStart = float64(s.Bounds.Max.Y) - textHeight
	default:
		break
	}

	// Some black magic to draw outline
	if s.OutlineColor != nil {
		offset := face.Metrics().Height / 256 * fixed.Int26_6(outlineWidth)
		for _, delta := range []fixed.Point26_6{
			{X: offset, Y: offset},
			{X: -offset, Y: offset},
			{X: -offset, Y: -offset},
			{X: offset, Y: -offset},
		} {
			x := xStart + float64(delta.X)/64
			y := yStart + float64(delta.Y)/64
			dc.DrawStringWrapped(text, x, y, 0, yAlign, float64(s.Bounds.Dx()), lineSpacing, xAlign)
		}
	}

	dc.SetColor(textColor)
	// dc.SetRGB(0.5, 0.5, 0.2)
	dc.DrawStringWrapped(text,
		xStart,
		yStart,
		0, yAlign, float64(s.Bounds.Dx()), lineSpacing, xAlign)

	if debug {
		fmt.Printf("X %v Y %v W %v H %v\n", float64(s.Bounds.Min.X), float64(s.Bounds.Min.Y),
			float64(s.Bounds.Dx()), float64(s.Bounds.Dy()))
		dc.SetRGB(1, 0, 0)
		dc.DrawRectangle(float64(s.Bounds.Min.X), float64(s.Bounds.Min.Y),
			float64(s.Bounds.Dx()), float64(s.Bounds.Dy()))
		dc.Stroke()
	}

	dc.Pop()
}

func faceForSlot(text string, fontt *opentype.Font, maxFontSize float64, width int, height int) (font.Face, float64, float64) {
	fontSize := maxFontSize
	if fontSize == 0.0 {
		fontSize = 80
	}
	w := 0.0
	h := 0.0
	dc := gg.NewContext(10, 10)
	face, _ := opentype.NewFace(fontt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	for fontSize >= 6.0 {
		face, _ = opentype.NewFace(fontt, &opentype.FaceOptions{
			Size:    fontSize,
			DPI:     72,
			Hinting: font.HintingNone,
		})
		dc.SetFontFace(face)
		lines := dc.WordWrap(text, float64(width))
		w, h = dc.MeasureMultilineString(strings.Join(lines, "\n"), lineSpacing)
		if w > float64(width) || h > float64(height)*lineSpacing {
			fontSize -= (fontSize + 9) / 10
			continue
		}
		break
	}

	return face, w, h
}
