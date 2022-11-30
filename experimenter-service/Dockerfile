FROM golang:1.19.2-alpine3.16

RUN apk add git

COPY go.mod /experimenter/
COPY go.sum /experimenter/

WORKDIR /experimenter

RUN go mod download

COPY . /experimenter

ENV CGO_ENABLED 0
ENV NUMBER_OF_THREADS 4

ENV MSC_CALCULATOR_SERVICE_ADDRESS=mastership-calculator:8001
ENV MSC_EXPERIMENTER_SERVICE_PORT=8000

ENTRYPOINT ["./entrypoint.sh"]
