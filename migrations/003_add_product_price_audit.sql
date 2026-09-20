CREATE TABLE product_price_audit (
    audit_id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    old_price NUMERIC(12, 2),
    new_price NUMERIC(12, 2),
    changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    changed_by TEXT NOT NULL
);

CREATE OR REPLACE FUNCTION log_product_price_change()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO product_price_audit (
        product_id,
        old_price,
        new_price,
        changed_by
    ) VALUES (
        OLD.product_id,
        OLD.price,
        NEW.price,
        CURRENT_USER
    );

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_product_price_audit
AFTER UPDATE OF price
ON product
FOR EACH ROW
WHEN (OLD.price IS DISTINCT FROM NEW.price)
EXECUTE FUNCTION log_product_price_change();
