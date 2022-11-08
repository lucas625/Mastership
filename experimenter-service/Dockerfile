FROM golang:1.19.2-alpine3.16

RUN apk add git

COPY go.mod /experimenter/
COPY go.sum /experimenter/

WORKDIR /experimenter

RUN go mod download

COPY . /experimenter

ENV CGO_ENABLED 0
ENV NUMBER_OF_THREADS 4

ENTRYPOINT ["./entrypoint.sh"]
