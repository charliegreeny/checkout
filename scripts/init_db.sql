#!/bin/bash

CREATE DATABASE checkout;

USE checkout; 

CREATE TABLE offers(
    id VARCHAR(256), 
    quantity INTEGER,
    price INTEGER,
    discount DECIMAL(3,2) NULL,
    PRIMARY KEY (id)
);

CREATE TABLE items(
    SKU VARCHAR(256),
    price INTEGER,
    offers_id VARCHAR(256),
    PRIMARY KEY (SKU),
    CONSTRAINT fk_itemOffer FOREIGN KEY (offers_id) REFERENCES offers(id)
);


CREATE TABLE customers(
    id VARCHAR(256), 
    PRIMARY KEY (id)
);

CREATE TABLE carts (
    id VARCHAR(256) UNIQUE,
    customers_id VARCHAR(256),
    cart_line_items_id VARCHAR(256),
    updated_at TIMESTAMP,
    created_at TIMESTAMP,
    is_complete BOOLEAN, 
    PRIMARY KEY (id, customers_id),
    CONSTRAINT fk_cartCustomer FOREIGN KEY (customers_id) REFERENCES customers(id)
);

CREATE TABLE cart_line_items(
    cart_id VARCHAR(256),
    items_SKU VARCHAR(256),
    quantity INTEGER,
    PRIMARY KEY (cart_id, items_SKU, quantity),
    CONSTRAINT fk_cartLineItemItems FOREIGN KEY (items_SKU) REFERENCES items(SKU),
    CONSTRAINT fk_cartCartLineItem FOREIGN KEY (cart_id) REFERENCES carts(id)
);

INSERT INTO offers(id, price, quantity) VALUES('123', 130, 5);
INSERT INTO offers(id, price, quantity) VALUES('456', 2, 45);

INSERT INTO items VALUES('A', 50, '123'); 
INSERT INTO items VALUES('B', 30, '456'); 
INSERT INTO items(SKU, price) VALUES('C', 20);
INSERT INTO items(SKU, price)  VALUES('D', 15);
 
INSERT INTO customers VALUES('customer_1');
INSERT INTO customers VALUES('customer_2');
