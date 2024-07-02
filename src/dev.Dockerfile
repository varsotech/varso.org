FROM cosmtrek/air AS goapp

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apt upgrade && apt install -y git make openssh-client curl
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app 
ENV airwd=/app

ARG SERVICE

ENV SERVICE="${SERVICE}"
# dlv exec --listen=:2345 --log 
ENTRYPOINT air --build.include_dir "src/services/${SERVICE},src/common,src/services/auth/client,src/services/app/client,src/services/fileserver/client" --build.cmd "go build -gcflags \"all=-N -l\" -o /app/tmp/${SERVICE} src/services/${SERVICE}/main.go" --build.bin "/app/tmp/${SERVICE}"

FROM goapp AS goappapi
ARG PORT 
EXPOSE ${PORT}
