APIURL_prod := https://api.nicecock.eu/api
APIURL_staging := https://api.nicecock.eu/api/testing

build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm -ldflags="-X 'main.ApiURL=${APIURL_staging}'"  ./src
	go build -o app -ldflags="-X 'main.ApiURL=${APIURL_staging}'" ./src

build-prod:
	GOARCH=wasm GOOS=js go build -o web/app.wasm -ldflags="-X 'main.ApiURL=${APIURL_prod}'" ./src
	go build -o app -ldflags="-X 'main.ApiURL=${APIURL_prod}'" ./src

run: build
	./app

run-prod: build-prod
	./app