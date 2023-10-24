CREATE TABLE IF NOT EXISTS profile
(
    id uuid NOT NULL PRIMARY KEY,
    login text NOT NULL UNIQUE,
    description text,
    imgsrc text NOT NULL DEFAULT 'default.png',
    passwordhash text NOT NULL
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
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS order_info
(
    id UUID NOT NULL PRIMARY KEY,
    delivery_date TIMESTAMP,
    creation_date TIMESTAMP NOT NULL,
    profile_id UUID NOT NULL,
    status TEXT NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    promocode_id UUID,
    CONSTRAINT "ProfilePromocode_unique" UNIQUE (profile_id, promocode_id),
    FOREIGN KEY (promocode_id) REFERENCES promocode(id) ON DELETE RESTRICT
);


CREATE TABLE IF NOT EXISTS order_item
(
    id UUID NOT NULL PRIMARY KEY,
    order_id UUID NOT NULL,
    FOREIGN KEY (order_id) REFERENCES order_info(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT "OrderProduct_unique" UNIQUE (order_id, product_id),
    quantity INT NOT NULL,
    CHECK (quantity > 0)
);


CREATE TABLE IF NOT EXISTS address
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    city TEXT NOT NULL,
    street TEXT NOT NULL,
    house TEXT NOT NULL,
    flat TEXT NOT NULL,
    CONSTRAINT "Address_unique" UNIQUE (city, street, house, flat),
    is_current BOOLEAN
);


CREATE TABLE IF NOT EXISTS category_reference
(
    id UUID NOT NULL PRIMARY KEY,
    category_id UUID,
    FOREIGN KEY (id) REFERENCES category(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS shopping_cart_item
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT "ProfileProduct_unique" UNIQUE (profile_id, product_id),
    quantity INT NOT NULL,
    CHECK (quantity > 0)
);

CREATE TABLE IF NOT EXISTS favorite
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT "ProfileProduct_unique" UNIQUE (profile_id, product_id)
);
