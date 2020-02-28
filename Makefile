help:
	@echo "make help               Show this help"
	@echo "make build              Build the binaries in build directory"
	@echo "make test               Run the unit tests"
	@echo "make test-integration   Run the all the tests, including integration tests"

build: test
	GOOS=linux GOARCG=amd64 go build -o ./build/linux_amd64/opype .
	GOOS=darwin GOARCG=amd64 go build -o ./build/darwin_amd64/opype . 
	GOOS=windows GOARCG=amd64 go build -o ./build/windows_amd64/opype.exe . 

test:
	go fmt ./... && go vet ./... && golint --set_exit_status ./... && go test ./...

test_integration: test
	go test -tags=integration