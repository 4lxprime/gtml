build:
	go build -o bin/app cmd/main.go

run: build
	./cmd/app