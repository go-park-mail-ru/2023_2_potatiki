DROP TABLE IF EXISTS profile;

DROP TABLE IF EXISTS product;

DROP TABLE IF EXISTS order_info;

DROP TABLE IF EXISTS order_item;

DROP TABLE IF EXISTS category;

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
    CONSTRAINT "ProfileLogin_unique" UNIQUE (login),
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
    CHECK (rating >= 0)
        CHECK (price > 0)
);

CREATE TABLE IF NOT EXISTS promocode
(
    id UUID NOT NULL PRIMARY KEY,
    name TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS order_info
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID,
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