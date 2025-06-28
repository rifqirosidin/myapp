ENV ?= development
COMPOSE = docker-compose -f docker-compose.yml

ifeq ($(ENV), production)
	COMPOSE += --f docker-compose.override.yml --env-file .env.production
else
	COMPOSE += -f docker-compose.override.yml --env-file .env.development
endif

up:
	$(COMPOSE) up --build

down:
	$(COMPOSE) down

logs:
	$(COMPOSE) logf -f

restart: down up

ps:
	$(COMPOSE) ps

deploy:
	./scripts/deploy.sh