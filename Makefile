SHELL := /bin/sh

COMPOSE_FILE := dev/docker/docker-compose.yml
COMPOSE := docker compose -f $(COMPOSE_FILE)

.PHONY: help up down restart logs ps build rebuild reset api sqlc swagger

help:
	@echo "Targets disponíveis:"
	@echo "  make up       - Sobe os containers em background"
	@echo "  make down     - Para e remove os containers"
	@echo "  make restart  - Reinicia os containers"
	@echo "  make logs     - Mostra logs (follow)"
	@echo "  make ps       - Lista status dos serviços"
	@echo "  make build    - Faz build das imagens"
	@echo "  make rebuild  - Recria os containers com build"
	@echo "  make reset    - Remove containers e volumes"
	@echo "  make api      - Sobe a API com variáveis do .env"
	@echo "  make sqlc     - Regenera o código sqlc a partir das queries SQL"
	@echo "  make swagger  - Regenera a documentação Swagger"

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down --remove-orphans

restart: down up

logs:
	$(COMPOSE) logs -f --tail=200

ps:
	$(COMPOSE) ps

build:
	$(COMPOSE) build

rebuild:
	$(COMPOSE) up -d --build --force-recreate

reset:
	$(COMPOSE) down -v --remove-orphans

api:
	@set -a; \
	if [ -f .env ]; then . ./.env; fi; \
	set +a; \
	go run ./cmd/api

sqlc:
	@echo ">> Iniciando geração de código sqlc..."
	@echo ">> Config: sqlc.yaml"
	@echo ">> Queries: core/database/postgres/query"
	@echo ">> Schema:  core/database/postgres/migration"
	@echo ">> Output:  core/database/postgres/query/sqlc"
	sqlc generate
	@echo ">> Código sqlc gerado com sucesso."

swagger:
	@echo ">> Iniciando geração da documentação Swagger..."
	@echo ">> Arquivo principal: main.go (dir: cmd/api)"
	@echo ">> Diretórios de busca: cmd/api e handlers/dtos"
	@echo ">> Saída: cmd/api/docs"
	swag init -g main.go \
		-d cmd/api,core/modules/resale/presentation/http,core/modules/resale/application/dto/output,core/modules/retailer/presentation/http,core/modules/retailer/application/dto/input,core/modules/retailer/application/dto/output \
		-o cmd/api/docs
	@echo ">> Documentação Swagger gerada com sucesso."
