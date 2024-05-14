build:
	@cp "$(GOROOT)/misc/wasm/wasm_exec.js" ./assets/wasm_exec.js
	@env GOOS=js GOARCH=wasm go build -o ./assets/game.wasm ./

start:
	@go run ./server/main.go
