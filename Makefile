COVERAGEOUT := coverage.out
SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

services-up:
	@docker-compose up -d

services-down:
	@docker-compose down

test:
	richgo test  ./... -v -coverprofile=$(COVERAGEOUT)
	go tool cover -func=$(COVERAGEOUT)

ci-test:
	go test  ./... -v -coverprofile=$(COVERAGEOUT)
	go tool cover -func=$(COVERAGEOUT)

cover:
	go test -v -coverprofile=$(COVERAGEOUT) ./...
	@go tool cover -html=$(COVERAGEOUT)

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-api

build-api:
	go build -tags $(LIBRARY_ENV) -o ./bin/api api/main.go

run:
	go run api/main.go

# export GOPATH=$HOME/go
# export PATH=$GOPATH/bin:$PATH