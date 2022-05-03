FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN go mod init goctr

COPY . .

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm

RUN go build

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
