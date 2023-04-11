install:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

path:
	export PATH="$PATH:$(go env GOPATH)/bin"

protoc:
	mkdir pb && protoc --go_out=./pb --go-grpc_out=./pb proto/*.proto

clean:
	rm -rf pb

grpc_doesnt_work:
	source ~/.bash_profile

main:	
	go run main.go

test:
	go test -v ./...

evans:
	evans -r repl -p 50051

go_test_with_coverage:
	go test -coverprofile=coverage.out ./...

remove_coverage:
	rm -rf coverage.out

# check code coverage
test_code_coverage:
	go test -cover ./utils
