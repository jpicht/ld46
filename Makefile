all: dist/main.wasm dist/wasm_exec.js dist/index.html dist/gfx.png dist/.htaccess

dist/main.wasm: Makefile
	GOOS=js GOARCH=wasm go build -o $@

ifndef GOROOT
GOROOT=$(shell go env | grep GOROOT | cut -d'"' -f 2)
endif

dist/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js dist
	cp $< $@

dist/index.html: static/index.html
	cp $< $@

dist/gfx.png: gfx/gfx.xcf Makefile gfx/gfx.sh
	gfx/gfx.sh $< $@

dist/.htaccess: static/.htaccess
	cp $< $@

dist:
	mkdir -p dist

clean:
	rm -rf dist

serve: all
	SERVE_DIR=$${PWD}/dist go run ./serve

.PHONY: \
	all \
	clean \
	dist/main.wasm \
	serve
