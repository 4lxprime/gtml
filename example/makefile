build-server:
	@go build -o srv server/main.go

server: build-server
	@./srv

build-wasm:
	@GOOS=js GOARCH=wasm go build -o app/static/main.wasm app/main.go