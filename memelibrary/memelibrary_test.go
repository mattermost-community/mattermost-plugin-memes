package memelibrary

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsImageAsset(t *testing.T) {
	assert.True(t, isImageAsset("all-the-things.jpg"))
	assert.False(t, isImageAsset("all-the-things.json"))
}

func TestMustLoadImage(t *testing.T) {
	assert.NotPanics(t, func() {
		img := mustLoadImage("images/all-the-things.jpg")
		assert.NotNil(t, img)
	})

	assert.Panics(t, func() {
		mustLoadImage("this-asset-does-not-exist.jpg")
	})

	assert.Panics(t, func() {
		mustLoadImage("metadata/all-the-things.yaml")
	})
}

func TestTemplate(t *testing.T) {
	assert.Nil(t, Template("not-a-template"))

	template := Template("all-the-things")
	require.NotNil(t, template)
	assert.NotNil(t, template.Image)
}

func TestPatternMatch(t *testing.T) {
	for name, metadata := range Memes() {
		for _, pattern := range metadata.Patterns {
			template, text := PatternMatch(pattern.Example)
			assert.Equal(t, Template(name), template)
			assert.NotNil(t, text)
		}
	}
}
