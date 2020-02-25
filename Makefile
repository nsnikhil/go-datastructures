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