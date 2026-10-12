ifeq ($(OS),Windows_NT)
    GIT_BASH := C:/Program Files/Git
    SHELL := $(GIT_BASH)/usr/bin/bash.exe
    .SHELLFLAGS := -c
    export PATH := $(GIT_BASH)/usr/bin:$(GIT_BASH)/mingw64/bin:$(PATH)
    export LANG := C.UTF-8
    export LC_ALL := C.UTF-8
	export MSYS_NO_PATHCONV := 1
endif

include .env
export

PROJECT_ROOT := $(CURDIR)
export PROJECT_ROOT

env-up:
	@docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Clean all environment volumes? Data loss risk. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		rm -rf "$(PROJECT_ROOT)/out/pgdata" && \
		echo "Environment files cleaned"; \
	else \
		echo "Cleanup cancelled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder


migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Required parameter 'seq' is missing. Example: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate create \
		-ext sql -dir /migrations -seq "$(seq)"	


migrate-up:
	@make migrate-action action=up


migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Required parameter 'action' is missing. Example: make migrate-create action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

logs-cleanup:
	@read -p "Clean all log files? Log loss risk. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf "$(PROJECT_ROOT)/out/logs" && \
		echo "Logs files cleaned"; \
	else \
		echo "Cleanup cancelled"; \
	fi

todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/todoapp/main.go

todoapp-deploy:
	@docker compose up -d --build todoapp

ps:
	@docker compose ps