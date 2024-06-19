install-mockgen:
	go install go.uber.org/mock/mockgen@latest


mockgen:
	go generate ./...

test:
	go test ./...