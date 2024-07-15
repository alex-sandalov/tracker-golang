# Тестовое задание
**Задача**: реализовать тайм-трекер. Если пользователя нет в базе данных, миграция которой описана в микросервисе `info`, то пользователь не может создать задачу.

## Описание взаимодействие между `Tracker` и `Info`
`Tracker` — это микросервис для управления пользователями. Он взаимодействует с сервисом `Info` для получения подробной информации о пользователях.

Внутри `Tracker` и `Info` лежит подробная информация как работать с микросервисами

## Стек технологий
- Golang
- Gin
- PostgreSQL
- Docker
- Swagger

## Запуск приложений
```sh
make build && make run
```

## Запуск первый раз
```sh
make build && make run && make migrate
```

## Цели Makefile
```sh
make help
```
