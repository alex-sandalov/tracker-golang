#!/bin/bash

if [ ! -f ".env" ]; then
	echo ".env file not found."
	exit 1
else
	export $(grep -v '^#' .env | xargs)

	IFS=',' read -r -a db_array <<< "$DB_NAME"

	create_database() {
		local db=$1
		docker exec -it $CONTAINER_NAME psql -U $DB_USER -c "CREATE DATABASE $db;"
	}

	for db in "${db_array[@]}"; do
		create_database $db
	done
fi
