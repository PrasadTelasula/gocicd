FROM golang:alpine AS builder

MAINTAINER prasad.telasula@yahoo.com

WORKDIR /go/src/app

COPY main.go .

RUN go build -o webserver .

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/app/ /app

EXPOSE 8000

CMD ["./webserver"]
