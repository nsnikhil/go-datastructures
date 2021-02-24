APP=erx
APP_VERSION:=0.1
APP_COMMIT:=$(shell git rev-parse HEAD)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

deps:
	go mod download
tidy:
	go mod tidy

check: fmt vet lint

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

lint:
	golint $(ALL_PACKAGES)

test:
	go clean -testcache
	go test ./...

test-cover-html:
	go clean -testcache
	mkdir -p out/
	go test ./... -coverprofile=out/coverage.out
	go tool cover -html=out/coverage.out