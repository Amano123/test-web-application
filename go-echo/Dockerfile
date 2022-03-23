FROM golang:latest

# ワーキングディレクトリの設定
WORKDIR /app
COPY ./app /app

RUN go mod init main \
  && go mod tidy \
  && go build

RUN go install github.com/cosmtrek/air@latest

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64