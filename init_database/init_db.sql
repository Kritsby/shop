CREATE SCHEMA IF NOT EXISTS shop;

CREATE TABLE IF NOT EXISTS shop.storages (
                                             id INT UNIQUE NOT NULL PRIMARY KEY,
                                             name TEXT,
                                             is_availability bool
);

CREATE TABLE IF NOT EXISTS shop.products (
                                             name TEXT,
                                             size INT,
                                             id INT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS shop.product_amount (
                                                   storage_id INT,
                                                   product_id INT,
                                                   amount INT,
                                                   reserved INT DEFAULT 0
);

INSERT INTO shop.storages(id, name, is_availability)VALUES ('1', 'Москва', TRUE);
INSERT INTO shop.storages(id, name, is_availability)VALUES ('2', 'Воронеж', TRUE);
INSERT INTO shop.storages(id, name, is_availability)VALUES ('3', 'Владивосток', TRUE);

INSERT INTO shop.products(name, size, id) VALUES('футболка', 51, 1);
INSERT INTO shop.products(name, size, id) VALUES('футболка', 52, 2);
INSERT INTO shop.products(name, size, id) VALUES('футболка', 53, 3);

INSERT INTO shop.products(name, size, id) VALUES('джинсы', 51, 11);
INSERT INTO shop.products(name, size, id) VALUES('джинсы', 52, 12);
INSERT INTO shop.products(name, size, id) VALUES('джинсы', 53, 13);

INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('1', 1, 7);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('1', 2, 3);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('1', 3, 9);

INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('1', 11, 5);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('1', 12, 6);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('1', 13, 2);

INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('2', 1, 11);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('2', 2, 1);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('2', 3, 0);

INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('2', 11, 14);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('2', 12, 5);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('2', 13, 8);

INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('3', 1, 2);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('3', 2, 0);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('3', 3, 6);

INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('3', 11, 6);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('3', 12, 1);
INSERT INTO shop.product_amount(storage_id, product_id, amount)VALUES ('3', 13, 8);
