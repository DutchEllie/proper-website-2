APIURL_prod := https://api.nicecock.eu/api/comment
APIURL_staging := https://api.nicecock.eu/api/testingcomment

build:
	GOARCH=wasm GOOS=js go build -ldflags="-X 'dutchellie.nl/DutchEllie/proper-website-2/components.ApiURL=${APIURL_staging}'" -o web/app.wasm
	go build -ldflags="-X 'dutchellie.nl/DutchEllie/proper-website-2/components.ApiURL=${APIURL_staging}'" -o app

build-prod:
	GOARCH=wasm GOOS=js go build -ldflags="-X 'dutchellie.nl/DutchEllie/proper-website-2/components.ApiURL=${APIURL_prod}'" -o web/app.wasm
	go build -ldflags="-X 'dutchellie.nl/DutchEllie/proper-website-2/components.ApiURL=${APIURL_prod}'" -o app

run: build
	./app

run-prod: build-prod
	./app