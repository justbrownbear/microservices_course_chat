#!/bin/bash
source .env

export MIGRATION_DSN="host=postgresql port=$POSTGRES_PORT dbname=$POSTGRES_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

sleep 2 && /bin/goose -dir /migrations postgres "${MIGRATION_DSN}" up -v
