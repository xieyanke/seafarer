BINARY = seafarer
GOOS = linux
GOARCH = amd64
GO_BUILD_TAGS = ""


.PHONY: run
run:
	go run cmd/main.go


.PHONY: build
build:
	CGO_ENABLED=0 go build -trimpath -o ./bin/$(BINARY)_$(GOOS)_$(GOARCH) \
		-tags $(GO_BUILD_TAGS) ./cmd/main.go


.PHONY: swaginit
swaginit:
	swag init -g cmd/main.go


.PHONY: clean
clean:
	rm -rf bin logs


.PHONY: gomoddownload
gomoddownload:
	go mod download
