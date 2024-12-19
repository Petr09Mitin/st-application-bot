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

.PHONY: mocks
MOCKS_DESTINATION=mocks
mocks: 
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $(shell find usecase -type f -name '*.go' ! -name '*_test.go'); do \
		mkdir -p $(MOCKS_DESTINATION)/`dirname $$file` && mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; \
	done
	@for file in $(shell find internal/rest/application_bot -type f -name '*.go' ! -name '*_test.go'); do \
		mkdir -p $(MOCKS_DESTINATION)/`dirname $$file | sed 's/internal\///'` && mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/`echo $$file | sed 's/internal\///'`; \
	done
	@echo "Mocks generated."

.PHONY: test
test:
	go test ./... -coverpkg ./... -coverprofile cover.out.tmp
	cat cover.out.tmp | grep -v "docs" | grep -v "mocks" | grep -v "proto" | grep -v "easyjson" > cover.out

.PHONY: coverage
coverage:
	go tool cover -func cover.out
