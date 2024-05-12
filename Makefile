build:
	go build cmd/cli/main.go

run:
	go run cmd/cli/main.go

test:
	go test -v ./...

coverage:
	mkdir -p .coverage
	go test -v ./... -coverprofile=.coverage/coverage
	go tool cover -func=.coverage/coverage
