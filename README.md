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

From Mattermost 5.16 and later, the Memes Plugin is included in the Plugin Marketplace which can be accessed from **Main Menu > Plugins Marketplace**. You can install the Memes plugin there.

In Mattermost 5.15 and earlier, follow these steps:

1. Go to https://github.com/mattermost/mattermost-plugin-memes/releases/latest to download the latest release file in zip or tar.gz format.
2. Upload the file through **System Console > Plugins > Management**. See [documentation](https://docs.mattermost.com/administration/plugins.html#set-up-guide) for more details.

## Development

This plugin contains a server portion only. Read our documentation about the [Developer Workflow](https://developers.mattermost.com/integrate/plugins/developer-workflow/) and [Developer Setup](https://developers.mattermost.com/integrate/plugins/developer-setup/) for more information about developing and extending plugins.

Run `make memelibrary` to bundle up the meme assets (images, fonts, etc.).

For convenience, you can run the plugin from your terminal to preview an image for a given input. For example, on macOS, you can run the following to generate the above meme and open it in Preview:

`go run server/plugin.go -out demo.jpg 'memes. memes everywhere' && open demo.jpg`

This is especially useful when adding or modifying memes as you can quickly modify assets, `make memelibrary`, and preview the result using the above command. (See the files in ` memelibrary/assets` to get started with that.)

### Releasing new versions

The version of a plugin is determined at compile time, automatically populating a `version` field in the [plugin manifest](plugin.json):
* If the current commit matches a tag, the version will match after stripping any leading `v`, e.g. `1.3.1`.
* Otherwise, the version will combine the nearest tag with `git rev-parse --short HEAD`, e.g. `1.3.1+d06e53e1`.
* If there is no version tag, an empty version will be combined with the short hash, e.g. `0.0.0+76081421`.

To disable this behaviour, manually populate and maintain the `version` field.
