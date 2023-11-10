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


CREATE TABLE IF NOT EXISTS profile
(
    id uuid NOT NULL PRIMARY KEY,
    login text NOT NULL UNIQUE,
    description text,
    imgsrc text NOT NULL DEFAULT 'default.png',
    phone text NOT NULL,
    passwordhash bytea NOT NULL
);


INSERT INTO profile (id, login, description, imgsrc, phone, passwordhash)
VALUES
    ('c6e4e63c-8b64-4b98-aebd-76b1ff1c0e9a', 'scremyda', 'Описание пользователя 1', 'user1.png', '+79164424126', E'\\x53b483ac6ff31100c8af51a7815ddd0f4669f9bbff49f7f5e7bdfcf9cc58eea50b5fbaff99343a1c'),
    ('9f85360d-7c1b-4c44-bc13-d73a3e5d4ac3', 'scremyda)', 'Описание пользователя 2', 'user2.png', '79164424126', E'\\x14dc16379c5f2455511b562922bc43cd283470a20287d23c0cef5e08f125281dece36cdc4bb59cbc'),
    ('a7e06ef1-76b5-4e85-a3b8-832745e6d416', 'user3', 'Описание пользователя 3', 'user3.png', '79164424126', E'\\xFEDCBA9876543210'),
    ('4d26e8e7-af08-42d1-8160-8d0d8e7d24b6', 'user4', 'Описание пользователя 4', 'user4.png', '79164424126', E'\\x1234567890ABCDEF'),
    ('f34b43b6-2e4a-4aa3-babf-6e6217c21bf9', 'user5', 'Описание пользователя 5', 'user5.png', '79164424126', E'\\xEFCDAB8967452301'),
    ('39d8c3f9-2f6e-4a3d-8a9b-2b6a8f7e63ab', 'user6', 'Описание пользователя 6', 'user6.png', '79164424126', E'\\xABCDEF0123456789'),
    ('d3a4c7c0-7a6b-4e4a-bc6b-4e4d6a8d7a3c', 'user7', 'Описание пользователя 7', 'user7.png', '79164424126', E'\\xFEDCBA9876543210'),
    ('7e6b3a7d-2e3b-4c0b-8c7c-4e7b3c8e0d3a', 'user8', 'Описание пользователя 8', 'user8.png', '79164424126', E'\\x1234567890ABCDEF'),
    ('1a0e0d0f-0e0c-4c44-8b08-4d6e7b7d6e3e', 'user9', 'Описание пользователя 9', 'user9.png', '79164424126', E'\\xEFCDAB8967452301'),
    ('1a0e0d0f-0e0c-4c44-8167-832745e6d416', 'user10', 'Описание пользователя 10', 'user10.png', '79164424126', E'\\xEFCDAB8967452301');

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
    (7, 'Стиральные машины', 5);
    (8, 'Пылесосы', 5),

--     (9, 'Игровые консоли', 1),
--     (10, 'PlayStation', 9),
--     (11, 'Xbox', 9),
--     (12, 'Nintendo', 9);

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
CREATE TABLE IF NOT EXISTS cart
(
    id UUID NOT NULL PRIMARY KEY,
    profile_id UUID NOT NULL,
    is_current BOOLEAN,
    FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE
    );

INSERT INTO cart (id, profile_id, is_current)
VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'c6e4e63c-8b64-4b98-aebd-76b1ff1c0e9a', true),
    ('98d460d4-3f6e-46f2-a9c7-5e36924a3e0c', '9f85360d-7c1b-4c44-bc13-d73a3e5d4ac3', true),
    ('4d26e8e7-af08-42d1-8160-8d0d8e7d24b6', 'a7e06ef1-76b5-4e85-a3b8-832745e6d416', true),
    ('f34b43b6-2e4a-4aa3-babf-6e6217c21bf9', '4d26e8e7-af08-42d1-8160-8d0d8e7d24b6', true),
    ('39d8c3f9-2f6e-4a3d-8a9b-2b6a8f7e63ab', 'f34b43b6-2e4a-4aa3-babf-6e6217c21bf9', true),
    ('a7e06ef1-76b5-4e85-a3b8-832745e6d416', '39d8c3f9-2f6e-4a3d-8a9b-2b6a8f7e63ab', true),
    ('c6e4e63c-8b64-4b98-aebd-76b1ff1c0e9a', 'd3a4c7c0-7a6b-4e4a-bc6b-4e4d6a8d7a3c', true),
    ('9f85360d-7c1b-4c44-bc13-d73a3e5d4ac3', '7e6b3a7d-2e3b-4c0b-8c7c-4e7b3c8e0d3a', true),
    ('d3a4c7c0-7a6b-4e4a-bc6b-4e4d6a8d7a3c', '1a0e0d0f-0e0c-4c44-8b08-4d6e7b7d6e3e', true),
    ('7e6b3a7d-2e3b-4c0b-8c7c-4e7b3c8e0d3a', '1a0e0d0f-0e0c-4c44-8167-832745e6d416', true);
------------------------------------------------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS product
(
    id uuid NOT NULL PRIMARY KEY,
    name text NOT NULL UNIQUE ,
    description text NOT NULL,
    price INT NOT NULL,
    imgsrc text NOT NULL,
    rating NUMERIC(3, 2) NOT NULL,
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE RESTRICT,
    CHECK (rating >= 0),
    CHECK (price > 0)
    );

INSERT INTO product (id, name, price, imgsrc, description, rating, category_id)
VALUES
    ('550e8400-e29b-41d4-a716-446655440000', 'Apple MacBook Air 15 (M2, 8C CPU/10C GPU, 2023), 8 ГБ, 512 ГБ SSD, «полуночный черный»',
     189990, 'MacbookAir15.png',
     'Apple MacBook Air 15 — ноутбук, объединивший в себе инновационные технологии, большой четкий дисплей,' ||
     ' высокую производительность, небольшие габариты и отличную эргономику. Это первая модель из линейки Air,' ||
     ' которая получила экран диагональю 15 дюймов.', 4.7, 4),

    ('be2c8b1b-8d27-4142-a31a-ac6676cf678a', 'Apple MacBook Pro 14" (M3 Max 14C CPU, 30C GPU, 2023) 36 ГБ, 1 ТБ SSD, серебристый',
     229999, 'MacbookPro14M3Grey.jpg',
     'Apple MacBook Pro — мощный и легкий ноутбук, предназначенный для профессиональной деятельности. ' ||
     'В устройстве объединили бескомпромиссную мощь, четкий и яркий дисплей, продуманную эргономику и инновационные ' ||
     'технологии. Корпус выполнен в строгом дизайне.', 4.85, 4),

    ('007749b5-7e07-4be8-8c91-8db273ace4c1', 'Apple MacBook Pro 16" (M2 Max 12C CPU, 30C GPU, 2023) 64 ГБ, 1 ТБ SSD, «серый космос»',
     249999, 'MacbookPro16M2MaxGrey.jpeg', 'Apple MacBook Pro — мощный и легкий ноутбук, предназначенный для профессиональной деятельности. ' ||
                                           'В устройстве объединили бескомпромиссную мощь, четкий и яркий дисплей, продуманную эргономику и инновационные ' ||
                                           'технологии. Корпус выполнен в строгом дизайне.', 4.95, 4),

    ('0d1261e6-3d6f-4eb2-8acd-38fbb8611c5d', '14 Ультрабук HUAWEI MateBook D 14 NbD-WDI9 серый (53013FCE)',
     79999, 'HUAWEIMateBookD14.jpg',
     'Ноутбук HUAWEI MateBook D 14 i5-1155G7 BoDE-WFH9 SpaceGray — устройство в корпусе из алюминиевого сплава,' ||
     ' которое работает на базе ОС Windows 11. Для управления предусмотрена клавиатура с английской и русской раскладкой. ' ||
     'Веб-камера разрешением на 1 Мп дает возможность принимать участие в видеоконференциях. ', 4.75, 4),

    ('3fdc3e65-589d-4aea-be26-5d011dbf7dbd', 'Asus TUF Gaming A15 FX506HE-HN012 (90NR0704-M02050)',
     99999, 'AsusTUFGamingA15.jpg', 'Игровой ноутбук Asus TUF Gaming A15 FX506HE-HN012 i5 11400H/8Gb/512Gb SSD/15.6" ' ||
                                    'FHD IPS 144Ghz/RTX 3050Ti 4Gb/DOS/Graphite Black. Для управления предусмотрена клавиатура с английской и русской раскладкой.', 4.65, 4),

    ('007749b5-7e07-4be8-8c91-8db273ace4c2', 'MSI Modern 14 C12M-249XBY-BB31215U8GXXDXX (Modern 14 C12M-249XBY-BB31215U8GXXDXX)',
     52990, 'MSIModern14.jpg',
     'Ноутбук MSI Modern 14 C12M-249XBY-BB31215U8GXXDXX 14" FHD IPS i3-1215U/8Gb/SSD 256Gb/Iris ' ||
     'Xe/DOS/Classic Black (Modern 14 C12M-249XBY-BB31215U8GXXDXX) Для управления предусмотрена клавиатура с английской и русской раскладкой.', 4.95, 4),

    ('007749b5-7e07-4be8-8c91-8db273ace4c3', 'Ноутбук Lenovo IdeaPad 3 15ITL6 (82H801B6RK)', 118990,
     'lenovo3.jpg',
     'Ноутбук Lenovo IdeaPad 3 15ITL6 i5 1135G7/12Gb/1Tb HDD + 256Gb SSD/noDVD/GeForce MX350 2Gb/15.6" ' ||
     'IPS FHD/noOS/Grey (82H801B6RK) Для управления предусмотрена клавиатура с английской и русской раскладкой.', 5.00, 4),

    ('007749b5-7e07-4be8-8c91-8db273ace4c4', 'Apple iPad Pro (2022) 11" Wi-Fi 256 ГБ, серебристый', 120990,
     'Ipadpro2022.jpg',
     'Работайте, учитесь или отдыхайте с Apple iPad Pro 2022 Wi-Fi Cell.' ||
     ' Планшетный компьютер имеет 11-дюймовый дисплей Multi‑Touch со светодиодной подсветкой' ||
     ' и технологией IPS и объем памяти 128 ГБ.', 4.75, 3),

    ('007749b5-7e07-4be8-8c91-8db273ace4c5', 'Apple iPad mini (2021) Wi-Fi + Cellular 64 ГБ, фиолетовый', 108990,
     'IPAD2022.jpg',  'Работайте, учитесь или отдыхайте с Apple iPad Mini 2021 Wi-Fi Cell.' ||
                      ' Планшетный компьютер имеет 11-дюймовый дисплей Multi‑Touch со светодиодной подсветкой' ||
                      ' и технологией IPS и объем памяти 64 ГБ.', 4.85, 3),

    ('007749b5-7e07-4be8-8c91-8db273ace4c6', 'Apple MacBook Air 13 2023', 98990, 'macbook.png', '13-inch lightweight laptop', 4.65, 4),

    ('550e8400-e29b-41d4-a716-446655440100', 'Apple MacBook Air 13', 89999, 'macbook.png', '13-inch lightweight laptop', 4.5, 4),
    ('be2c8b1b-8d27-4142-a31a-ac6676cf648a', 'Apple MacBook Pro 15', 189999, 'macbook.png', '15-inch professional laptop', 4.85, 4),
    ('007749b5-7e07-4be8-8c91-8db273ace3c1', 'Apple MacBook Pro 16', 219999, 'macbook.png', '16-inch high-performance laptop', 4.95, 4),
    ('0d1261e6-3d6f-4eb2-8acd-38fbb8611c7d', 'Apple MacBook Pro 14', 149999, 'macbook.png', '14-inch professional laptop', 4.75, 4),
    ('3fdc3e65-589d-4aea-be26-5d011dbf4dbd', 'Apple MacBook Pro 13', 99999, 'macbook.png', '13-inch professional laptop', 4.65, 4),
    ('007749b5-7e07-4be8-8c91-8db273ace1c2', 'Apple MacBook Air 15', 137990, 'macbook.png', '15-inch high-performance laptop', 4.95, 4),
    ('007749b5-7e07-4be8-8c91-8db273ace8c3', 'Apple MacBook Air 14', 118990, 'macbook.png', '13-inch high-performance laptop', 5.00, 4),
    ('007749b5-7e07-4be8-8c91-8db273ace4c9', 'Apple MacBook Air 19', 299999, 'macbook.png', '13-inch high-performance laptop', 5.00, 4),

--------------------Холодильники-6
    ('007749b5-434c-4be8-8c91-8db273ace100', 'Холодильник Samsung RL4362RBASL/WT', 125999, 'holodos1.jpg', 'Холодильник Samsung RL4362RBASL/WT позволит уберечь ваши продукты питания от порчи. В холодильнике имеются две камеры: холодильная и морозильная, которая располагается в нижней части. Камеры в INDESIT DS 4180 S B размораживаются двумя способами: холодильная - капельным, а морозильная - ручным. В случае, если прекратится подача электроэнергии, устройство еще в течение 18 часов будет способно сохранять холод. В компрессоре холодильника используется хладагент R600a (изобутан). Возможность перевешивания дверей позволит вам выбрать, в какую сторону будет открываться холодильник.', 4.95, 6),
    ('007749b5-434c-4be8-8c91-8db273ace101', 'Холодильник Toshiba GR-RB449WE-PMJ(06)', 64999, 'holodos2.jpg', 'Холодильник Toshiba GR-RB449WE-PMJ(06) позволит уберечь ваши продукты питания от порчи. В холодильнике имеются две камеры: холодильная и морозильная, которая располагается в нижней части. Камеры в INDESIT DS 4180 S B размораживаются двумя способами: холодильная - капельным, а морозильная - ручным. В случае, если прекратится подача электроэнергии, устройство еще в течение 18 часов будет способно сохранять холод. В компрессоре холодильника используется хладагент R600a (изобутан). Возможность перевешивания дверей позволит вам выбрать, в какую сторону будет открываться холодильник.', 4.91, 6),
--------------------Стиральные машины-7
    ('007749b5-434c-4be8-8c91-8db273ace102', 'Стиральная машина Weissgauff WM 4606', 22370, 'stiralka1.jpg', 'Белая стиральная машина с нагревательным элементом Hi-Tech с защитой от накипи, максимальной загрузкой барабана 6 кг и электронным контролем дисбаланса, который предотвращает шум и вибрацию во время стирки и отжима. Стиральная машина оснащена технологией SteamCure, которая обрабатывает белье паром до и после стирки, облегчает отстирывание пятен и разглаживает ткани после отжима. 15 автоматических программ стирки позволяют ухаживать за одеждой из разных тканей, от хлопка до шерсти, а также джинсовыми, спортивными вещами и пуховиками.', 4.99, 7),
    ('007749b5-434c-4be8-8c91-8db273ace103', 'Cтиральная машина Beko WSPE6H616A', 25898, 'stiralka2.jpg', 'Черная стиральная машина с нагревательным элементом Hi-Tech с защитой от накипи, максимальной загрузкой барабана 6 кг и электронным контролем дисбаланса, который предотвращает шум и вибрацию во время стирки и отжима. Стиральная машина оснащена технологией SteamCure, которая обрабатывает белье паром до и после стирки, облегчает отстирывание пятен и разглаживает ткани после отжима. 15 автоматических программ стирки позволяют ухаживать за одеждой из разных тканей, от хлопка до шерсти, а также джинсовыми, спортивными вещами и пуховиками.', 4.66, 7);


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


INSERT INTO shopping_cart_item (cart_id, product_id, quantity)
VALUES
    ( 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 1),
    ( 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 3),

    ('98d460d4-3f6e-46f2-a9c7-5e36924a3e0c', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 1),
    ( '98d460d4-3f6e-46f2-a9c7-5e36924a3e0c', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 2),

    ( '4d26e8e7-af08-42d1-8160-8d0d8e7d24b6', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 4),
    ( '4d26e8e7-af08-42d1-8160-8d0d8e7d24b6', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 1),

    ( 'f34b43b6-2e4a-4aa3-babf-6e6217c21bf9', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 1),
    ( 'f34b43b6-2e4a-4aa3-babf-6e6217c21bf9', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 1),

    ( '39d8c3f9-2f6e-4a3d-8a9b-2b6a8f7e63ab', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 2),
    ( '39d8c3f9-2f6e-4a3d-8a9b-2b6a8f7e63ab', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 2),

    ( 'a7e06ef1-76b5-4e85-a3b8-832745e6d416', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 1),
    ( 'a7e06ef1-76b5-4e85-a3b8-832745e6d416', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 2),

    ( 'c6e4e63c-8b64-4b98-aebd-76b1ff1c0e9a', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 3),
    ('c6e4e63c-8b64-4b98-aebd-76b1ff1c0e9a', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 4),

    ( '9f85360d-7c1b-4c44-bc13-d73a3e5d4ac3', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 1),
    ( '9f85360d-7c1b-4c44-bc13-d73a3e5d4ac3', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 2),

    ( 'd3a4c7c0-7a6b-4e4a-bc6b-4e4d6a8d7a3c', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 3),
    ( 'd3a4c7c0-7a6b-4e4a-bc6b-4e4d6a8d7a3c', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 3),

    ( '7e6b3a7d-2e3b-4c0b-8c7c-4e7b3c8e0d3a', 'be2c8b1b-8d27-4142-a31a-ac6676cf678a', 9),
    ( '7e6b3a7d-2e3b-4c0b-8c7c-4e7b3c8e0d3a', 'be2c8b1b-8d27-4142-a31a-ac6676cf648a', 1);
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


-- INSERT INTO order_info (id, profile_id, status_id, delivery_at)
-- VALUES
--     ('1a0e0d0f-0e0c-4c44-8167-832745e6d416', 'c6e4e63c-8b64-4b98-aebd-76b1ff1c0e9a', 3,  '2023-10-31 12:00:00'),
--     ('1a0e0d0f-0e0c-4c44-8b08-4d6e7b7d6e3e', '9f85360d-7c1b-4c44-bc13-d73a3e5d4ac3', 3,  '2023-10-31 14:00:00'),
--     ('7e6b3a7d-2e3b-4c0b-8c7c-4e7b3c8e0d3a', 'a7e06ef1-76b5-4e85-a3b8-832745e6d416', 3,  '2023-10-31 16:00:00'),
--     ('d3a4c7c0-7a6b-4e4a-bc6b-4e4d6a8d7a3c', '4d26e8e7-af08-42d1-8160-8d0d8e7d24b6', 3,  '2023-10-31 18:00:00'),
--     ('39d8c3f9-2f6e-4a3d-8a9b-2b6a8f7e63ab', 'f34b43b6-2e4a-4aa3-babf-6e6217c21bf9', 3,  '2023-10-31 20:00:00'),
--     ('f34b43b6-2e4a-4aa3-babf-6e6217c21bf9', '39d8c3f9-2f6e-4a3d-8a9b-2b6a8f7e63ab', 3,  '2023-10-31 22:00:00'),
--     ('c6e4e63c-8b64-4b98-aebd-76b1ff1c0e9a', 'd3a4c7c0-7a6b-4e4a-bc6b-4e4d6a8d7a3c', 3, '2023-11-01 12:00:00'),
--     ('9f85360d-7c1b-4c44-bc13-d73a3e5d4ac3', '7e6b3a7d-2e3b-4c0b-8c7c-4e7b3c8e0d3a', 3,  '2023-11-01 14:00:00'),
--     ('a7e06ef1-76b5-4e85-a3b8-832745e6d416', '1a0e0d0f-0e0c-4c44-8b08-4d6e7b7d6e3e', 3,  '2023-11-01 16:00:00'),
--     ('4d26e8e7-af08-42d1-8160-8d0d8e7d24b6', '1a0e0d0f-0e0c-4c44-8167-832745e6d416', 3, '2023-11-01 18:00:00');

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

