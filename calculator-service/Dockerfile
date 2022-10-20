FROM golang:1.19.2-alpine3.16

RUN apk add git

COPY go.mod /calculator/
COPY go.sum /calculator/

WORKDIR /calculator

RUN go mod download

COPY . /calculator

ENV CGO_ENABLED 0
ENV NUMBER_OF_THREADS 4

ENTRYPOINT ["./entrypoint.sh"]
