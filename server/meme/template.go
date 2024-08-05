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
	dc := gg.NewContextForImage(t.Image)

	for i, slot := range t.TextSlots {
		if i >= len(text) {
			break
		}
		slot.Render(dc, text[i], DEBUG)
	}

	return dc.Image(), nil
}
