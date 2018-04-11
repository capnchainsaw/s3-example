FROM golang:stretch

MAINTAINER  James Johnson <tiasdungeon@gmail.com>

ENV GOPATH /local/go

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    python-pip

RUN pip install setuptools
RUN pip install awscli --upgrade

COPY s3x /s3x

RUN mkdir -p $GOPATH/src/github.com/capnchainsaw/s3-example
ADD . $GOPATH/src/github.com/capnchainsaw/s3-example
WORKDIR $GOPATH/src/github.com/capnchainsaw/s3-example

RUN go get .
RUN go build -o s3-example .

EXPOSE 8080
CMD ["./s3-example"]
