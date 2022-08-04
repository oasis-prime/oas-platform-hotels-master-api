generate:
	- go run github.com/99designs/gqlgen generate
make-env:
	export $(grep -v '^#' .env | xargs)

dcup-local:
	docker-compose up

dcup-prod:
	docker-compose -f ./docker-compose.prod.yaml up

dc-down:
	docker-compose down

dc-clear:
	docker-compose down
	docker rmi -f ss-platform-api

deploy-trigger:
	gcloud functions deploy HelloAuth \
  --trigger-event providers/firebase.auth/eventTypes/user.create \
	--entry-point trigger/HelloAuth \
  --trigger-resource oas-platform \
  --runtime go116 \
	--set-env-vars APP_ENV="develop",DATABASE_NAME="oasis-master",DATABASE_READ_URL="oasis-trigger-event@cloudsql(oas-platform:asia-southeast1:oasis-prime)/oasis-master?charset=utf8&parseTime=True&loc=UTC",DATABASE_WRITE_URL="oasis-trigger-event@cloudsql(oas-platform:asia-southeast1:oasis-prime)/oasis-master?charset=utf8&parseTime=True&loc=UTC" \
	--region asia-southeast1

# hosts:
# 	sudo -- sh -c "echo 127.0.0.1  larler-dev.com >> /etc/hosts"
# 	sudo -- sh -c "echo 127.0.0.1  api.larler-dev.com >> /etc/hosts"
# 	sudo -- sh -c "echo 127.0.0.1  admin.larler-dev.com >> /etc/hosts"

# rm-hosts:
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 larler-dev.com/d' /etc/hosts"
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 api.larler-dev.com/d' /etc/hosts"
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 admin.larler-dev.com/d' /etc/hosts"