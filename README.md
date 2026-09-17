# ByteMarket · Hito 1: aplicación vulnerable

> Aplicación educativa e intencionalmente vulnerable. Usar sólo en el entorno local descartable.

Este hito presenta la aplicación inicial: React → Go → PostgreSQL. El esquema base está en
[`migrations/001_create_initial_schema.sql`](migrations/001_create_initial_schema.sql) y contiene
clientes, productos, pedidos y sus ítems.

La búsqueda concatena directamente la entrada externa en
[`backend/internal/products/search_products_vulnerable.go`](backend/internal/products/search_products_vulnerable.go):

```go
query := `SELECT ... FROM product WHERE name ILIKE '%` + search + `%'`
```

Además, el backend usa `bytemarket_admin`, un usuario con privilegios excesivos. La combinación
permite mostrar por qué una SQL Injection puede tener un impacto tan grande.

## Probarlo

```bash
docker compose up --build
```

Abrir <http://localhost:5173> y buscar `keyboard`, `mouse` o `monitor`.

Archivos clave:

- [`docker-compose.yml`](docker-compose.yml): servicios y credenciales iniciales.
- [`backend/internal/http/handler.go`](backend/internal/http/handler.go): endpoint de productos.
- [`database/seed.sql`](database/seed.sql): datos para la demostración.
