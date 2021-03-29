FROM golang:1.16.0-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

# RUN set -ex; \
#     apk update; \
#     apk add --no-cache \
#     git \
#     gcc \
#     musl-dev \
#     librdkafka-dev \
#     pkgconf \
#     openssh-client \
#     ca-certificates && \
#     echo "machine gitlab.com login ${GITLAB_USERNAME} password ${GITLAB_TOKEN}" > ~/.netrc

RUN go mod download

COPY . .

# RUN go test ./... --cover -v -tags musl

RUN GOOS=linux GOARCH=amd64 go build -o goapp main.go

FROM alpine:3.12.0 as release

WORKDIR /app

COPY --from=builder /app/goapp ./goapp

EXPOSE 3000

CMD ["./goapp"]
