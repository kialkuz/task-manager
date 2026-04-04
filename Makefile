include .env

# Переменные
DOCKER_COMPOSE = docker compose
MIGRATE_SERVICE = migrate

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

migrate-up:
	$(DOCKER_COMPOSE) run --rm $(MIGRATE_SERVICE) -path /migrations -database $(DATABASE_URL) up

migrate-down:
	$(DOCKER_COMPOSE) run --rm $(MIGRATE_SERVICE) -path /migrations -database $(DATABASE_URL) down

migrate-create:
	$(DOCKER_COMPOSE) run --rm $(MIGRATE_SERVICE) create -ext sql -dir /migrations $(name)
