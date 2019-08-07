module github.com/mattermost/mattermost-plugin-memes

go 1.12

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gorilla/mux v1.7.3
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51
	github.com/mattermost/mattermost-server v5.12.0+incompatible
	github.com/pkg/errors v0.8.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80 // indirect
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
	gopkg.in/yaml.v2 v2.2.2
)

// Workaround for https://github.com/golang/go/issues/30831 and fallout.
replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
