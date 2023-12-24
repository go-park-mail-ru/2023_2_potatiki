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

DROP TABLE IF EXISTS messages;

DROP TABLE IF EXISTS activities;

------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS messages (
    user_id uuid NOT NULL,
    created TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    message_info TEXT
);

------------------------------------------------------------------------------------------------------------------------

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

    (2, 'Электроника', 1),
    (21, 'Планшеты', 2),
    (22, 'Ноутбуки', 2),
    (23, 'Мониторы', 2),
    (23, 'Наушники', 2),

    (3, 'Бытовая техника', 1),
    (31, 'Холодильники', 3),
    (32, 'Стиральные машины', 3),
    (33, 'Пылесосы', 3),

    (4, 'Музыкальные инструменты', 1),
    (41, 'Губные гармошки', 4),
    (42, 'Гитары', 4),
    (43, 'Барабаны', 4),
    (44, 'Клавишные', 4),
    (45, 'Смычковые музыкальные инструменты', 4),
    (46, 'Духовые музыкальные инструменты', 4),
    (47, 'Виниловые пластинки', 4),

    (5, 'Спорт и активный отдых', 1),
    (51, 'Велосипеды', 5),
    (52, 'Горные лыжи', 5),
    (53, 'Сноуборды', 5),
    (54, 'Самокаты', 5),
    (55, 'Веревки альпинистские', 5),
    (56, 'Дартс', 5),

    (6, 'Красота и уход', 1),
    (61, 'Уход за лицом', 6),
    (62, 'Средства по уходу за волосами', 6),
    (63, 'Косметика для макияжа лица', 6),
    (64, 'Макияж глаз', 6),

    (7, 'Ювелирные изделия', 1),
    (71, 'Кольца', 7),
    (72, 'Серьги', 7),
    (73, 'Браслеты', 7),
    (74, 'Цепочки', 7),
    (75, 'Колье', 7),

    (8, 'Новогодние товары', 1),
    (81, 'Елки искусственные', 8),
    (82, 'Живые елки', 8),
    (83, 'Аксессуары для елок', 8),
    (84, 'Елочные украшения', 8),
    (85, 'Новогодние гирлянды светодиодные', 8),

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
    ('В обработке'),
    ('Передан в службу доставки'),
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
    count_comments INT DEFAULT 0,
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

CREATE EXTENSION IF NOT EXISTS pg_trgm;

--- DICTIONARY ---------------------------------------------------------------------------------------------------------


------------------------------------------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS promocode
(
    id SERIAL PRIMARY KEY,
    discount INT NOT NULL,
    name TEXT NOT NULL UNIQUE,
    leftover INT NOT NULL,
    deadline TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

INSERT INTO promocode (discount, name, leftover, deadline)
VALUES
    (10, 'PROMO10', 100, '2024-01-01 00:00:00'),
    (15, 'SALE15', 1000, '2024-01-01 00:00:00'),
    (20, 'DISCOUNT20', 100, '2024-01-01 00:00:00'),
    (25, 'SAVE25', 100, '2024-01-01 00:00:00'),
    (10, 'ZUZU10', 100000, '2024-01-01 00:00:00'),
    (30, 'OFF30', 5, '2024-01-01 00:00:00');
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
    delivery_at_date TEXT NOT NULL,
    delivery_at_time TEXT NOT NULL,
    creation_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    profile_id UUID NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE,
    status_id INT NOT NULL,
    FOREIGN KEY (status_id) REFERENCES status(id) ON DELETE RESTRICT,
    promocode_id INT,
    FOREIGN KEY (promocode_id) REFERENCES promocode(id) ON DELETE RESTRICT,
    CONSTRAINT uq_order_info_profile_id_promocode_id UNIQUE (profile_id, promocode_id),
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


-- Создание функции для обновления счетчика комментариев
CREATE OR REPLACE FUNCTION update_comment_count()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE product
    SET count_comments = count_comments + 1
    WHERE id = NEW.productID;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создание триггера, который вызывает функцию при добавлении нового комментария
CREATE TRIGGER comment_trigger
    AFTER INSERT ON comment
    FOR EACH ROW
EXECUTE FUNCTION update_comment_count();

------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS activities
(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    body jsonb
);

CREATE OR REPLACE FUNCTION order_created_trigger()
    RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM messages
    WHERE created < CURRENT_TIMESTAMP - interval '1 day';

    IF NEW.profile_id IS NOT NULL THEN
        INSERT INTO messages (user_id, created, message_info)
        VALUES (NEW.profile_id, CURRENT_TIMESTAMP, 'Заказ создан, загляните в раздел: "Заказы"');
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER order_created
    AFTER INSERT
    ON order_info
    FOR EACH ROW
EXECUTE FUNCTION order_created_trigger();

------------------------------------------------------------------------------------------------------------------------

CREATE OR REPLACE FUNCTION profile_created_trigger()
    RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM messages
    WHERE created < CURRENT_TIMESTAMP - interval '1 day';

    IF NEW.id IS NOT NULL THEN
        INSERT INTO messages (user_id, created, message_info)
        VALUES (NEW.id, CURRENT_TIMESTAMP + interval '1 second', 'Спасибо за регистрацию, мы дарим вам промокод: ZUZU10');
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER profile_created
    AFTER INSERT
    ON profile
    FOR EACH ROW
EXECUTE FUNCTION profile_created_trigger();
