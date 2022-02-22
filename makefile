.PHONY: all

all: clean build run

.PHONY: clean

clean:
	@rm -rf ./helloworld.app ./public
	@echo "[✔️] Clean complete!"

.PHONY: build

build:
	@npm install
	@npm run prod
	@mkdir -p ./helloworld.app/Contents/MacOS/
	@go build -o ./helloworld.app/Contents/MacOS/helloworld
	@cp .env ./helloworld.app/Contents/MacOS/
	@cp config.yml ./helloworld.app/Contents/MacOS/
	@cp -R public ./helloworld.app/Contents/MacOS/public
	@echo "[✔️] Build complete!"

.PHONY: run

run:
	@open ./helloworld.app
	@echo "[✔️] App is running!"



include .env
export

RUNNER=docker-compose exec app sql-migrate

ifeq ($(p),host)
 	RUNNER=sql-migrate
endif

MIGRATE=$(RUNNER)

migrate-status:
	$(MIGRATE) status

migrate-up:
	$(MIGRATE) up
	sqlboiler psql --wipe

migrate-down:
	$(MIGRATE) down
	sqlboiler psql --wipe

redo:
	@read -p  "Are you sure to reapply the last migration? [y/n]" -n 1 -r; \
	if [[ $$REPLY =~ ^[Yy] ]]; \
	then \
		$(MIGRATE) redo; \
	fi

create:
	@read -p  "What is the name of migration?" NAME; \
	${MIGRATE} new $$NAME

.PHONY: status migrate-up migrate-down migrate redo create