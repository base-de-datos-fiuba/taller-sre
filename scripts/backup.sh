#!/bin/sh
set -eu

project_root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
backup_dir="$project_root/backups"
backup_file="$backup_dir/bytemarket.dump"

mkdir -p "$backup_dir"

docker compose -f "$project_root/docker-compose.yml" exec -T db \
  pg_dump -U bytemarket_admin -d bytemarket -Fc > "$backup_file"

echo "Backup created at $backup_file"
