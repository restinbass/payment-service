#!/bin/bash

set -o allexport
source .env
set +o allexport

COMMAND=$1

if [ -z "${PORT}" ]; then
    PORT=5432
fi

export GOOSE_DRIVER="postgres"
export GOOSE_DBSTRING="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

goose -dir migrations $COMMAND
