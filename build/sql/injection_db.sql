DROP TABLE IF EXISTS profiles;

DROP TABLE IF EXISTS products;

CREATE TABLE IF NOT EXISTS profiles
(
    id uuid NOT NULL,
    login text NOT NULL,
    description text NOT NULL,
    imgsrc text NOT NULL,
    passwordhash text NOT NULL,
    CONSTRAINT "ProfileId_pkey" PRIMARY KEY (id),
    CONSTRAINT "ProfileLogin_unique" UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS products
(
    id uuid NOT NULL,
    nameProduct text NOT NULL,
    description text NOT NULL,
    price int NOT NULL,
    imgsrc text NOT NULL,
    rating NUMERIC(3, 2) NOT NULL,
    CONSTRAINT "ProductId_pkey" PRIMARY KEY (id),
    CONSTRAINT "ProductName_unique" UNIQUE (nameProduct)
);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('550e8400-e29b-41d4-a716-446655440000', 'Apple MacBook Air 13 2020', 89999, 'macbook.png', '13-inch lightweight laptop', 4.5);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('be2c8b1b-8d27-4142-a31a-ac6676cf678a', 'Apple MacBook Pro 15 2020', 189999, 'macbook.png', '15-inch professional laptop', 4.85);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c1', 'Apple MacBook Pro 16 2020', 219999, 'macbook.png', '16-inch high-performance laptop', 4.95);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('0d1261e6-3d6f-4eb2-8acd-38fbb8611c5d', 'Apple MacBook Pro 14 2020', 149999, 'macbook.png', '14-inch professional laptop', 4.75);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('3fdc3e65-589d-4aea-be26-5d011dbf7dbd', 'Apple MacBook Pro 13 2020', 99999, 'macbook.png', '13-inch professional laptop', 4.65);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c2', 'Apple MacBook Air 15 2020', 137990, 'macbook.png', '15-inch high-performance laptop', 4.95);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c3', 'Apple MacBook Air 13 2022', 118990, 'macbook.png', '13-inch high-performance laptop', 5.00);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c4', 'Apple MacBook Air 13 2021', 120990, 'macbook.png', '13-inch high-performance laptop', 4.75);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c5', 'Apple MacBook Air 15 2022', 108990, 'macbook.png', '15-inch professional laptop', 5.00);

insert into products (id, nameProduct, price, imgsrc, description, rating)
values ('007749b5-7e07-4be8-8c91-8db273ace4c6', 'Apple MacBook Air 13 2023', 98990, 'macbook.png', '13-inch lightweight laptop', 4.65);

GRANT ALL PRIVILEGES ON DATABASE zuzu to potatiki;

