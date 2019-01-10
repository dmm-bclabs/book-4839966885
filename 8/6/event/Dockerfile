# Build
FROM golang

RUN mkdir -p /go/src \
    && mkdir -p /go/bin \
    && mkdir -p /go/pkg

ARG APP_PATH="src/app"

RUN mkdir -p $GOPATH/$APP_PATH
ADD . $GOPATH/$APP_PATH
WORKDIR $GOPATH/src

RUN apt-get update \
    && apt-get install -y mysql-client \
    && go get -u github.com/go-sql-driver/mysql \
    && go get -u github.com/jinzhu/gorm \
    && go get -u github.com/ethereum/go-ethereum

WORKDIR $GOPATH/$APP_PATH
EXPOSE 8081
CMD ["/bin/bash"]