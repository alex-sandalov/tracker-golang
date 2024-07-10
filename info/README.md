template

```sh
migrate create -ext sql -dir ./migrations -seq init
```

## Start
```sh
docker-compose up —build -d
chmod +x /run_migrations.sh
docker-compose down && docker-compose up —build -d
```