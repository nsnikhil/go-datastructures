ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

test:
	go clean -testcache
	go test ./...

test-cover-html:
	go clean -testcache
	mkdir -p out/
	go test ./... -coverprofile=out/coverage.out
	go tool cover -html=out/coverage.out

clean:
	rm -rf out/

fmt:
	gofmt -l -s -w .

vet:
	go vet ./...

setup: download deps test

download:
	go get -u golang.org/x/lint/golint

deps:
	go mod vendor

lint:
	golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; }
