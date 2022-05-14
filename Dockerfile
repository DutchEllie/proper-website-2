#FROM golang:1.17.8-alpine AS builder
#ARG APIURL
#WORKDIR /project
#ADD . /project/
#RUN go mod tidy
#RUN GOARCH=wasm GOOS=js go build -o web/app.wasm -ldflags="-X 'main.ApiURL=$APIURL'" ./src
#RUN	go build -o app -ldflags="-X 'main.ApiURL=$APIURL'" ./src

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
RUN mkdir ./web
COPY ./web ./web
COPY ./build/web/app.wasm ./web/
COPY ./build/app ./
EXPOSE 8000
CMD ["./app"]  