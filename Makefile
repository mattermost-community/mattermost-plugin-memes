.PHONY: test

all: test dist

TAR_PLUGIN_EXE_TRANSFORM = --transform 'flags=r;s|dist/intermediate/plugin_.*|plugin.exe|'
ifneq (,$(findstring bsdtar,$(shell tar --version)))
	TAR_PLUGIN_EXE_TRANSFORM = -s '|dist/intermediate/plugin_.*|plugin.exe|'
endif

dist: vendor memelibrary/assets.go $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' ./...) plugin.yaml
	rm -rf ./dist
	go get github.com/mitchellh/gox
	$(shell go env GOPATH)/bin/gox -osarch='darwin/amd64 linux/amd64 windows/amd64' -output 'dist/intermediate/plugin_{{.OS}}_{{.Arch}}'
	tar -czvf dist/mattermost-meme-plugin-darwin-amd64.tar.gz $(TAR_PLUGIN_EXE_TRANSFORM) dist/intermediate/plugin_darwin_amd64 plugin.yaml
	tar -czvf dist/mattermost-meme-plugin-linux-amd64.tar.gz $(TAR_PLUGIN_EXE_TRANSFORM) dist/intermediate/plugin_linux_amd64 plugin.yaml
	tar -czvf dist/mattermost-meme-plugin-windows-amd64.tar.gz $(TAR_PLUGIN_EXE_TRANSFORM) dist/intermediate/plugin_windows_amd64.exe plugin.yaml
	rm -rf dist/intermediate

memelibrary/assets.go: $(shell find memelibrary/assets)
	go get github.com/jteeuwen/go-bindata/...
	$(shell go env GOPATH)/bin/go-bindata -o $@ -pkg memelibrary -prefix memelibrary/assets/ -ignore '(^|/)\..*' memelibrary/assets/...

mattermost-memes-plugin.tar.gz: vendor memelibrary/assets.go $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' ./...) plugin.yaml 
	go build -o plugin.exe
	tar -czvf $@ plugin.exe plugin.yaml
	rm plugin.exe

test: vendor
	go test -v ./...

vendor: glide.lock
	go get github.com/Masterminds/glide
	$(shell go env GOPATH)/bin/glide install
