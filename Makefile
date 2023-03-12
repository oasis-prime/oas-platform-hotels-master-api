gen:
	- go get github.com/99designs/gqlgen
	- go run github.com/99designs/gqlgen generate
bu:
	- go build -ldflags="-s -w" cmd/api/main.go
	- strip main 
make-env:
	export $(grep -v '^#' .env | xargs)

run-dev:
	-export $(grep -v '^#' .env | xargs);
	-air

set-mod:
	go env -w GOPRIVATE=github.com/oasis-prime/oas-platform-core,github.com/oasis-prime/oas-platform-firebase-core
	git config --global url."https://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:x-oauth-basic@github.com".insteadOf "https://github.com"

dcup-build:
	docker build \
		--build-arg ACCESS_TOKEN= \
		-t oas-platform-hotels-master-api -f ./build/Dockerfile .

dcup-local:
	docker-compose up

dcup-prod:
	docker-compose -f ./docker-compose.prod.yaml up --build

dc-down:
	docker-compose down

dc-clear:
	docker-compose down
	docker rmi -f ss-platform-api

# hosts:
# 	sudo -- sh -c "echo 127.0.0.1  larler-dev.com >> /etc/hosts"
# 	sudo -- sh -c "echo 127.0.0.1  api.larler-dev.com >> /etc/hosts"
# 	sudo -- sh -c "echo 127.0.0.1  admin.larler-dev.com >> /etc/hosts"

# rm-hosts:
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 larler-dev.com/d' /etc/hosts"
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 api.larler-dev.com/d' /etc/hosts"
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 admin.larler-dev.com/d' /etc/hosts"
