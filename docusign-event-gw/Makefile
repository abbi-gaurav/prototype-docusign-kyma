RELEASE?=kyma-1.15
APP?=docusign-event-gw
GOOS?=linux
GOARCH?=amd64
CONTAINER_IMAGE?=gabbi/${APP}:${RELEASE}

clean:
	rm -rf ${APP}

vet:
	GO111MODULE=on go vet .

test:
	GO111MODULE=on go test -v -race ./...

build: clean vet test
	GO111MODULE=on CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w" \
		-o ${APP}

build-image: build
	docker build -t $(CONTAINER_IMAGE) .

push-image:
	docker push $(CONTAINER_IMAGE)

tidy:
	GO111MODULE=on go mod tidy