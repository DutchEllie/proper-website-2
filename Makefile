APIURL_prod := https://api.quenten.nl/api
APIURL_staging := https://api.quenten.nl/api/testing

build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm -ldflags="-X 'main.ApiURL=${APIURL_staging}'"  ./src
	go build -o app -ldflags="-X 'main.ApiURL=${APIURL_staging}'" ./src

build-new:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./src
	go build -o app ./src

build-prod:
	GOARCH=wasm GOOS=js go build -o web/app.wasm -ldflags="-X 'main.ApiURL=${APIURL_prod}'" ./src
	go build -o app -ldflags="-X 'main.ApiURL=${APIURL_prod}'" ./src

run: build
	./app

run-new: build-new
	APIURL=${APIURL_staging} ./app

run-prod: build-prod
	./app