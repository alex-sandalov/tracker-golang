.DEFAULT_GOAL := help

help:
	@echo "Доступные цели"
	@echo "help           - справка"
	@echo "run            - запустить приложения"
	@echo "up             - запустить контейнеры"
	@echo "build          - собрать образы"
	@echo "create_dbs     - создать базы данных"
	@echo "network-up     - создать сеть"
	@echo "network-delete - удалить сеть"

up:
	docker-compose up --build -d tracker-db

create_dbs:
	chmod +x wait_for_postgres.sh
	chmod +x init_postgres.sh

	./wait_for_postgres.sh ./init_postgres.sh

	@echo "Базы данных postgres созданы"

build: up create_dbs
	cd tracker && make build
	@echo "Образ tracker собран"

	cd info && make build
	@echo "Образ info собран"

run:
	docker-compose up -d tracker-db
	cd tracker && make run
	cd info && make run

clean:
	docker-compose down tracker-db
	cd tracker && make clean
	cd info && make clean

migrate:
	cd tracker && make migrate
	cd info && make migrate

network-up:
	docker network create app-net
	@echo "Сеть создана"

network-delete:
	docker network rm app-net
	@echo "Сеть удалена"
