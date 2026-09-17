-- The bootstrap administrator owns the schema and runs migrations, but the
-- application connects with a separate role that cannot create or drop it.
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'bytemarket_app') THEN
        CREATE ROLE bytemarket_app
            LOGIN
            PASSWORD 'classroom_demo_only';
    END IF;
END
$$;

COMMENT ON ROLE bytemarket_app IS
    'Restricted login used by the ByteMarket backend';
