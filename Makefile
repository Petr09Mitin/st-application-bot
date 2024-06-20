app-build:
	go build cmd/app.go

run:
	make app-build
	sudo ./app
