ifneq (,$(wildcard .env))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

.PHONY: docker.network docker.build docker.run docker.postgres migrate.up migrate.down migrate.force dev

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
	docker run --rm -d \
    --name planigramme-postgres \
    --network planigramme-network \
    -e POSTGRES_USER=${DB_USER} \
    -e POSTGRES_PASSWORD=${DB_PASSWORD} \
    -e POSTGRES_DB=${DB_NAME} \
    -v ${HOME}/planigramme-postgres/data/:/var/lib/postgresql/data \
    -p ${DB_PORT}:5432 \
    postgres

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" force $(version)
