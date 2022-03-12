build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	cp web/app.wasm staticsite/web/app.wasm
	scp staticsite/web/app.wasm ellieserver:/home/ellie/nicecock/test/web/
	go build -o app

build-all: build
	cp -r web/* staticsite/web/
	scp -r staticsite/* ellieserver:/home/ellie/nicecock/test

run: build
	./app