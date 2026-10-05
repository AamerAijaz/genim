build:
	go build -o ./bin/genim.exe genim.go

run:
	./bin/genim.exe

wasm:
	powershell -NoProfile -Command "$$env:GOOS = 'js'; $$env:GOARCH = 'wasm'; go build -o web/main.wasm ./cmd/wasm; Copy-Item -Force (Join-Path (go env GOROOT) 'lib\wasm\wasm_exec.js') 'web\wasm_exec.js'"