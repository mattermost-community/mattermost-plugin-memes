# Disclaimer

**This repository is community supported and not maintained by Mattermost. Mattermost disclaims liability for integrations, including Third Party Integrations and Mattermost Integrations. Integrations may be modified or discontinued at any time.**

# Mattermost Memes Plugin

[![Build Status](https://img.shields.io/circleci/project/github/mattermost/mattermost-plugin-memes/master.svg)](https://circleci.com/gh/mattermost/mattermost-plugin-memes)
[![Code Coverage](https://img.shields.io/codecov/c/github/mattermost/mattermost-plugin-memes/master.svg)](https://codecov.io/gh/mattermost/mattermost-plugin-memes)

**Maintainer:** [@hanzei](https://github.com/hanzei)

This plugin will create a slash command that you can use to create memes!

<img src="screenshot.png" width="583" height="426" />

`/meme everywhere "memes." "memes everywhere"`

For more information like available memes or command syntax type `/meme ` and press enter.

## Installation

1. Go to https://github.com/mattermost/mattermost-plugin-memes/releases/latest to download the latest release file in zip or tar.gz format.
2. Upload the file through **System Console > Plugins > Management**. See [documentation](https://docs.mattermost.com/administration/plugins.html#set-up-guide) for more details.

## Development

Read our documentation about the [Developer Workflow](https://developers.mattermost.com/extend/plugins/developer-workflow/) and [Developer Setup](https://developers.mattermost.com/extend/plugins/developer-setup/) for more information about developing and extending plugins.

Run `make memelibrary` to bundle up the meme assets (images, fonts, etc.).

For convenience, you can run the plugin from your terminal to preview an image for a given input. For example, on macOS, you can run the following to generate the above meme and open it in Preview:

`go run server/plugin.go -out demo.jpg 'memes. memes everywhere' && open demo.jpg`

This is especially useful when adding or modifying memes as you can quickly modify assets, `make memelibrary`, and preview the result using the above command. (See the files in ` memelibrary/assets` to get started with that.)
