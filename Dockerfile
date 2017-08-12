FROM golang
ADD . /go/src/github.com/englishcriminal/fizzbuzzapi
RUN go install github.com/englishcriminal/fizzbuzzapi
ENTRYPOINT /go/bin/fizzbuzzapi
EXPOSE 80
