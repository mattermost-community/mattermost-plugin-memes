package memelibrary

import (
	"image"
	"image/color"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/mattermost/mattermost-plugin-memes/server/meme"
)

type Pattern struct {
	Pattern string
	Text    []string
	Example string

	pattern *regexp.Regexp
}

type Slot struct {
	X         int
	Y         int
	Width     int
	Height    int
	Font      string
	TextColor []int `yaml:"text_color"`
}

type Metadata struct {
	Aliases  []string
	Patterns []*Pattern
	Slots    []*Slot
	Example  []string
}

func ParseMetadata(in []byte) (*Metadata, error) {
	var m Metadata
	if err := yaml.Unmarshal(in, &m); err != nil {
		return nil, err
	}
	for _, pattern := range m.Patterns {
		r, err := regexp.Compile("(?i)^" + pattern.Pattern + "$")
		if err != nil {
			return nil, err
		}
		pattern.pattern = r
	}
	return &m, nil
}

func (m *Metadata) PatternMatch(input string) []string {
	for _, pattern := range m.Patterns {
		if text := pattern.Match(input); text != nil {
			return text
		}
	}
	return nil
}

func (m *Metadata) TextSlots(bounds image.Rectangle) (slots []*meme.TextSlot) {
	if m.Slots != nil {
		for _, slot := range m.Slots {
			textSlot := &meme.TextSlot{
				Bounds: image.Rect(slot.X, slot.Y, slot.X+slot.Width, slot.Y+slot.Height),
			}
			if slot.Font != "" {
				textSlot.Font = fonts[slot.Font]
			} else {
				textSlot.Font = fonts["Anton-Regular"]
				textSlot.TextColor = color.White
				textSlot.OutlineColor = color.Black
				textSlot.AllUppercase = true
			}
			if c := sliceToColor(slot.TextColor); c != nil {
				textSlot.TextColor = c
			}
			slots = append(slots, textSlot)
		}
		return
	}

	padding := bounds.Dy() / 20
	return []*meme.TextSlot{
		{
			Bounds:       image.Rect(padding, padding, bounds.Dx()-padding, bounds.Dy()/4),
			Font:         fonts["Anton-Regular"],
			TextColor:    color.White,
			OutlineColor: color.Black,
			AllUppercase: true,
		},
		{
			Bounds:       image.Rect(padding, bounds.Dy()*3/4, bounds.Dx()-padding, bounds.Dy()-padding),
			Font:         fonts["Anton-Regular"],
			TextColor:    color.White,
			OutlineColor: color.Black,
			AllUppercase: true,
		},
	}
}

func sliceToColor(s []int) color.Color {
	switch len(s) {
	case 1:
		return color.Gray16{uint16(s[0]) << 1}
	case 2:
		return color.RGBA{uint8(s[0]), uint8(s[0]), uint8(s[0]), uint8(s[1])}
	case 3:
		return color.RGBA{uint8(s[0]), uint8(s[1]), uint8(s[2]), 255}
	case 4:
		return color.RGBA{uint8(s[0]), uint8(s[1]), uint8(s[2]), uint8(s[3])}
	}
	return nil
}

func (p *Pattern) Match(input string) (text []string) {
	if matches := p.pattern.FindStringSubmatchIndex(input); matches != nil {
		for _, slotText := range p.Text {
			text = append(text, strings.TrimSpace(string(p.pattern.ExpandString(nil, slotText, input, matches))))
		}
	}
	return
}
