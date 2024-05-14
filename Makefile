build:
	@cp "$(GOROOT)/misc/wasm/wasm_exec.js" ./docs/wasm_exec.js
	@env GOOS=js GOARCH=wasm go build -o ./docs/game.wasm ./

start:
	@go run ./server/main.go
