FROM golang:alpine
# アップデートとgitのインストール！！
RUN apk add --update \
&&  apk add git
# appディレクトリの作成
RUN mkdir /go/src/go-echo
# ワーキングディレクトリの設定
WORKDIR /go/src/go-echo
ADD . /go/src/go-echo

RUN go get github.com/labstack/echo \
# &&  go get github.com/pilu/fresh
&&  go get -u github.com/cosmtrek/air 