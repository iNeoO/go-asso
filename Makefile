ifneq (,$(wildcard .env))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# Normalize quoted env vars from .env (compose dislikes embedded quotes)
DB_USER_CLEAN := $(subst ",,$(DB_USER))
DB_PASSWORD_CLEAN := $(subst ",,$(DB_PASSWORD))
DB_NAME_CLEAN := $(subst ",,$(DB_NAME))

.PHONY: docker.network docker.build docker.run docker.postgres migrate.up migrate.down migrate.force seed seed.docker reset.db dev

docker.network:
	docker network inspect planigramme-network >/dev/null 2>&1 || \
	docker network create -d bridge planigramme-network

docker.build:
	docker build -t go-planigramme .

docker.run:
	docker run --rm -it \
		--network planigramme-network \
		-p 5000:5000 \
		go-planigramme

docker.postgres:
	docker run -d --rm \
    --name planigramme-postgres \
    --network planigramme-network \
    -e POSTGRES_USER=${DB_USER_CLEAN} \
    -e POSTGRES_PASSWORD=${DB_PASSWORD_CLEAN} \
    -e POSTGRES_DB=${DB_NAME_CLEAN} \
    -v planigramme-pgdata:/var/lib/postgresql \
    -p ${DB_PORT}:5432 \
    postgres

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database '$(MIGRATIONS_DB_URL)' up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database '$(MIGRATIONS_DB_URL)' down

migrate.force: VERSION ?= $(word 2,$(MAKECMDGOALS))
migrate.force:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make migrate.force VERSION=<n>  (or: make migrate.force <n>)"; \
		exit 1; \
	fi
	migrate -path $(MIGRATIONS_FOLDER) -database '$(MIGRATIONS_DB_URL)' force $(VERSION)

SWAG_VERSION ?= v1.16.4
swag:
	@[ -x "$(PWD)/bin/swag" ] || GOBIN="$(PWD)/bin" go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)
	PATH=$(PWD)/bin:$$PATH swag init -g cmd/api/main.go -o swagger

seed:
	psql '$(MIGRATIONS_DB_URL)' -f platform/seed/seed.sql

seed.docker: DB_HOST ?= planigramme-postgres
seed.docker: DB_PORT = 5432
seed.docker:
	docker run --rm -i \
		--network planigramme-network \
		-v $(PWD)/platform/seed:/seed \
		postgres \
		psql "postgres://${DB_USER_CLEAN}:${DB_PASSWORD_CLEAN}@${DB_HOST}:${DB_PORT}/${DB_NAME_CLEAN}" -f /seed/seed.sql

reset.db: DB_HOST ?= planigramme-postgres
reset.db:
	# Stop and remove any Postgres from compose or manual runs, plus its volume
	-docker compose -f docker-compose.dev.yml down -v --remove-orphans
	-docker rm -f planigramme-postgres
	-docker volume rm -f planigramme-pgdata
	DB_USER=$(DB_USER_CLEAN) DB_PASSWORD=$(DB_PASSWORD_CLEAN) DB_NAME=$(DB_NAME_CLEAN) docker compose -f docker-compose.dev.yml up -d postgres
	@echo "Waiting for postgres to be ready..."
	@until DB_USER=$(DB_USER_CLEAN) DB_PASSWORD=$(DB_PASSWORD_CLEAN) DB_NAME=$(DB_NAME_CLEAN) docker compose -f docker-compose.dev.yml exec -T postgres sh -c 'PGPASSWORD="$$DB_PASSWORD" pg_isready -U "$$DB_USER" -d "$$DB_NAME"' >/dev/null 2>&1; do \
		sleep 1; \
	done
	$(MAKE) migrate.up
	$(MAKE) seed.docker DB_HOST=$(DB_HOST)

# Allow extra arguments (e.g. migrate.force version) without breaking the build
EXTRA_GOALS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(EXTRA_GOALS):
	@:
