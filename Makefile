NAME:=go-ddd
DC=docker-compose -f ./docker/docker-compose.yaml

tidy:
	rm -f go.sum; go mod tidy -compat=1.22.1

vet:
	go vet ./...

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

install-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.57.2

lint: install-linter
	./bin/golangci-lint run

stop:
	$(DC) stop

clean: stop
	$(DC) down --rmi local --remove-orphans -v
	$(DC) rm -f -v

build: clean
	$(DC) build

run: stop
	$(DC) up

helm-charts:
	helm lint helm/*
	helm package helm/*