FROM golang:1.17-alpine as builder

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

COPY . /goapi

EXPOSE 8000

#CMD ["go", "run", "main.go"]
RUN go build -o goapi main.go


FROM alpine:3.10
COPY --from=builder /goapi/goapi /goapi
COPY --from=builder /goapi/db/dbconfig.ini /db/dbconfig.ini
CMD ["./goapi"]
