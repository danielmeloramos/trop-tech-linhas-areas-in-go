COVERAGEOUT := coverage.out

services-up:
	@docker-compose up -d

services-down:
	@docker-compose down

run: services-up
	gin -b linhas-areas -t cmd/linhas-areas/ -a 8889 -i run

test:
	richgo test  ./... -v -coverprofile=$(COVERAGEOUT)
	go tool cover -func=$(COVERAGEOUT)

ci-test:
	go test  ./... -v -coverprofile=$(COVERAGEOUT)
	go tool cover -func=$(COVERAGEOUT)

cover:
	go test -v -coverprofile=$(COVERAGEOUT) ./...
	@go tool cover -html=$(COVERAGEOUT)