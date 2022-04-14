FROM golang:1.18-alpine

WORKDIR /go/src
COPY src /go/src

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
