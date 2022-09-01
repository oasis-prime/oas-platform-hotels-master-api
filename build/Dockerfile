FROM golang:1.16-alpine as builder

ENV APP_HOME /go/src/app

WORKDIR /app

RUN set -ex; \
    apk update; \
    apk add --no-cache \
    git \
    gcc \
    musl-dev \
    make

COPY go.mod ./
COPY go.sum ./

#Set Env
ARG ACCESS_TOKEN

ENV ACCESS_TOKEN=$ACCESS_TOKEN

#Set github login
RUN go env -w GOPRIVATE=github.com/oasis-prime/oas-platform-core && \
    git config --global url."https://${ACCESS_TOKEN}:x-oauth-basic@github.com".insteadOf "https://github.com"

RUN go mod download

COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -o goapp ./cmd/api/main.go

##########################################

FROM alpine:3.13.1 as release

WORKDIR /app

COPY --from=builder ./app/goapp ./goapp

EXPOSE 8080

CMD ["./goapp"]