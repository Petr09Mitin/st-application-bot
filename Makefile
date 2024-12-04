app-build:
	docker build -t st-application-bot/main -f cmd/Dockerfile .

docker-run:
	docker-compose up -d

migrate:
	cd ./internal/repository/postgres/migrations && tern migrate && cd ../../../../
