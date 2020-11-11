FROM golang:latest
RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work

RUN apt update -y && apt upgrade -y
RUN apt install default-mysql-client -y

WORKDIR /go
RUN go get github.com/oxequa/realize
RUN go get github.com/pressly/goose/cmd/goose
#RUN go get bitbucket.org/liamstask/goose/cmd/goose

EXPOSE 8080
WORKDIR /go/src/work
CMD ["realize", "start", "--server"]
