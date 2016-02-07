FROM golang

ADD main.go /go/src/goscripts/main.go

RUN go get goscripts

VOLUME /data

ENTRYPOINT ["/go/bin/goscripts"]
