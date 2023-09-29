DROP TABLE IF EXISTS profiles;

DROP TABLE IF EXISTS products;

CREATE TABLE IF NOT EXISTS profiles
(
    id uuid NOT NULL,
    login text,
    description text,
    imgsrc text,
    passwordhash text,
    CONSTRAINT "ProfileId_pkey" PRIMARY KEY (id),
    CONSTRAINT "ProfileLogin_unique" UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS products
(
    id uuid NOT NULL,
    nameProduct text,
    description text,
    price int,
    CONSTRAINT "ProductId_pkey" PRIMARY KEY (id),
    CONSTRAINT "ProductName_unique" UNIQUE (nameProduct)
);

GRANT ALL PRIVILEGES ON DATABASE zuzu to potatiki;

