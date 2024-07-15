help:
	@echo "Доступные цели"
	@echo "help           - справка"
	@echo "run            - запустить приложения"
	@echo "build          - собрать образы"
	@echo "clean          - удалить контейнеры"
	@echo "migrate        - сделать миграцию базы данных"

build:
	docker-compose build info-golang

clean:
	docker-compose down info-golang

run:
	docker-compose up -d info-golang

stop:
	docker-compose stop info-golang

migrate:
	chmod +x ./run-migrations.sh
	./run-migrations.sh
