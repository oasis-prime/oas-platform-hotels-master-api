FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.* ./

RUN apk update && apk upgrade && \
    apk add git

#Set github login
RUN go env -w GOPRIVATE=github.com/oasis-prime/oas-platform-core
RUN GOCACHE=OFF
ARG GITHUB_USER
ARG ACCESS_TOKEN
ENV ACCESS_TOKEN=$ACCESS_TOKEN
ENV GITHUB_USER=$GITHUB_USER
RUN git config --global url."https://${ACCESS_TOKEN}:x-oauth-basic@github.com".insteadOf "https://github.com"

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o goapp cmd/api/main.go

##########################################

FROM alpine:3.13.1 as release

ARG APP_ENV
ARG VAULT_TOKEN
ARG APP_VERSION

ENV APP_ENV=${APP_ENV} \
    VAULT_TOKEN=${VAULT_TOKEN} \
    APP_VERSION=${APP_VERSION}

WORKDIR /app

COPY --from=builder /app/goapp ./goapp

CMD ["./goapp"]