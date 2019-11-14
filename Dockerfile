FROM ubuntu

## Dev tool set - remove on prod
RUN apt-get update && apt-get -y install curl

RUN groupadd -g 999 appuser && \
    useradd -r -u 999 -g appuser appuser
USER appuser

WORKDIR /appuser
COPY moloon /go/bin/moloon

ENTRYPOINT ["/go/bin/moloon"]
EXPOSE 3000