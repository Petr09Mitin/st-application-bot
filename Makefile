docker-local:
	docker build -t petr09mitin/st-application-bot:latest -f cmd/Dockerfile . \
	&& docker build -t petr09mitin/st-application-bot:second-roster -f cmd/Dockerfile.sr . \
	&& docker-compose down \
	&& docker-compose up

app-build:
	docker build -t petr09mitin/st-application-bot:latest -f cmd/Dockerfile . --network=host

docker-run:
	make app-build && docker compose up -d

migrate:
	cd ./internal/repository/postgres/migrations \
	&& tern migrate --config ./tern.conf \
	&& tern migrate --config ./second_roster.conf \
	&& cd ../../../../
