FROM golang:1.19-alpine as builder

ENV APP_HOME /go/src/app

WORKDIR /app

RUN set -ex; \
    apk update; \
    apk add --no-cache \
    tzdata \
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
RUN go env -w GOPRIVATE=github.com/oasis-prime/oas-platform-core,github.com/oasis-prime/oas-platform-firebase-core && \
    git config --global url."https://${ACCESS_TOKEN}:x-oauth-basic@github.com".insteadOf "https://github.com"

RUN go mod download

COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o goapp ./cmd/api/main.go

RUN strip goapp

##########################################

FROM alpine as release

WORKDIR /app

RUN apk update && apk upgrade && \
    apk add git && \
    apk add build-base && \
    apk add tzdata

#Set Env
ARG GIN_MODE \
    APP_ENV \
    APP_PORT \
    DATABASE_NAME \
    DATABASE_USER \
    DATABASE_PWD \
    DATABASE_INSTANCE_CONNECTION_NAME \
    DATABASE_SOCKET_DIR \
    JWT_SECRET \
    JWT_EXPIRATIONDELTA \
    REDIS_ADDR \
    REDIS_PASSWORD \
    REDIS_DB \
    HOTELBEDS_ENDPOINT \
    HOTELBEDS_SECUREENDOPINT \
    HOTELBEDS_KEY \
    HOTELBEDS_SECRET \
    HOTELBEDS_FORMAT \
    GOOGLE_PLACES \
    GOOGLE_PUBSUBKEY \
    GOOGLE_PROJECTID \
    CHILLPAY_MD5 \
    CHILLPAY_MERCHANTCODE \
    CHILLPAY_APIKEY \
    CHILLPAY_URL \
    FIREBASE_KEY

ENV TZ=Asia/Bangkok

ENV GIN_MODE=$GIN_MODE \
    APP_ENV=$APP_ENV \
    APP_PORT=$APP_PORT \
    DATABASE_NAME=$DATABASE_NAME \
    DATABASE_USER=$DATABASE_USER \
    DATABASE_PWD=$DATABASE_PWD \
    DATABASE_INSTANCE_CONNECTION_NAME=$DATABASE_INSTANCE_CONNECTION_NAME \
    DATABASE_SOCKET_DIR=$DATABASE_SOCKET_DIR \
    JWT_SECRET=$JWT_SECRET \
    JWT_EXPIRATIONDELTA=$JWT_EXPIRATIONDELTA \
    REDIS_ADDR=$REDIS_ADDR \
    REDIS_PASSWORD=$REDIS_PASSWORD \
    REDIS_DB=$REDIS_DB \
    HOTELBEDS_ENDPOINT=$HOTELBEDS_ENDPOINT \
    HOTELBEDS_SECUREENDOPINT=$HOTELBEDS_SECUREENDOPINT \
    HOTELBEDS_KEY=$HOTELBEDS_KEY \
    HOTELBEDS_SECRET=$HOTELBEDS_SECRET \
    HOTELBEDS_FORMAT=$HOTELBEDS_FORMAT \
    GOOGLE_PLACES=$GOOGLE_PLACES \
    GOOGLE_PUBSUBKEY=$GOOGLE_PUBSUBKEY \
    GOOGLE_PROJECTID=$GOOGLE_PROJECTID \
    CHILLPAY_MD5=$CHILLPAY_MD5 \
    CHILLPAY_MERCHANTCODE=$CHILLPAY_MERCHANTCODE \
    CHILLPAY_APIKEY=$CHILLPAY_APIKEY \
    CHILLPAY_URL=$CHILLPAY_URL \
    FIREBASE_KEY=$FIREBASE_KEY

COPY --from=builder ./app/goapp ./goapp

EXPOSE 8080

CMD ["./goapp"]