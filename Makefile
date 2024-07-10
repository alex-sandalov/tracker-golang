.DEFAULT_GOAL := help

help:
	@echo "Доступные цели"
	@echo "help           - справка"
	@echo "start          - запустить приложения"
	@echo "up             - запустить контейнеры"
	@echo "build          - собрать образы"
	@echo "create_dbs     - создать базы данных"
	@echo "network-up     - создать сеть"
	@echo "network-delete - удалить сеть"

build:
	docker-compose build db
	@echo "Образ базы данных собран"

up:
	docker-compose up --build -d db

create_dbs:
	chmod +x wait_for_postgres.sh
	chmod +x init_postgres.sh

	./wait_for_postgres.sh ./init_postgres.sh

	@echo "Базы данных postgres созданы"

start: up create_dbs
	@echo "Старт приложений"

network-up:
	docker network create app-net
	@echo "Сеть создана"

network-delete:
	docker network rm app-net
	@echo "Сеть удалена"

