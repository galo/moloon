FROM golang:alpine as builder
ADD . /go/src/github.com/galo/moloon
RUN go install github.com/galo/moloon

FROM alpine
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/moloon /root/moloon

ENTRYPOINT /root/moloon server
EXPOSE 3000