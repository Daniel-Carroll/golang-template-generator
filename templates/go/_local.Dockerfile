FROM us.gcr.io/heb-cx-utils/buildpack-golang-alp:latest

WORKDIR /usr/local/go/src/{{ repoUrl }}

COPY . /usr/local/go/src/{{ repoUrl }}

RUN apk --no-cache add tzdata

RUN go get github.com/githubnemo/CompileDaemon && go get -d -v ./...