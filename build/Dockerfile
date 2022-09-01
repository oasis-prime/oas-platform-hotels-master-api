FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.* ./

RUN apk update && apk upgrade && \
    apk add git

#Set Env
ARG ACCESS_TOKEN
# ARG DATABASE_NAME=oasis-master
# ARG DATABASE_USER=root
# ARG DATABASE_PWD=iHA3i4JumM4Ap
# ARG DATABASE_INSTANCE_CONNECTION_NAME=127.0.0.1
# ARG DATABASE_SOCKET_DIR=true

# ARG JWT_SECRET=xxxx
# ARG JWT_EXPIRATIONDELTA=72

# ARG REDIS_ADDR=127.0.0.1:6379
# ARG REDIS_PASSWORD=xxxx
# ARG REDIS_DB=0

# ARG HOTELBEDS_ENDPOINT=https://api.test.hotelbeds.com
# ARG HOTELBEDS_SECUREENDOPINT=https://api-secure.test.hotelbeds.com
# ARG HOTELBEDS_KEY=rq9mtpfyjmedv77rdy3h7uhm
# ARG HOTELBEDS_SECRET=TqzaWFcvg6
# ARG HOTELBEDS_FORMAT=application/json

ENV ACCESS_TOKEN=$ACCESS_TOKEN

#Set github login
RUN go env -w GOPRIVATE=github.com/oasis-prime/oas-platform-core
RUN GOCACHE=OFF
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