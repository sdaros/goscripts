FROM golang

ADD goscripts.go /go/src/goscripts/goscripts.go

RUN go get goscripts

VOLUME /data

ENTRYPOINT ["/go/bin/goscripts"]
