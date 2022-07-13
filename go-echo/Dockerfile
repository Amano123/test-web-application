FROM golang:1.18-alpine

WORKDIR /go/src
# COPY src/go.mod .
# COPY src/go.sum .
# COPY src/.air.toml .

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go install github.com/cosmtrek/air@latest

COPY ./setup.sh /

CMD ["sh", "/setup.sh"]
