mkdir -p ../../../../packages/varsoapi/src/app

echo src/services/analytics/internal/proto: Installing protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

echo src/services/analytics/internal/proto: Generating files
protoc -I=./ --plugin=$(go env GOPATH)/bin/protoc-gen-go --plugin=../../../../common/proto/node_modules/.bin/protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true --go_out=../../ --ts_proto_out=../../../../packages/varsoapi/src/analytics analyticsbase.proto analyticsrequests.proto

mkdir -p ../../../www/src/proto/src/analytics

if [ -e "../../../../packages/varsoapi/src/analytics" ] ; then 
  cp -R ../../../../packages/varsoapi/src/analytics ../../../www/src/proto/src/
fi
