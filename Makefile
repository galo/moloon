all: container

ENVVAR=GOARCH=amd64
TAG=1.0.0

build: clean
	$(ENVVAR) go build

build-linux: clean
	$(ENVVAR) GOOS=linux go build 

container:
	docker build \
	--build-arg http_proxy=$(http_proxy) \
	--build-arg https_proxy=$(https_proxy) \
	-t gcr.io/print-cloud-software/moloon:$(TAG) .

clean-container: build-linux
	docker build --no-cache --force-rm -t gcr.io/print-cloud-software/moloon:1.0.0:$(TAG) .

clean:
	rm -f moloon

.PHONY: all build container clean clean-container c
