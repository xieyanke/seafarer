BINARY = seafarer
GOOS = linux
GOARCH = amd64
GO_BUILD_TAGS = ""

.PHONY: run
run:
	go run main.go


.PHONY: seafarer
seafarer:
	CGO_ENABLED=0 go build -trimpath -o ./bin/$(BINARY)_$(GOOS)_$(GOARCH) \
		-tags $(GO_BUILD_TAGS) ./cmd/main.go


.PHONY: gomoddownload
gomoddownload:
	go mod download

clean:
	rm -rf bin