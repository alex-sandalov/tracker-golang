#!/bin/sh
# wait-for-postgres.sh


if [ ! -f ".env" ]; then
	echo ".env file not found."
	exit 1
else
  export $(grep -v '^#' .env | xargs)
  
  set -e

  cmd="$@"

  until docker exec -it $CONTAINER_NAME psql -U $DB_USER -c "\q"; do
    >&2 echo "Postgres is unavailable - sleeping"
    sleep 1
  done

  >&2 echo "Postgres is up - executing command"
  exec $cmd

fi