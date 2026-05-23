.PHONY: all build clean

all: build

build:
	@go build -gcflags="all=-B" -ldflags="all=-s -w" -o ~/.local/bin/unibox

clean:
	@rm -f ~/.local/bin/unibox

push:
	@git add .
	@git commit -m "update" | true
	@git push
