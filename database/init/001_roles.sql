-- bytemarket_admin is created by the official PostgreSQL image from
-- POSTGRES_USER. It is intentionally a superuser for the first lesson.
-- A later lesson can create a restricted login role in this file and grant
-- only SELECT on the application tables.
COMMENT ON ROLE bytemarket_admin IS
    'INTENTIONALLY OVERPRIVILEGED role for the disposable classroom demo';
