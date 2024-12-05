app-build:
	docker build -t petr09mitin/st-application-bot:latest -f cmd/Dockerfile . --network=host

docker-run:
	make app-build && docker compose up -d

migrate:
	cd ./internal/repository/postgres/migrations && tern migrate && cd ../../../../
