.PHONY: dev build build_clean

dev:
	@go run cmd/server/main.go

build: build_clean
	@go build -o build/main cmd/server/main.go

build_clean:
	@rm -rf build
	@mkdir build
