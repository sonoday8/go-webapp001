#!/bin/bash

source .env

GOOSE_DRIVER=$DB_DRIVER GOOSE_DBSTRING="$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST)/$DB_DATABASE" goose -dir db/migrations $@