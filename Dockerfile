FROM alpine
WORKDIR /root
RUN apk --no-cache add ca-certificates
COPY moloon /go/bin/moloon

ENTRYPOINT ["/go/bin/moloon"]
EXPOSE 3000