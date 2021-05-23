# Build Stage
FROM us.gcr.io/constellation-utils/golang-alp:latest as build-go

WORKDIR /usr/local/go/src/{{ module }}

ARG CURRENT_ENVIRONMENT=$ENVIRONMENT

COPY . /usr/local/go/src/{{ repo_url }}

# Compiler flags are set. Reference: http://www.jeffsloyer.io/post/cross-compiling-docker-alpine-golang/
RUN go get -d -v ./... \
    && cd /usr/local/go/src/{{ repo_url }}/cmd/fulfillment/ \
    && GOOS=linux go build -ldflags "-s" -a -installsuffix cgo .

# Final Stage
FROM alpine 

# Certificates needed for https to google function
RUN apk add --update --no-cache ca-certificates \
    && mkdir /app

RUN apk --no-cache add tzdata

WORKDIR /app

COPY --from=build-go /usr/local/go/src/{{ repo_url }}/cmd/fulfillment .
COPY --from=build-go /usr/local/go/src/{{ repo_url }}/scripts/start-app .

CMD ["sh", "-c", "./start-app"]