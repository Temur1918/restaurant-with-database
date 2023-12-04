CREATE DATABASE restaurant;

CREATE TABLE table_ (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    number INTEGER UNIQUE NOT NULL
);

CREATE TABLE waiter (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE product(
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(8, 2) DEFAULT 0
);

CREATE TABLE order_ (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    is_paid BOOLEAN DEFAULT false,
    table_id VARCHAR(100) REFERENCES table_(id),
    waiter_id VARCHAR(100) REFERENCES waiter(id)
);

CREATE TABLE order_products(
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    quantity INTEGER DEFAULT 1,
    price NUMERIC(9, 2) DEFAULT 0,
    order_id VARCHAR(100) REFERENCES order_(id),
    product_id VARCHAR(100) REFERENCES product(id)
);

