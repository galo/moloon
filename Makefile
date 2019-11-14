all: container

ENVVAR=GOARCH=amd64 CGO_ENABLED=0
TAG=1.0.0

build: clean
	$(ENVVAR) go build

build-linux: clean
	$(ENVVAR) GOOS=linux go build

container: build-linux
	docker build \
	--build-arg http_proxy=$(http_proxy) \
	--build-arg https_proxy=$(https_proxy) \
	-t r.jdkr.io/galo/moloon:$(TAG) .

clean-container: build-linux
	docker build --no-cache --force-rm -t r.jdkr.io/galo/moloon:$(TAG) .

clean:
	rm -f moloon

.PHONY: all build container clean clean-container c