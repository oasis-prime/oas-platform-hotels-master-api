generate:
	- go get github.com/99designs/gqlgen
	- go run github.com/99designs/gqlgen generate
make-env:
	export $(grep -v '^#' .env | xargs)

run-dev:
	-export $(grep -v '^#' .env | xargs);
	-air

dcup-local:
	docker-compose up

dcup-prod:
	docker-compose -f ./docker-compose.prod.yaml up

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