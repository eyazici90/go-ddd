NAME:=go-ddd


tidy:
	rm -f go.sum; go mod tidy -compat=1.17

vet:
	go vet ./...

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

install-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.57.2

lint: install-linter
	./bin/golangci-lint run