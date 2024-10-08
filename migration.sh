#!/bin/bash
source .env

export MIGRATION_DSN="host=$DB_HOST port=5432 dbname=$POSTGRES_DB_NAME user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v