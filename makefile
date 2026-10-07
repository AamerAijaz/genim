build:
	powershell -NoProfile -Command "$$env:CGO_ENABLED = '0'; go build -o ./bin/genim.exe ./cmd"

run: build
	./bin/genim.exe
