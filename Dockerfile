# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.work ./
COPY web-service-gin/go.mod ./
#COPY web-service-gin/go.sum ./
#COPY datetime-utils/go.mod ./
#COPY image-utils/go.mod ./
#COPY image-utils/go.sum ./
#COPY users/go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /web-service-gin

EXPOSE 8080

CMD [ "/web-service-gin" ]