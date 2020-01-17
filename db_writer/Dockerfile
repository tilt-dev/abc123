FROM golang:1.10

ADD . /go/src/github.com/windmilleng/abc123/db_writer
RUN go install github.com/windmilleng/abc123/db_writer
ENTRYPOINT /go/bin/db_writer
