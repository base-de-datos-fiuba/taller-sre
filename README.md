# ByteMarket · Hito 5: auditoría de precios

La migración [`migrations/003_add_product_price_audit.sql`](migrations/003_add_product_price_audit.sql)
crea la tabla de auditoría, la función y el trigger para registrar cambios de precio.

```sql
CREATE TRIGGER trg_product_price_audit
AFTER UPDATE OF price ON product
FOR EACH ROW
WHEN (OLD.price IS DISTINCT FROM NEW.price)
EXECUTE FUNCTION log_product_price_change();
```

Cada cambio guarda `old_price`, `new_price`, `changed_at` y `CURRENT_USER`. El trigger registra una
fila por cada producto afectado, incluso si un único `UPDATE` modifica muchos productos.

## Probarlo

```sql
UPDATE product SET price = 1500 WHERE product_id = 3;
SELECT * FROM product_price_audit;
```

La búsqueda ya está parametrizada y el backend usa `bytemarket_app`. El rol tiene permisos de
escritura sobre las tablas para las operaciones normales de la aplicación.
