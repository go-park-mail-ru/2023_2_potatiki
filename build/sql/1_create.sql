DROP TABLE IF EXISTS profile CASCADE;

DROP TABLE IF EXISTS product CASCADE;

DROP TABLE IF EXISTS order_info CASCADE;

DROP TABLE IF EXISTS order_item;

DROP TABLE IF EXISTS category CASCADE;

DROP TABLE IF EXISTS shopping_cart_item;

DROP TABLE IF EXISTS cart;

DROP TABLE IF EXISTS promocode;

DROP TABLE IF EXISTS address;

DROP TABLE IF EXISTS favorite;

DROP TABLE IF EXISTS status;

DROP TABLE IF EXISTS comment;

DROP TABLE IF EXISTS survey CASCADE;

DROP TABLE IF EXISTS question_type CASCADE;

DROP TABLE IF EXISTS question CASCADE;

DROP TABLE IF EXISTS answer;

DROP TABLE IF EXISTS results;

CREATE TABLE IF NOT EXISTS survey
(
    id uuid NOT NULL PRIMARY KEY,
    name text UNIQUE
);

INSERT INTO survey (id, name)
VALUES
    ('1e461708-6b04-45b9-a4fa-77c32c14d982', 'Опрос про товары'),
    ('1e461708-6b04-45b9-a4fa-77c32c14d487', 'Опрос про вид'),
    ('1e461708-6b04-45b9-a4fa-77c32c14d318', 'Опрос про что-то еще');

CREATE TABLE IF NOT EXISTS question_type
(
    id serial NOT NULL PRIMARY KEY,
    type text NOT NULL
);

INSERT INTO question_type (id, type)
VALUES
    (1, 'CSAT'),
    (2, 'NPS'),
    (3, 'CSI');

CREATE TABLE IF NOT EXISTS question
(
    id uuid NOT NULL PRIMARY KEY,
    type int NOT NULL,
    FOREIGN KEY (type) REFERENCES question_type(id) ON DELETE RESTRICT,
    id_survey uuid,
    FOREIGN KEY (id_survey) REFERENCES survey(id) ON DELETE RESTRICT,
    name text UNIQUE
);

INSERT INTO question (id, type, id_survey, name)
VALUES
    ('1e461708-6b04-45b9-a4fa-77c32c14d382', 1, '1e461708-6b04-45b9-a4fa-77c32c14d982', 'Крутой сайт?'),
    ('1e461708-6b04-45b9-a4fa-77c32c14d387', 1,'1e461708-6b04-45b9-a4fa-77c32c14d982', 'Круто дизайн?'),
    ('1e461708-6b04-45b9-a4fa-77c32c14d388', 2,'1e461708-6b04-45b9-a4fa-77c32c14d982', 'Ваше мнение?');

CREATE TABLE IF NOT EXISTS results
(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid,
    survey_id uuid,
    FOREIGN KEY (survey_id) REFERENCES survey(id) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS answer
(
    id uuid NOT NULL PRIMARY KEY,
    question uuid,
    FOREIGN KEY (question) REFERENCES question(id) ON DELETE RESTRICT,
    result_id uuid,
    FOREIGN KEY (result_id) REFERENCES results(id) ON DELETE RESTRICT,
    answer int
);

--- setup ru_RU dictionary ---
CREATE TEXT SEARCH DICTIONARY russian_ispell (
    TEMPLATE = ispell,
    DictFile = russian,
    AffFile = russian,
    Stopwords = russian
);
CREATE TEXT SEARCH CONFIGURATION ru (COPY=russian);
ALTER TEXT SEARCH CONFIGURATION ru
    ALTER MAPPING FOR hword, hword_part, word
    WITH russian_ispell, russian_stem;
--- setup ru_RU dictionary ---

------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS profile
(
    id uuid NOT NULL PRIMARY KEY,
    login text NOT NULL UNIQUE,
    description text,
    imgsrc text NOT NULL DEFAULT 'default.png',
    phone text NOT NULL,
    passwordhash bytea NOT NULL
);

------------------------------------------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS category
(
    id INT PRIMARY KEY,
    name TEXT UNIQUE,
    parent INT DEFAULT NULL REFERENCES category (id)
    );

INSERT INTO category
VALUES
    (1, 'Все товары', NULL),
    (2, 'Ноутбуки и планшеты', 1),
    (3, 'Планшеты', 2),
    (4, 'Ноутбуки', 2),
    (5, 'Бытовая техника', 1),
    (6, 'Холодильники', 5),
    (7, 'Стиральные машины', 5),
    (8, 'Пылесосы', 5),
    (9, 'Мебель', 1),
    (91, 'Стулья', 9),
    (92, 'Рабочие столы', 9),
    (93, 'Диваны', 9),
    (94, 'Кресла', 9),
    (10, 'Канцелярия', 1),
    (101, 'Тетради', 10),
    (102, 'Письменные принадлежности', 10),
    (103, 'Пеналы', 10),
    (104, 'Клей', 10),
    (11, 'Товары для геймеров', 1),
    (111, 'Nintendo', 11),
    (112, 'Xbox', 11),
    (113, 'PlayStation', 11);

------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS status
(
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE
);

INSERT INTO status (name)
VALUES
    ('cart'),
    ('created'),
    ('processed'),
    ('delivery'),
    ('delivered'),
    ('received'),
    ('returned');

------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS product
(
    id uuid NOT NULL PRIMARY KEY,
    name text NOT NULL,
    description text NOT NULL,
    price INT NOT NULL,
    imgsrc text NOT NULL,
    category_id INT,
    rating NUMERIC(3, 2) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE RESTRICT,
    creation_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CHECK (rating >= 0),
    CHECK (price > 0)
    );
------------------------------------------------------------------------------------------------------------------------
CREATE TABLE comment
(
    id uuid NOT NULL PRIMARY KEY,
    productID uuid REFERENCES product (id) ON DELETE CASCADE,
    userID uuid REFERENCES profile (id) ON DELETE CASCADE,
    pros text NOT NULL,
    cons text NOT NULL,
    comment text NOT NULL,
    creation_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    rating NUMERIC(3, 2) NOT NULL,
    CHECK (rating >= 0)
);
--- DICTIONARY ---------------------------------------------------------------------------------------------------------

CREATE EXTENSION pg_trgm;

--- DICTIONARY ---------------------------------------------------------------------------------------------------------


------------------------------------------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS promocode
(
    id SERIAL PRIMARY KEY,
    discount INT NOT NULL,
    name TEXT NOT NULL UNIQUE
);

INSERT INTO promocode (discount, name)
VALUES
    (10, 'PROMO10'),
    (15, 'SALE15'),
    (20, 'DISCOUNT20'),
    (25, 'SAVE25'),
    (30, '30OFF');
------------------------------------------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS cart
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID NOT NULL,
    is_current BOOLEAN,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE
    );
------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS shopping_cart_item
(
    id SERIAL PRIMARY KEY,
    cart_id UUID NOT NULL,
    FOREIGN KEY (cart_id) REFERENCES cart(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT uq_shopping_cart_item_cart_id_product_id UNIQUE (cart_id, product_id),
    quantity INT NOT NULL,
    CHECK (quantity > 0)
    );
------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS address
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    city TEXT NOT NULL,
    street TEXT NOT NULL,
    house TEXT NOT NULL,
    flat TEXT NOT NULL,
    is_current BOOLEAN,
    is_deleted BOOLEAN DEFAULT FALSE
    );

------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS order_info
(
    id UUID NOT NULL PRIMARY KEY,
    delivery_at TIMESTAMPTZ,
    creation_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    profile_id UUID NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    status_id INT NOT NULL,
    FOREIGN KEY (status_id) REFERENCES status(id) ON DELETE RESTRICT,
    promocode_id INT,
    CONSTRAINT uq_order_info_profile_id_promocode_id UNIQUE (profile_id, promocode_id),
    FOREIGN KEY (promocode_id) REFERENCES promocode(id) ON DELETE RESTRICT,
    address_id UUID NOT NULL,
    FOREIGN KEY (address_id) REFERENCES address(id) ON DELETE RESTRICT
    );

------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS order_item
(
    id UUID NOT NULL PRIMARY KEY,
    order_id UUID NOT NULL,
    FOREIGN KEY (order_id) REFERENCES order_info(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT uq_order_item_order_id_product_id UNIQUE (order_id, product_id),
    price INT NOT NULL,
    quantity INT NOT NULL,
    CHECK (quantity > 0),
    CHECK (price > 0)
    );

------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS favorite
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT uq_favorite_profile_id_product_id UNIQUE (profile_id, product_id)
    );

------------------------------------------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION update_is_current()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW.is_current = true THEN
UPDATE address
SET is_current = false
WHERE profile_id = NEW.profile_id AND id <> NEW.id;
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;



CREATE TRIGGER set_is_current_on_insert
    BEFORE INSERT ON address
    FOR EACH ROW
    EXECUTE FUNCTION update_is_current();


CREATE TRIGGER set_is_current_on_update
    BEFORE UPDATE ON address
    FOR EACH ROW
    EXECUTE FUNCTION update_is_current();
