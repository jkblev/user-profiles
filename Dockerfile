# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.work ./
# The below is VERY CLUNKY, and I assume there is a better way to do this
# using go workspaces, or by doing my go package dependencies in a different way.
# But "go mod download" will only work on go.mod files, not go.work. :(

# As a total newcomer to Go, I will have to go with this clunky method for now
# and hope that a more experienced Go dev can teach me a better way.
COPY /datetime-utils/go.mod ./datetime-utils/go.mod
COPY /image-utils/go.mod ./image-utils/go.mod
COPY /image-utils/go.sum ./image-utils/go.sum
COPY /users/go.mod ./users/go.mod
COPY /web-service-gin/go.mod ./web-service-gin/go.mod
COPY /web-service-gin/go.sum ./web-service-gin/go.sum

RUN go mod download

COPY /datetime-utils/*.go ./datetime-utils/
COPY /image-utils/*.go ./image-utils/
COPY /users/*.go ./users/
COPY /web-service-gin/*.go ./web-service-gin/

WORKDIR /app/web-service-gin

EXPOSE 3000

RUN go build -o /web-service-gin


CMD [ "/web-service-gin" ]