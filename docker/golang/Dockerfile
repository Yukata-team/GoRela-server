# golang のイメージを取得
FROM golang:1.14

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/github.com/Yukata-team/GoRela-server
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work

EXPOSE 80

RUN go get github.com/dgrijalva/jwt-go
RUN	go get github.com/jinzhu/gorm
RUN	go get github.com/k0kubun/pp
RUN	go get github.com/labstack/gommon
RUN	go get github.com/labstack/echo
