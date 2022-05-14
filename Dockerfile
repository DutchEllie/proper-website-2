FROM golang:1.17.8-alpine AS builder
ARG APIURL
WORKDIR /project
ADD . /project/
RUN go mod tidy
RUN GOARCH=wasm GOOS=js go build -o web/app.wasm -ldflags="-X 'main.ApiURL=$APIURL'" ./src
RUN	go build -o app -ldflags="-X 'main.ApiURL=$APIURL'" ./src

FROM alpine:latest AS staging
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /project/web ./web/
COPY --from=builder /project/app ./
EXPOSE 8000
ENV TESTING true
CMD ["./app"]  