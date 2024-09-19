CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          base_price NUMERIC(10, 2) NOT NULL
);

INSERT INTO products (name, base_price) VALUES ('Product1', 100.00);
INSERT INTO products (name, base_price) VALUES ('Product2', 200.00);