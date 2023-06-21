BINARY = seafarer
GOOS = linux
GOARCH = amd64
GO_BUILD_TAGS = ""
GO_BUILD_LDFLAGS = "-s -w"


.PHONY: run
run:
	go run main.go


.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -ldflags $(GO_BUILD_LDFLAGS) -tags $(GO_BUILD_TAGS) -trimpath -o ./bin/$(BINARY)_$(GOOS)_$(GOARCH) \
		 main.go


.PHONY: swaginit
swaginit:
	swag init


.PHONY: clean
clean:
	rm -rf bin logs


.PHONY: gomoddownload
gomoddownload:
	go mod download
