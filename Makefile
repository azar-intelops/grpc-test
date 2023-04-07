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