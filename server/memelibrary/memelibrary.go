package memelibrary

import (
	"bytes"
	"image"
	"path/filepath"
	"strings"

	// registering decoder functions
	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/image/font/opentype"

	"github.com/mattermost/mattermost-plugin-memes/server/meme"
)

var fonts = make(map[string]*opentype.Font)
var images = make(map[string]image.Image)
var metadata = make(map[string]*Metadata)
var templates = make(map[string]*meme.Template)

func isImageAsset(assetName string) bool {
	ext := strings.ToLower(filepath.Ext(assetName))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

func mustLoadImage(assetName string) image.Image {
	img, _, err := image.Decode(bytes.NewReader(MustAsset(assetName)))
	if err != nil {
		panic(err)
	}
	return img
}

func init() {
	fontAssets, _ := AssetDir("fonts")
	for _, assetName := range fontAssets {
		fontName := strings.TrimSuffix(assetName, filepath.Ext(assetName))
		font, err := opentype.Parse(MustAsset(filepath.Join("fonts", assetName)))
		if err != nil {
			panic(err)
		}
		fonts[fontName] = font
	}

	imageAssets, _ := AssetDir("images")
	for _, assetName := range imageAssets {
		if !isImageAsset(assetName) {
			continue
		}
		templateName := strings.TrimSuffix(assetName, filepath.Ext(assetName))
		images[templateName] = mustLoadImage(filepath.Join("images", assetName))
	}

	metadataAssets, _ := AssetDir("metadata")
	for _, assetName := range metadataAssets {
		ext := filepath.Ext(assetName)
		if ext != ".yaml" {
			continue
		}

		templateName := strings.TrimSuffix(assetName, ext)

		m, err := ParseMetadata(MustAsset(filepath.Join("metadata", assetName)))
		if err != nil {
			panic(err)
		}
		metadata[templateName] = m
	}

	for templateName, metadata := range metadata {
		img := images[templateName]

		template := &meme.Template{
			Name:      templateName,
			Image:     img,
			TextSlots: metadata.TextSlots(img.Bounds()),
		}
		templates[templateName] = template
		for _, alias := range metadata.Aliases {
			templates[alias] = template
		}
	}
}

func Memes() map[string]*Metadata {
	return metadata
}

func Template(name string) *meme.Template {
	return templates[name]
}

func PatternMatch(input string) (*meme.Template, []string) {
	for templateName, metadata := range metadata {
		if text := metadata.PatternMatch(input); text != nil {
			return templates[templateName], text
		}
	}
	return nil, nil
}
