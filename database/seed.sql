INSERT INTO customer (first_name, last_name, email) VALUES
    ('Alex', 'Rivera', 'alex.rivera@example.test'),
    ('Sam', 'Chen', 'sam.chen@example.test'),
    ('Jordan', 'Patel', 'jordan.patel@example.test'),
    ('Taylor', 'Kim', 'taylor.kim@example.test'),
    ('Casey', 'Smith', 'casey.smith@example.test'),
    ('Morgan', 'Diaz', 'morgan.diaz@example.test'),
    ('Riley', 'Brown', 'riley.brown@example.test'),
    ('Jamie', 'Wilson', 'jamie.wilson@example.test'),
    ('Avery', 'Martin', 'avery.martin@example.test'),
    ('Cameron', 'Lee', 'cameron.lee@example.test');

INSERT INTO product (name, description, price, stock) VALUES
    ('Mechanical Keyboard', 'Hot-swappable switches and white backlight', 89.99, 18),
    ('Gaming Mouse', 'Lightweight mouse with adjustable DPI', 49.90, 31),
    ('USB-C Hub', 'Seven-port hub with HDMI and card reader', 39.50, 24),
    ('Laptop Stand', 'Adjustable aluminum desktop stand', 42.00, 12),
    ('Webcam', '1080p webcam with privacy shutter', 55.75, 20),
    ('27-inch Monitor', 'IPS display with 1440p resolution', 279.00, 8),
    ('Wireless Headphones', 'Over-ear headphones with active noise cancellation', 129.95, 15),
    ('Portable SSD 1TB', 'Fast USB-C solid-state storage', 99.00, 17),
    ('Desk Microphone', 'USB condenser microphone with desk stand', 74.50, 9),
    ('Wi-Fi 6 Router', 'Dual-band router for home networks', 119.00, 11),
    ('Bluetooth Speaker', 'Compact water-resistant speaker', 45.00, 26),
    ('Ergonomic Mouse Pad', 'Wrist-supported desk mouse pad', 18.25, 40),
    ('DisplayPort Cable', 'Two-meter DisplayPort 1.4 cable', 16.90, 36),
    ('Smart LED Bulb', 'Color-changing Wi-Fi LED bulb', 14.99, 50),
    ('Raspberry Pi Case', 'Ventilated case with cooling fan', 22.00, 14),
    ('Power Bank', '20,000 mAh USB-C portable charger', 48.80, 19),
    ('Cable Organizer', 'Magnetic cable clips, pack of six', 12.50, 45),
    ('Numeric Keypad', 'Slim wireless numeric keypad', 27.30, 16),
    ('Surge Protector', 'Eight outlets with USB charging', 34.99, 22),
    ('Screen Cleaning Kit', 'Microfiber cloth and alcohol-free spray', 11.50, 33);

INSERT INTO orders (status, customer_id) VALUES
    ('PAID', 1),
    ('PENDING', 3),
    ('SHIPPED', 6),
    ('CANCELLED', 8);

INSERT INTO order_item (order_id, product_id, quantity, unit_price) VALUES
    (1, 1, 1, 89.99),
    (1, 2, 1, 49.90),
    (2, 3, 2, 39.50),
    (3, 6, 1, 279.00),
    (3, 13, 1, 16.90),
    (4, 7, 1, 129.95);
