FROM alpine

RUN adduser -D appuser
USER appuser

WORKDIR /appuser
COPY moloon /go/bin/moloon

ENTRYPOINT ["/go/bin/moloon"]
EXPOSE 3000