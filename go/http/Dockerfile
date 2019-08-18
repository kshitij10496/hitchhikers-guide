FROM golang:1.12

MAINTAINER Kshitij Saraogi <kshitij.saraogi@gmail.com>

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]