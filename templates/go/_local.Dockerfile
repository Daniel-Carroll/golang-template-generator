FROM us.gcr.io/heb-cx-utils/buildpack-golang-alp:latest

WORKDIR /usr/local/go/src/gitlab.com/heb-engineering/teams/enterprise/digital-fulfillment/jolly-roger/df-efc-fulfillment

COPY . /usr/local/go/src/gitlab.com/heb-engineering/teams/enterprise/digital-fulfillment/jolly-roger/df-efc-fulfillment

RUN apk --no-cache add tzdata

RUN go get github.com/githubnemo/CompileDaemon && go get -d -v ./...