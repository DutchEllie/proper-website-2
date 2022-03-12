FROM golang:1.17.8-alpine AS builder
WORKDIR /project
ADD . /project/
RUN go mod tidy
RUN GOARCH=wasm GOOS=js go build -o web/app.wasm
RUN go build -o app

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /project/web ./web/
COPY --from=builder /project/app ./
EXPOSE 8000
CMD ["./app"]  

