.PHONY: dev build build_clean dev_up dev_down dev_clean

# Color section
# Ref: https://qastack.vn/programming/5947742/how-to-change-the-output-color-of-echo-in-linux
GREEN=\033[0;32m
NC=\033[0m # No Color

dev:
	@go run cmd/server/main.go

build: build_clean
	@go build -o build/main cmd/server/main.go

build_clean:
	@rm -rf build
	@mkdir build

dev_up:
	@docker compose -f deployments/dev/docker-compose.yaml up -d

dev_down:
	@docker compose -f deployments/dev/docker-compose.yaml down

dev_docker_volume_name := $(shell docker volume ls -q | grep "shape_db_volume")
dev_db_clean:
	@if [ ! -z "$(dev_docker_volume_name)" ]; then\
		echo "$(GREEN)> Removing docker volume ...$(NC)";\
		docker volume rm $(dev_docker_volume_name);\
	fi
	@echo "$(GREEN)> Cleaned!$(NC)";
