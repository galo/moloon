all: container

ENVVAR=GOARCH=amd64 CGO_ENABLED=0
TAG=1.0.0

build: clean
	$(ENVVAR) go build

build-linux: clean
	$(ENVVAR) GOOS=linux go build

container: build-linux
	docker build -t galo/moloon:$(TAG) .

clean-container: build-linux
	docker build --no-cache --force-rm -t galo/moloon:$(TAG) .

clean:
	rm -f moloon

.PHONY: all build container clean clean-container c