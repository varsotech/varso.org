set -e
sudo apt install protobuf-compiler
go mod download
go generate ./...