#!/bin/sh
set -eu

project_root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
backup_file="$project_root/backups/bytemarket.dump"

if [ ! -f "$backup_file" ]; then
  echo "Backup not found: $backup_file" >&2
  exit 1
fi

docker compose -f "$project_root/docker-compose.yml" exec -T db \
  psql -v ON_ERROR_STOP=1 -U bytemarket_admin -d bytemarket \
  -c "CREATE SCHEMA IF NOT EXISTS public"

docker compose -f "$project_root/docker-compose.yml" exec -T db \
  pg_restore -U bytemarket_admin -d bytemarket --exit-on-error < "$backup_file"

echo "Restore completed from $backup_file"
