# syntax=docker/dockerfile:1

FROM golang:1.21.4-alpine as builder

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
COPY . ./
RUN go get ./...
RUN go mod download


RUN go build -o /standup-management-api
#RUN go build -tags netgo -ldflags '-s -w' -o app

EXPOSE 5000

CMD [ "/standup-management-api" ]