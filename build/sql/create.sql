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
    (7, 'Стиральные машины', 5),
    (8, 'Пылесосы', 5);

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
    name text NOT NULL,
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
    ('007749b5-434c-4be8-8c91-8db273ace103', 'Cтиральная машина Beko WSPE6H616A', 25898, 'stiralka2.jpg', 'Черная стиральная машина с нагревательным элементом Hi-Tech с защитой от накипи, максимальной загрузкой барабана 6 кг и электронным контролем дисбаланса, который предотвращает шум и вибрацию во время стирки и отжима. Стиральная машина оснащена технологией SteamCure, которая обрабатывает белье паром до и после стирки, облегчает отстирывание пятен и разглаживает ткани после отжима. 15 автоматических программ стирки позволяют ухаживать за одеждой из разных тканей, от хлопка до шерсти, а также джинсовыми, спортивными вещами и пуховиками.', 4.66, 7),
    
    ('8b6989d0-8c8b-46fa-a1a2-79fa8fb9085d', 'Стиральная машина Haier HW60-BP10919B белый', 43120, '7-8b6989d0-8c8b-46fa-a1a2-79fa8fb9085d.jpeg', 'Самый лучший среди товаров на рынке Стиральная машина Haier HW60-BP10919B белый', 4.94, 7),
    ('5fbe7cd4-5327-47c0-b3fb-51ef85abf170', 'Стиральная машина Beko RSPE78612W', 37390, '7-5fbe7cd4-5327-47c0-b3fb-51ef85abf170.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko RSPE78612W', 4.20, 7),
    ('29901604-f4db-4ab6-8d7d-1d4d7bbdc292', 'Стиральная машина LG F2J3NS2W белый', 62850, '7-29901604-f4db-4ab6-8d7d-1d4d7bbdc292.png', 'Самый лучший среди товаров на рынке Стиральная машина LG F2J3NS2W белый', 4.67, 7),
    ('7ac82de4-5a5d-4fa7-809d-28a1fa9b5301', 'Стиральная машина Haier HW60-BP10919B белый', 45629, '7-7ac82de4-5a5d-4fa7-809d-28a1fa9b5301.jpeg', 'Самый лучший среди товаров на рынке Стиральная машина Haier HW60-BP10919B белый', 4.94, 7),
    ('11f32448-e66b-477b-a470-8de5ced847b6', 'Стиральная машина Indesit IWSB 5105 (CIS)', 31837, '7-11f32448-e66b-477b-a470-8de5ced847b6.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWSB 5105 (CIS)', 4.0, 7),
    ('8e3bec39-8881-45aa-83ad-490707a832c2', 'Стиральная машина Beko WRE 6512 BWW', 32840, '7-8e3bec39-8881-45aa-83ad-490707a832c2.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WRE 6512 BWW', 4.83, 7),
    ('dd7b0332-a867-48ef-a34e-3b9b317854c3', 'Стиральная машина Beko WSPE6H616W', 41039, '7-dd7b0332-a867-48ef-a34e-3b9b317854c3.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WSPE6H616W', 4.83, 7),
    ('e35ca5de-927c-484a-bd95-c3d08b603d2a', 'Стиральная машина Haier HW60-BP12919B белая', 48120, '7-e35ca5de-927c-484a-bd95-c3d08b603d2a.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Haier HW60-BP12919B белая', 4.67, 7),
    ('1d89544d-3e88-44ac-a57c-2837c86ee1b5', 'Стиральная машина Beko WSPE7H616W', 48849, '7-1d89544d-3e88-44ac-a57c-2837c86ee1b5.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WSPE7H616W', 4.21, 7),
    ('4a1d5e89-e742-4780-b03e-a27038d2195f', 'Стиральная машина LG F2T3HS6W', 59999, '7-4a1d5e89-e742-4780-b03e-a27038d2195f.jpg', 'Самый лучший среди товаров на рынке Стиральная машина LG F2T3HS6W', 4.54, 7),
    ('4075fcff-0279-4d4f-80da-65cce386ec0d', 'Стиральная машина ATLANT СМА-50 У 87', 28540, '7-4075fcff-0279-4d4f-80da-65cce386ec0d.jpg', 'Самый лучший среди товаров на рынке Стиральная машина ATLANT СМА-50 У 87', 4.65, 7),
    ('15297923-8cd8-4e44-990e-39e93d040ee4', 'Стиральная машина Beko RSPE78612W', 31860, '7-15297923-8cd8-4e44-990e-39e93d040ee4.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko RSPE78612W', 4.20, 7),
    ('573a6cd0-50b0-4d55-aff2-52208143b1bd', 'Стиральная машина LG F2J3NS2W белый', 54800, '7-573a6cd0-50b0-4d55-aff2-52208143b1bd.png', 'Самый лучший среди товаров на рынке Стиральная машина LG F2J3NS2W белый', 4.67, 7),
    ('16fc4d7f-355f-4420-8033-13bbe9ea5f1a', 'Стиральная машина Haier HW60-BP10919B белый', 38999, '7-16fc4d7f-355f-4420-8033-13bbe9ea5f1a.jpeg', 'Самый лучший среди товаров на рынке Стиральная машина Haier HW60-BP10919B белый', 4.94, 7),
    ('aa92262e-d135-4558-a7a1-5c0a6c9006ab', 'Стиральная машина Indesit IWSB 5105 (CIS)', 24843, '7-aa92262e-d135-4558-a7a1-5c0a6c9006ab.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWSB 5105 (CIS)', 4.0, 7),
    ('d0f47f77-9053-475d-bcfd-c3264031ff34', 'Стиральная машина Beko WRE 6512 BWW', 32740, '7-d0f47f77-9053-475d-bcfd-c3264031ff34.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WRE 6512 BWW', 4.83, 7),
    ('bf59e10d-f980-4547-b3bc-2c6835729308', 'Стиральная машина Viomi Master 2 Pro black', 67490, '7-bf59e10d-f980-4547-b3bc-2c6835729308.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Viomi Master 2 Pro black', 4.15, 7),
    ('b9f44869-f68e-43e7-b3a6-868b30fd141f', 'Стиральная машина Indesit IWUC 4105 (CIS)', 30524, '7-b9f44869-f68e-43e7-b3a6-868b30fd141f.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWUC 4105 (CIS)', 4.80, 7),
    ('17c63ad5-15af-4055-ae76-fd6ed54c1f51', 'Стиральная машина Beko WSPE6H616W', 44187, '7-17c63ad5-15af-4055-ae76-fd6ed54c1f51.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WSPE6H616W', 4.83, 7),
    ('433046cf-d0f0-44b8-a3f8-08552a9b5229', 'Стиральная машина Haier HW60-BP12919B белая', 47079, '7-433046cf-d0f0-44b8-a3f8-08552a9b5229.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Haier HW60-BP12919B белая', 4.67, 7),
    ('f79a46ae-3aae-4e61-b11f-ea3badbcdb1f', 'Стиральная машина Indesit IWSD 51051 CIS', 28990, '7-f79a46ae-3aae-4e61-b11f-ea3badbcdb1f.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWSD 51051 CIS', 4.1, 7),
    ('2527129f-63b3-474c-988a-622cf2728bd4', 'Стиральная машина Indesit IWUD 4105 (CIS)', 24440, '7-2527129f-63b3-474c-988a-622cf2728bd4.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWUD 4105 (CIS)', 4.14, 7),
    ('16c892b1-a674-40f9-bd74-da15cd94b6a1', 'Стиральная машина Indesit BWSA 51051 1', 34591, '7-16c892b1-a674-40f9-bd74-da15cd94b6a1.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit BWSA 51051 1', 4.56, 7),
    ('b0be4e91-ef5d-4ea5-a0be-89cd8927c5a6', 'Стиральная машина Beko WSPE7H616W', 44100, '7-b0be4e91-ef5d-4ea5-a0be-89cd8927c5a6.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WSPE7H616W', 4.21, 7),
    ('f2779f6f-c3b3-458a-8819-a76d1a2e8b51', 'Стиральная машина LG F2T3HS6W', 59999, '7-f2779f6f-c3b3-458a-8819-a76d1a2e8b51.jpg', 'Самый лучший среди товаров на рынке Стиральная машина LG F2T3HS6W', 4.54, 7),
    ('8a397e0f-9151-4dd4-b783-4b19baf869e8', 'Стиральная машина ATLANT СМА-50 У 87', 24744, '7-8a397e0f-9151-4dd4-b783-4b19baf869e8.jpg', 'Самый лучший среди товаров на рынке Стиральная машина ATLANT СМА-50 У 87', 4.65, 7),
    ('56663fe5-972f-49bc-83ea-469954b610ec', 'Стиральная машина Beko RSPE78612W', 31860, '7-56663fe5-972f-49bc-83ea-469954b610ec.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko RSPE78612W', 4.20, 7),
    ('4e240a4d-81b7-4c03-b662-88806535fd49', 'Стиральная машина LG F2J6NM7W белый', 62420, '7-4e240a4d-81b7-4c03-b662-88806535fd49.jpg', 'Самый лучший среди товаров на рынке Стиральная машина LG F2J6NM7W белый', 4.38, 7),
    ('feffd3ab-7ab2-4892-82d4-fc03326d9f9c', 'Стиральная машина Midea MF100W60 белая', 31970, '7-feffd3ab-7ab2-4892-82d4-fc03326d9f9c.jpeg', 'Самый лучший среди товаров на рынке Стиральная машина Midea MF100W60 белая', 4.17, 7),
    ('e2f0d309-4849-4cde-b634-32dd5c6df277', 'Стиральная машина Indesit IWSB 5105 (CIS)', 26820, '7-e2f0d309-4849-4cde-b634-32dd5c6df277.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWSB 5105 (CIS)', 4.0, 7),
    ('ec3eadcf-68b5-4bcb-9b9f-ecb166cf376c', 'Стиральная машина Beko WRE 6512 BWW', 28140, '7-ec3eadcf-68b5-4bcb-9b9f-ecb166cf376c.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WRE 6512 BWW', 4.83, 7),
    ('4d2bc3c9-e416-4595-a3da-6dfaf2fe7b3b', 'Стиральная машина LG F2T3HS6S', 63449, '7-4d2bc3c9-e416-4595-a3da-6dfaf2fe7b3b.jpg', 'Самый лучший среди товаров на рынке Стиральная машина LG F2T3HS6S', 4.56, 7),
    ('ceba5194-cff5-48af-ae0a-9605de215f84', 'Стиральная машина Viomi Master 2 Pro black', 111505, '7-ceba5194-cff5-48af-ae0a-9605de215f84.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Viomi Master 2 Pro black', 4.15, 7),
    ('d6237888-ac0e-4d14-823f-ef189f4fafc0', 'Стиральная машина Indesit IWUD 4085 (CIS)', 23546, '7-d6237888-ac0e-4d14-823f-ef189f4fafc0.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWUD 4085 (CIS)', 4.38, 7),
    ('6c26cfa9-d42c-449a-b899-5b2f5b6d3608', 'Стиральная машина Beko WSPE6H616A', 42080, '7-6c26cfa9-d42c-449a-b899-5b2f5b6d3608.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WSPE6H616A', 4.6, 7),
    ('1cfc23cc-182d-436f-83b4-8b43fc21be6f', 'Стиральная машина Indesit IWSD 51051 CIS', 28990, '7-1cfc23cc-182d-436f-83b4-8b43fc21be6f.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWSD 51051 CIS', 4.1, 7),
    ('6420556a-3f7e-4b21-8ed1-ed0728d7c0cb', 'Стиральная машина Indesit IWUD 4105 (CIS)', 21048, '7-6420556a-3f7e-4b21-8ed1-ed0728d7c0cb.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWUD 4105 (CIS)', 4.14, 7),
    ('3ef4ebfc-b0de-4011-a895-e287c1adcd4c', 'Стиральная машина Beko WSPE7612W', 31020, '7-3ef4ebfc-b0de-4011-a895-e287c1adcd4c.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WSPE7612W', 4.35, 7),
    ('67ffaec0-a1de-4368-ada8-6f5f1d64763c', 'Стиральная машина Indesit IWUB 4085 (CIS)', 28092, '7-67ffaec0-a1de-4368-ada8-6f5f1d64763c.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit IWUB 4085 (CIS)', 4.32, 7),
    ('529cac83-f707-4db4-aa4d-ee7cc999f8ff', 'Стиральная машина Indesit BWSA 51051 1', 32680, '7-529cac83-f707-4db4-aa4d-ee7cc999f8ff.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Indesit BWSA 51051 1', 4.56, 7),
    ('f4755f5b-8e14-450c-a8cb-dcc66d9d24b5', 'Стиральная машина Beko WSPE7H616W', 49387, '7-f4755f5b-8e14-450c-a8cb-dcc66d9d24b5.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko WSPE7H616W', 4.21, 7),
    ('4d457713-2823-4355-ba08-753bc3ed0733', 'Стиральная машина LG F2T3HS6W', 59999, '7-4d457713-2823-4355-ba08-753bc3ed0733.jpg', 'Самый лучший среди товаров на рынке Стиральная машина LG F2T3HS6W', 4.54, 7),
    ('7e92a2d5-b0ac-47b5-949e-e572861f75a5', 'Стиральная машина Hotpoint-Ariston WDS 7448 C7S VBW', 63990, '7-7e92a2d5-b0ac-47b5-949e-e572861f75a5.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Hotpoint-Ariston WDS 7448 C7S VBW', 4.81, 7),
    ('331a0375-0ad1-4992-b268-3cbd3c6f9e79', 'Стиральная машина ATLANT СМА-50 У 87', 27445, '7-331a0375-0ad1-4992-b268-3cbd3c6f9e79.jpg', 'Самый лучший среди товаров на рынке Стиральная машина ATLANT СМА-50 У 87', 4.65, 7),
    ('c4407bcd-9178-451d-b4d0-aae09e5a3a0e', 'Стиральная машина с фронтальной загрузкой Beko WSPE7612A Black', 32360, '7-c4407bcd-9178-451d-b4d0-aae09e5a3a0e.jpg', 'Самый лучший среди товаров на рынке Стиральная машина с фронтальной загрузкой Beko WSPE7612A Black', 4.11, 7),
    ('2dde40ee-e818-429d-beb8-695017eb8f99', 'Стиральная машина Beko RSPE78612W', 36129, '7-2dde40ee-e818-429d-beb8-695017eb8f99.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Beko RSPE78612W', 4.20, 7),
    ('b7d33c43-6ded-41b7-becb-cee8dcf31f25', 'Стиральная машина Hotpoint-Ariston NSB 7249 ZD AVE RU белая', 46563, '7-b7d33c43-6ded-41b7-becb-cee8dcf31f25.jpeg', 'Самый лучший среди товаров на рынке Стиральная машина Hotpoint-Ariston NSB 7249 ZD AVE RU белая', 4.15, 7),
    ('1ddc22ac-29cc-4f09-ab92-f9f2c10c8e87', 'Стиральная машина Candy 2D1140-07', 32748, '7-1ddc22ac-29cc-4f09-ab92-f9f2c10c8e87.jpg', 'Самый лучший среди товаров на рынке Стиральная машина Candy 2D1140-07', 4.5, 7),

    --------------------Пылесосы-8
    ('fc3dd9b2-865a-4b44-a47b-41efc9df6bc6', 'Робот-пылесос Xiaomi Robot Vacuum E10 EU белый', 17960, '8-fc3dd9b2-865a-4b44-a47b-41efc9df6bc6.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Xiaomi Robot Vacuum E10 EU белый', 4.26, 8),
    ('e872c39b-cea8-458f-b84f-b225dbff7cb1', 'Вертикальный пылесос Timberk T-VCH-65 белый', 23990, '8-e872c39b-cea8-458f-b84f-b225dbff7cb1.jpeg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Timberk T-VCH-65 белый', 4.75, 8),
    ('805815a4-232c-4699-88ce-91262db599e2', 'Вертикальный пылесос Deerma Suction Vacuum Cleaner DX700S (Европейская версия)', 5000, '8-805815a4-232c-4699-88ce-91262db599e2.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma Suction Vacuum Cleaner DX700S (Европейская версия)', 4.90, 8),
    ('3b4c1217-bf02-4de1-87e3-4d8eb273f672', 'Робот-пылесос Xiaomi Robot Vacuum S10 B106GL белый', 24394, '8-3b4c1217-bf02-4de1-87e3-4d8eb273f672.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Xiaomi Robot Vacuum S10 B106GL белый', 4.97, 8),
    ('33091de3-86c3-4fba-a860-2667a6cec2d2', 'Робот-пылесос Xiaomi Robot Vacuum S10 RU белый', 29808, '8-33091de3-86c3-4fba-a860-2667a6cec2d2.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Xiaomi Robot Vacuum S10 RU белый', 4.66, 8),
    ('d582c2f2-2410-4dc1-a1ef-fdb0e22ee00c', 'Робот-пылесос Dreame L10 Ultra белый', 88124, '8-d582c2f2-2410-4dc1-a1ef-fdb0e22ee00c.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Dreame L10 Ultra белый', 4.53, 8),
    ('7925b64d-ece0-4cff-a244-69ae1c3f0ed3', 'Пылесос Scarlett SC-VC80H16', 2290, '8-7925b64d-ece0-4cff-a244-69ae1c3f0ed3.jpg', 'Самый лучший среди товаров на рынке Пылесос Scarlett SC-VC80H16', 4.35, 8),
    ('edb58c08-b4b2-41c2-8ec5-23352aea9b4f', 'Робот-пылесос Xiaomi Robot Vacuum E10 EU белый', 17008, '8-edb58c08-b4b2-41c2-8ec5-23352aea9b4f.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Xiaomi Robot Vacuum E10 EU белый', 4.26, 8),
    ('fc37d3fc-87b0-4683-97f4-a090c5535505', 'Вертикальный пылесос Deerma DX118C White', 6469, '8-fc37d3fc-87b0-4683-97f4-a090c5535505.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma DX118C White', 4.8, 8),
    ('2ddfbca9-5775-4f4f-8461-5ebe676290d8', 'Вертикальный пылесос Deerma DX700S Black', 6990, '8-2ddfbca9-5775-4f4f-8461-5ebe676290d8.jpeg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma DX700S Black', 4.69, 8),
    ('f45fb4f1-30df-4a88-85aa-523937761e8d', 'Робот-пылесос Dreame Bot L10s Ultra белый', 121754, '8-f45fb4f1-30df-4a88-85aa-523937761e8d.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Dreame Bot L10s Ultra белый', 4.3, 8),
    ('160137de-8bee-40c8-9c25-f5d027c613ce', 'Вертикальный пылесос Deerma Suction Vacuum Cleaner DX700S (Европейская версия)', 5625, '8-160137de-8bee-40c8-9c25-f5d027c613ce.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma Suction Vacuum Cleaner DX700S (Европейская версия)', 4.90, 8),
    ('878ef412-7f07-4f53-bbc1-c818c7b23de8', 'Робот-пылесос HOBOT LEGEE-D8 белый', 72430, '8-878ef412-7f07-4f53-bbc1-c818c7b23de8.png', 'Самый лучший среди товаров на рынке Робот-пылесос HOBOT LEGEE-D8 белый', 4.79, 8),
    ('0d9e4ac3-3e4b-467a-8cb9-787b696643cb', 'Вертикальный пылесос Deerma Vacuum Cleaner DX700 (Российская версия)', 8238, '8-0d9e4ac3-3e4b-467a-8cb9-787b696643cb.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma Vacuum Cleaner DX700 (Российская версия)', 4.84, 8),
    ('aa0318a6-18d0-4c95-9f16-8d0214792c94', 'Робот-пылесос Xiaomi Robot Vacuum S10 B106GL белый', 22990, '8-aa0318a6-18d0-4c95-9f16-8d0214792c94.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Xiaomi Robot Vacuum S10 B106GL белый', 4.97, 8),
    ('efa2af8b-cc5e-43b5-af22-ca7fd87d5d20', 'Пылесос Polaris PVCS 7000 Energy WAY AQUA белый', 26505, '8-efa2af8b-cc5e-43b5-af22-ca7fd87d5d20.jpeg', 'Самый лучший среди товаров на рынке Пылесос Polaris PVCS 7000 Energy WAY AQUA белый', 4.2, 8),
    ('b0a98252-f112-4c8e-adf5-eb99336cd439', 'Вертикальный пылесос Scarlett SC-VC80H23 White', 2610, '8-b0a98252-f112-4c8e-adf5-eb99336cd439.png', 'Самый лучший среди товаров на рынке Вертикальный пылесос Scarlett SC-VC80H23 White', 4.0, 8),
    ('864f0833-5856-41e8-8465-9ca42723687c', 'Робот-пылесос Dreame L10 Ultra белый', 82360, '8-864f0833-5856-41e8-8465-9ca42723687c.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Dreame L10 Ultra белый', 4.53, 8),
    ('7c0b699a-92de-4e15-a1c8-0e4cfd26a065', 'Пылесос Scarlett SC-VC80H16', 3241, '8-7c0b699a-92de-4e15-a1c8-0e4cfd26a065.jpg', 'Самый лучший среди товаров на рынке Пылесос Scarlett SC-VC80H16', 4.35, 8),
    ('34e3a65e-007b-4b85-97e4-908e4f2109af', 'Пылесос аккумуляторный Xiaomi Vacuum Cleaner G9 Plus EU (BHR6185EU)', 23429, '8-34e3a65e-007b-4b85-97e4-908e4f2109af.jpg', 'Самый лучший среди товаров на рынке Пылесос аккумуляторный Xiaomi Vacuum Cleaner G9 Plus EU (BHR6185EU)', 4.79, 8),
    ('7d770a0b-7b99-425d-b945-dff8bf87fb5a', 'Робот-пылесос Xiaomi Robot Vacuum E10 EU белый', 17253, '8-7d770a0b-7b99-425d-b945-dff8bf87fb5a.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Xiaomi Robot Vacuum E10 EU белый', 4.26, 8),
    ('ebe1c612-85d9-4bf8-8888-a03687fc7091', 'Вертикальный пылесос Deerma DX118C White', 4193, '8-ebe1c612-85d9-4bf8-8888-a03687fc7091.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma DX118C White', 4.8, 8),
    ('60326162-2025-4d74-b467-7e1714b3f25f', 'Робот-пылесос Dreame Bot L10s Ultra белый', 113430, '8-60326162-2025-4d74-b467-7e1714b3f25f.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Dreame Bot L10s Ultra белый', 4.3, 8),
    ('fe10e45e-7950-4b98-bd95-7bd4b2a23a81', 'Вертикальный пылесос Deerma Suction Vacuum Cleaner DX700S (Европейская версия)', 5659, '8-fe10e45e-7950-4b98-bd95-7bd4b2a23a81.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma Suction Vacuum Cleaner DX700S (Европейская версия)', 4.90, 8),
    ('e41965c6-f020-4881-a3ea-447d73fe27af', 'Вертикальный пылесос Deerma DX118C White', 4193, '8-e41965c6-f020-4881-a3ea-447d73fe27af.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma DX118C White', 4.8, 8),
    ('26f81af5-ce55-4f40-919d-f2fa59706f73', 'Робот-пылесос Dreame Bot L10s Ultra белый', 113430, '8-26f81af5-ce55-4f40-919d-f2fa59706f73.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Dreame Bot L10s Ultra белый', 4.3, 8),
    ('ffbd6bc7-476d-44d2-a08b-805ef714a70d', 'Вертикальный пылесос Deerma Suction Vacuum Cleaner DX700S (Европейская версия)', 5659, '8-ffbd6bc7-476d-44d2-a08b-805ef714a70d.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma Suction Vacuum Cleaner DX700S (Европейская версия)', 4.90, 8),
    ('e821a393-ad9e-4974-b703-4766d259bd9f', 'Вертикальный пылесос Dreame Vacuum Cleaner T10 (Европейская версия)', 25220, '8-e821a393-ad9e-4974-b703-4766d259bd9f.png', 'Самый лучший среди товаров на рынке Вертикальный пылесос Dreame Vacuum Cleaner T10 (Европейская версия)', 4.48, 8),
    ('f2bef340-9524-4de3-a114-0feb2a1e99d4', 'Робот-пылесос Scarlett SC-VC80R21 White', 14201, '8-f2bef340-9524-4de3-a114-0feb2a1e99d4.jpg', 'Самый лучший среди товаров на рынке Робот-пылесос Scarlett SC-VC80R21 White', 4.84, 8),
    ('bc9f52ce-665c-49f4-8574-6163b5a46d7b', 'Вертикальный пылесос Starwind  SCH1010 Black', 5030, '8-bc9f52ce-665c-49f4-8574-6163b5a46d7b.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Starwind  SCH1010 Black', 4.68, 8),
    ('ec4afbdc-8145-4afb-846f-011108954dd8', 'Вертикальный пылесос Deerma Vacuum Cleaner DX700 (Российская версия)', 6490, '8-ec4afbdc-8145-4afb-846f-011108954dd8.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma Vacuum Cleaner DX700 (Российская версия)', 4.84, 8),
    ('a0f4b0ef-4d0f-467f-88ec-ac139482a6f6', 'Робот-пылесос Roborock Q7 Max Plus (Русская версия) White', 76799, '8-a0f4b0ef-4d0f-467f-88ec-ac139482a6f6.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Roborock Q7 Max Plus (Русская версия) White', 4.81, 8),
    ('b5e4235b-6551-4309-a544-39e0ba3f306b', 'Вертикальный пылесос Deerma DX700 (Европейская версия)', 3958, '8-b5e4235b-6551-4309-a544-39e0ba3f306b.jpeg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma DX700 (Европейская версия)', 4.93, 8),
    ('86a75039-e496-4ec2-b94e-c5df0d08f093', 'Пылесос LG  VK 69662 N Blue', 6690, '8-86a75039-e496-4ec2-b94e-c5df0d08f093.jpg', 'Самый лучший среди товаров на рынке Пылесос LG  VK 69662 N Blue', 4.70, 8),
    ('2aebfe5e-f303-4afc-a76e-8c81aa5acd6d', 'Вертикальный пылесос Scarlett SC-VC80H22 White', 4780, '8-2aebfe5e-f303-4afc-a76e-8c81aa5acd6d.jpeg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Scarlett SC-VC80H22 White', 4.80, 8),
    ('cab23fe3-1077-48f9-a6c6-d1534d87c96e', 'Пылесос LG  VK 69662 N Blue', 13322, '8-cab23fe3-1077-48f9-a6c6-d1534d87c96e.jpg', 'Самый лучший среди товаров на рынке Пылесос LG  VK 69662 N Blue', 4.70, 8),
    ('2ba91979-8d0f-4ea3-9b41-5b5d19f99fba', 'Пылесос Scarlett SC-VC80H16', 3241, '8-2ba91979-8d0f-4ea3-9b41-5b5d19f99fba.jpg', 'Самый лучший среди товаров на рынке Пылесос Scarlett SC-VC80H16', 4.35, 8),
    ('f995379f-835b-4ea2-b242-25fbec8d2f04', 'Пылесос аккумуляторный Xiaomi Vacuum Cleaner G9 Plus EU (BHR6185EU)', 27681, '8-f995379f-835b-4ea2-b242-25fbec8d2f04.jpg', 'Самый лучший среди товаров на рынке Пылесос аккумуляторный Xiaomi Vacuum Cleaner G9 Plus EU (BHR6185EU)', 4.79, 8),
    ('d177d946-33ff-4621-80ec-38e21361157e', 'Робот-пылесос Roborock S8 Pro Ultra (Русская версия) белый', 115990, '8-d177d946-33ff-4621-80ec-38e21361157e.png', 'Самый лучший среди товаров на рынке Робот-пылесос Roborock S8 Pro Ultra (Русская версия) белый', 4.80, 8),
    ('3538726d-9ae5-4498-95c1-bd6d84366738', 'Вертикальный пылесос Deerma Vacuum Cleaner DX700 (Российская версия)', 4909, '8-3538726d-9ae5-4498-95c1-bd6d84366738.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma Vacuum Cleaner DX700 (Российская версия)', 4.84, 8),
    ('c07a4aff-2030-470a-a901-dec9adb78d63', 'Робот-пылесос Dreame D10 Plus белый', 49318, '8-c07a4aff-2030-470a-a901-dec9adb78d63.jpg', 'Самый лучший среди товаров на рынке Робот-пылесос Dreame D10 Plus белый', 4.83, 8),
    ('c7a018cb-7d40-426b-9f88-b7686183f344', 'Робот-пылесос Xiaomi Robot Vacuum S10+EU белый', 53521, '8-c7a018cb-7d40-426b-9f88-b7686183f344.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Xiaomi Robot Vacuum S10+EU белый', 4.23, 8),
    ('9c8222be-e280-4611-a01c-c7d9e9642859', 'Робот-пылесос Xiaomi Robot Vacuum E10 EU белый', 17178, '8-9c8222be-e280-4611-a01c-c7d9e9642859.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Xiaomi Robot Vacuum E10 EU белый', 4.26, 8),
    ('ffadecde-0096-4ed8-89e7-3003edcf66af', 'Вертикальный пылесос Deerma DX118C White', 2686, '8-ffadecde-0096-4ed8-89e7-3003edcf66af.jpg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma DX118C White', 4.8, 8),
    ('977b004e-bedd-41cb-bf1f-aa9ef031978a', 'Вертикальный пылесос Deerma DX700S Black', 3950, '8-977b004e-bedd-41cb-bf1f-aa9ef031978a.jpeg', 'Самый лучший среди товаров на рынке Вертикальный пылесос Deerma DX700S Black', 4.69, 8),
    ('d0ba8212-06f7-45d5-b9c5-de9ac4649eb0', 'Робот-пылесос Scarlett SC-VC80R21 White', 13027, '8-d0ba8212-06f7-45d5-b9c5-de9ac4649eb0.jpg', 'Самый лучший среди товаров на рынке Робот-пылесос Scarlett SC-VC80R21 White', 4.84, 8),
    ('98d8a525-9568-4b95-b3b1-37e6e6f3eef7', 'Робот-пылесос Dreame Bot L10s Ultra белый', 102350, '8-98d8a525-9568-4b95-b3b1-37e6e6f3eef7.jpeg', 'Самый лучший среди товаров на рынке Робот-пылесос Dreame Bot L10s Ultra белый', 4.3, 8),
    ('516af109-29e8-4c85-9329-7ff70d27216d', 'Робот-пылесос Honor Choice Robot Cleaner R2 White белый', 29178, '8-516af109-29e8-4c85-9329-7ff70d27216d.png', 'Самый лучший среди товаров на рынке Робот-пылесос Honor Choice Robot Cleaner R2 White белый', 4.61, 8);
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

