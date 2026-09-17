# ByteMarket · Hito 2: backup y restore

> Aplicación educativa e intencionalmente vulnerable. Usar sólo en el entorno local descartable.

Sobre la aplicación vulnerable del hito anterior, esta branch incorpora un flujo reproducible de
respaldo y recuperación.

[`scripts/backup.sh`](scripts/backup.sh) genera un dump binario fuera del contenedor:

```sh
docker compose exec -T db \
  pg_dump -U bytemarket_admin -d bytemarket -Fc > "$backup_file"
```

[`scripts/restore.sh`](scripts/restore.sh) recrea el schema cuando fue eliminado y restaura el dump
con `pg_restore`. Los archivos generados quedan excluidos por
[`backups/.gitignore`](backups/.gitignore).

## Probarlo

```bash
docker compose up --build -d
./scripts/backup.sh
```

Después de la demostración destructiva:

```bash
./scripts/restore.sh
```

El objetivo no es solamente crear un archivo: es comprobar que el sistema puede recuperarse.
