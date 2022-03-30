package meme

import (
	"image"

	"github.com/fogleman/gg"
)

type Template struct {
	Name      string
	Image     image.Image
	TextSlots []*TextSlot
}

func (t *Template) Render(text []string) (image.Image, error) {
	// b := t.Image.Bounds()
	// img := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	dc := gg.NewContextForImage(t.Image)

	// draw.Draw(img, img.Bounds(), t.Image, b.Min, draw.Src)

	for i, slot := range t.TextSlots {
		if i >= len(text) {
			break
		}
		slot.Render(dc, text[i], DEBUG)
	}

	return dc.Image(), nil
}
