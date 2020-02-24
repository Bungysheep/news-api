#!/bin/sh

migrationPath="$1"
databaseUrl="$2"

until psql "$databaseUrl" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec ./migrate -path "$migrationPath" -database "$databaseUrl" up