#!/bin/bash
source .env.local

export MIGRATION_DSN="host=postgresql port=5432 dbname=$POSTGRES_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_PATH}" postgres "${MIGRATION_DSN}" up -v
