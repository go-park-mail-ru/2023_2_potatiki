DROP TABLE IF EXISTS profile CASCADE;

DROP TABLE IF EXISTS product CASCADE;

DROP TABLE IF EXISTS order_info CASCADE;

DROP TABLE IF EXISTS order_item;

DROP TABLE IF EXISTS category CASCADE;

DROP TABLE IF EXISTS category_reference;

DROP TABLE IF EXISTS shopping_cart_item;

DROP TABLE IF EXISTS promocode;

DROP TABLE IF EXISTS address;

DROP TABLE IF EXISTS favorite;

CREATE TABLE IF NOT EXISTS profile
(
    id uuid NOT NULL PRIMARY KEY,
    login text NOT NULL UNIQUE,
    description text,
    imgsrc text NOT NULL DEFAULT 'default.png',
    passwordhash text NOT NULL,
    CONSTRAINT "ProfileLogin_unique" UNIQUE (login)
    );

CREATE TABLE IF NOT EXISTS category
(
    id UUID NOT NULL PRIMARY KEY,
    name TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS product
(
    id uuid NOT NULL PRIMARY KEY,
    name_product text NOT NULL,
    description text NOT NULL,
    price INT NOT NULL,
    imgsrc text NOT NULL,
    rating NUMERIC(3, 2) NOT NULL,
    category_id UUID,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE RESTRICT,
    CONSTRAINT "ProductName_unique" UNIQUE (name_product),
    CHECK (rating >= 0),
    CHECK (price > 0)
    );

CREATE TABLE IF NOT EXISTS promocode
(
    id UUID NOT NULL PRIMARY KEY,
    discount INT NOT NULL,
    name TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS order_info
(
    id UUID NOT NULL PRIMARY KEY,
    delivery_date TIMESTAMP,
    creation_date TIMESTAMP,
    profile_id UUID,
    status TEXT,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    promocode_id UUID,
    FOREIGN KEY (promocode_id) REFERENCES promocode(id) ON DELETE RESTRICT
    );


CREATE TABLE IF NOT EXISTS order_item
(
    id UUID NOT NULL PRIMARY KEY,
    order_id UUID,
    FOREIGN KEY (order_id) REFERENCES order_info(id) ON DELETE CASCADE,
    product_id UUID,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    quantity INT,
    CHECK (quantity >= 0)
    );


CREATE TABLE IF NOT EXISTS address
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    city TEXT,
    street TEXT,
    house TEXT,
    flat TEXT,
    is_current BOOLEAN
    );


CREATE TABLE IF NOT EXISTS category_reference
(
    id UUID NOT NULL PRIMARY KEY,
    category_id UUID,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
    );


CREATE TABLE IF NOT EXISTS shopping_cart_item
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    product_id UUID,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    quantity INT,
    CHECK (quantity >= 0)
    );

CREATE TABLE IF NOT EXISTS favorite
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    product_id UUID,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE
    );

GRANT ALL PRIVILEGES ON DATABASE zuzu to potatiki;

insert into product (id, name_product, price, imgsrc, description, rating)
values ('550e8400-e29b-41d4-a716-446655440000', 'Apple MacBook Air 13 2020', 89999, 'macbook.png', '13-inch lightweight laptop', 4.5);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('be2c8b1b-8d27-4142-a31a-ac6676cf678a', 'Apple MacBook Pro 15 2020', 189999, 'macbook.png', '15-inch professional laptop', 4.85);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c1', 'Apple MacBook Pro 16 2020', 219999, 'macbook.png', '16-inch high-performance laptop', 4.95);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('0d1261e6-3d6f-4eb2-8acd-38fbb8611c5d', 'Apple MacBook Pro 14 2020', 149999, 'macbook.png', '14-inch professional laptop', 4.75);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('3fdc3e65-589d-4aea-be26-5d011dbf7dbd', 'Apple MacBook Pro 13 2020', 99999, 'macbook.png', '13-inch professional laptop', 4.65);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c2', 'Apple MacBook Air 15 2020', 137990, 'macbook.png', '15-inch high-performance laptop', 4.95);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c3', 'Apple MacBook Air 13 2022', 118990, 'macbook.png', '13-inch high-performance laptop', 5.00);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c4', 'Apple MacBook Air 13 2021', 120990, 'macbook.png', '13-inch high-performance laptop', 4.75);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c5', 'Apple MacBook Air 15 2022', 108990, 'macbook.png', '15-inch professional laptop', 5.00);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c6', 'Apple MacBook Air 13 2023', 98990, 'macbook.png', '13-inch lightweight laptop', 4.65);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('550e8400-e29b-41d4-a716-446655440100', 'Apple MacBook Air 13', 89999, 'macbook.png', '13-inch lightweight laptop', 4.5);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('be2c8b1b-8d27-4142-a31a-ac6676cf648a', 'Apple MacBook Pro 15', 189999, 'macbook.png', '15-inch professional laptop', 4.85);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace3c1', 'Apple MacBook Pro 16', 219999, 'macbook.png', '16-inch high-performance laptop', 4.95);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('0d1261e6-3d6f-4eb2-8acd-38fbb8611c7d', 'Apple MacBook Pro 14', 149999, 'macbook.png', '14-inch professional laptop', 4.75);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('3fdc3e65-589d-4aea-be26-5d011dbf4dbd', 'Apple MacBook Pro 13', 99999, 'macbook.png', '13-inch professional laptop', 4.65);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace1c2', 'Apple MacBook Air 15', 137990, 'macbookair.png', '15-inch high-performance laptop', 4.95);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace8c3', 'Apple MacBook Air 14', 118990, 'macbookair.png', '13-inch high-performance laptop', 5.00);

insert into product (id, name_product, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c9', 'Apple MacBook Air 19', 299999, 'macbookair.png', '13-inch high-performance laptop', 5.00);
