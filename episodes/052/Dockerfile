FROM golang:latest
MAINTAINER Kris Nova "knova@heptio.com"

RUN mkdir -p /go/src/github.com/heptio/tgik/episodes/052
ADD . /go/src/github.com/heptio/tgik/episodes/052
WORKDIR /go/src/github.com/heptio/tgik/episodes/052
CMD ["go", "run", "main.go"]