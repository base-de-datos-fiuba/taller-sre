#!/bin/sh
set -eu

run_sql() {
  echo "Running $1"
  psql -v ON_ERROR_STOP=1 \
    --username "$POSTGRES_USER" \
    --dbname "$POSTGRES_DB" \
    --file "$1"
}

run_sql /bytemarket-init/001_roles.sql

for migration in /bytemarket-init/migrations/*.sql; do
  run_sql "$migration"
done

run_sql /bytemarket-init/seed.sql
