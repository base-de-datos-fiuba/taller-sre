# ByteMarket · Hito 3: consulta parametrizada

Este hito corrige la SQL Injection de la búsqueda. En
[`backend/internal/products/search_products.go`](backend/internal/products/search_products.go), el
SQL y la entrada externa ahora viajan separados:

```go
query := `SELECT ... FROM product WHERE name ILIKE $1`
rows, err := r.db.Query(ctx, query, "%"+search+"%")
```

También se reemplaza `SearchVulnerable` por `Search` en
[`backend/internal/http/handler.go`](backend/internal/http/handler.go). Ya no se usa el protocolo
simple que permitía ejecutar sentencias apiladas.

## Verificar la mejora

```bash
docker compose up --build
```

Las búsquedas normales siguen funcionando. Una entrada que contiene SQL ahora se trata como texto
y no modifica el esquema.

Todavía queda una defensa pendiente: el backend continúa conectado con un usuario administrador.
