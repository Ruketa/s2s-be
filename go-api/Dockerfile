FROM golang:1.17-alpine

WORKDIR /goapi

RUN apk add --no-cache build-base \
 && apk --no-cache add postgresql-dev \
 && go mod init goapi \
 && go get github.com/gin-gonic/gin \
 && go get github.com/gin-contrib/cors \
 && go get github.com/go-ini/ini \
 && go get github.com/lib/pq \
 && go get github.com/google/uuid \
 && go get gorm.io/gorm \
 && go get gorm.io/driver/postgres \
 && go get gopkg.in/ini.v1 

EXPOSE 8000