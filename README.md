# mattermost-plugin-memes [![CircleCI](https://circleci.com/gh/mattermost/mattermost-plugin-memes.svg?style=svg)](https://circleci.com/gh/mattermost/mattermost-plugin-memes)

This plugin will create a slash command that you can use to create memes!

<img src="screenshot.png" width="583" height="426" />

`/meme memes. memes everywhere`

For more information like avaliable memes or command syntax type `/meme ` and press enter.

## Installation

Go to the GitHub releases tab and download the latest release for your server architecture. You can upload this file in the Mattermost system console to install the plugin.

## Developing

Run `make vendor` to install third-party dependencies and `make memelibrary/assets.go` to bundle up the meme assets (images, fonts, etc.). From there, you can develop like any other Go project: Hack away and use `go test`, `go build`, etc.

For convenience, you can run the plugin from your terminal to preview an image for a given input. For example, on macOS, you can run the following to generate the above meme and open it in Preview:

`go run plugin.go -out demo.jpg 'memes. memes everywhere' && open demo.jpg`

This is especially useful when adding or modifying memes as you can quickly modify assets, `make memelibrary/assets.go`, and preview the result using the above command. (See the files in memelibrary/assets to get started with that.)

If you want to create a fully bundled plugin that will run on a local server, you can use `make mattermost-memes-plugin.tar.gz`.

## Releasing

To make a release, update the version number in plugin.yaml, and create a release via the GitHub interface. Travis will upload the distributables for you.
