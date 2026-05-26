.PHONY: all build clean push

all: build

build:
	@go build -gcflags="" -ldflags="all=-s -w" -trimpath -o ~/.local/bin/unibox

clean:
	@rm -f ~/.local/bin/unibox

push:
	@git add .
	@git commit -m "update" | true
	@git push
