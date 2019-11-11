FROM golang:1.13-alpine as builder

RUN apk update && apk add --virtual build-dependencies build-base gcc wget

WORKDIR $GOPATH/src/github.com/galo/moloon/
ADD . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/moloon

FROM alpine
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/moloon /go/bin/moloon

ENTRYPOINT ["/go/bin/moloon"]
EXPOSE 3000