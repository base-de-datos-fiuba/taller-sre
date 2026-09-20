# ByteMarket · Hito 4: mínimo privilegio

Este hito agrega una segunda defensa: el backend deja de conectarse como administrador.

[`database/init/001_roles.sql`](database/init/001_roles.sql) crea el usuario de aplicación:

```sql
CREATE ROLE bytemarket_app
LOGIN PASSWORD 'classroom_demo_only';
```

[`migrations/002_grant_application_permissions.sql`](migrations/003_grant_application_permissions.sql)
otorga sólo acceso al schema y operaciones normales sobre tablas y secuencias. Luego
[`docker-compose.yml`](docker-compose.yml) configura:

```yaml
DB_USER: bytemarket_app
```

## Qué mejora

- La consulta sigue parametrizada.
- La aplicación puede consultar y modificar sus datos.
- El usuario del backend no puede crear ni eliminar el schema.

Estas defensas se complementan: limitar permisos no reemplaza las consultas parametrizadas.
