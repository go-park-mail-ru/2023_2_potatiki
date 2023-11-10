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
    (8, 'Пылесосы', 5),
    (9, 'Мебель', 1),
    (91, 'Стулья', 9),
    (92, 'Рабочии столы', 9),
    (93, 'Диваны', 9),
    (94, 'Кресла', 9);

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

INSERT INTO product (id, name, price, imgsrc, description, rating, category_id)
	VALUES

	('1e461708-6b04-45b9-a4fa-77c32c14d382', 'Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 35', 35990, '93-1e461708-6b04-45b9-a4fa-77c32c14d382.jpeg', 'Самый лучший среди товаров на рынке Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 35', 4.16, 93),
	('0545799d-b57a-4930-a66c-d4bdfd0ed03d', 'Угловой диван-кровать Gupan Nordkisa, Еврокнижка, ППУ, цвет Amigo Navy, угол слева', 51199, '93-0545799d-b57a-4930-a66c-d4bdfd0ed03d.png', 'Самый лучший среди товаров на рынке Угловой диван-кровать Gupan Nordkisa, Еврокнижка, ППУ, цвет Amigo Navy, угол слева', 4.95, 93),
	('e6ac9016-3b2b-4c01-a718-24308803461a', 'Диван-кровать МногоМеб Стелла серый велюр, еврокнижка', 18520, '93-e6ac9016-3b2b-4c01-a718-24308803461a.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать МногоМеб Стелла серый велюр, еврокнижка', 4.30, 93),
	('f31cef95-7b69-4a16-8808-bb279389d4ef', 'Диван-кровать МногоМеб Книжка в велюре Астра серый 120 ППУ', 10200, '93-f31cef95-7b69-4a16-8808-bb279389d4ef.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать МногоМеб Книжка в велюре Астра серый 120 ППУ', 4.5, 93),
	('bc61c5ee-2e18-4d9b-9602-f67f8ab31c1f', 'Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 25', 35990, '93-bc61c5ee-2e18-4d9b-9602-f67f8ab31c1f.jpeg', 'Самый лучший среди товаров на рынке Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 25', 4.15, 93),
	('2189223a-8ac5-4dd3-a70c-2067e3781080', 'Диван прямой ОБСТАНОВКА.РУ Босс Фабио Melissa 35', 36000, '93-2189223a-8ac5-4dd3-a70c-2067e3781080.jpeg', 'Самый лучший среди товаров на рынке Диван прямой ОБСТАНОВКА.РУ Босс Фабио Melissa 35', 4.6, 93),
	('e6117881-7d1d-483c-a808-09070bdbea23', 'Угловой диван-кровать Gupan Nordkisa, Еврокнижка, ППУ, цвет Amigo Grafit, угол справа', 51199, '93-e6117881-7d1d-483c-a808-09070bdbea23.png', 'Самый лучший среди товаров на рынке Угловой диван-кровать Gupan Nordkisa, Еврокнижка, ППУ, цвет Amigo Grafit, угол справа', 4.95, 93),
	('5d5fba4b-641e-4cc1-b455-23ab472495f7', 'Диван-кровать NRAVA Hit 2180х870х940 Corvette 18 Серый', 35447, '93-5d5fba4b-641e-4cc1-b455-23ab472495f7.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать NRAVA Hit 2180х870х940 Corvette 18 Серый', 4.16, 93),
	('2c55d2b3-da06-4c07-be0d-175d35647788', 'Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 39', 35990, '93-2c55d2b3-da06-4c07-be0d-175d35647788.jpeg', 'Самый лучший среди товаров на рынке Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 39', 4.3, 93),
	('25692680-7fd4-4f77-8fba-783fe18f5461', 'Диван-кровать Gupan Норд, материал Велюр, Amigo Blue, беспружинный', 44382, '93-25692680-7fd4-4f77-8fba-783fe18f5461.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать Gupan Норд, материал Велюр, Amigo Blue, беспружинный', 4.95, 93),
	('51c67ffe-47ef-4be2-9a8f-3ab3b47ef7c8', 'Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 2', 35990, '93-51c67ffe-47ef-4be2-9a8f-3ab3b47ef7c8.jpeg', 'Самый лучший среди товаров на рынке Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 2', 4.7, 93),
	('40060c76-43bf-4faa-a4a3-866aec29488f', 'Диван-кровать NRAVA Fabi 1420х770х808 Vivaldi 4 бежевый', 28245, '93-40060c76-43bf-4faa-a4a3-866aec29488f.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать NRAVA Fabi 1420х770х808 Vivaldi 4 бежевый', 4.10, 93),
	('081e4600-a353-4262-8ae8-ffb743a5abf4', 'Угловой диван Лига Диванов Атланта лайт правый угол', 36990, '93-081e4600-a353-4262-8ae8-ffb743a5abf4.jpg', 'Самый лучший среди товаров на рынке Угловой диван Лига Диванов Атланта лайт правый угол', 4.4, 93),
	('3b51331e-c63b-459f-8440-7b2f63966ec0', 'Диван прямой ОБСТАНОВКА.РУ Босс Фабио Melissa 35', 34790, '93-3b51331e-c63b-459f-8440-7b2f63966ec0.jpeg', 'Самый лучший среди товаров на рынке Диван прямой ОБСТАНОВКА.РУ Босс Фабио Melissa 35', 4.6, 93),
	('2cdf4317-c177-41eb-971b-832b260b1056', 'Диван SALON TRON Монреаль серый, в холл, в прихожую, в гостиную, 120х54х75', 15990, '93-2cdf4317-c177-41eb-971b-832b260b1056.jpeg', 'Самый лучший среди товаров на рынке Диван SALON TRON Монреаль серый, в холл, в прихожую, в гостиную, 120х54х75', 4.2, 93),
	('a129acf0-d743-4af6-b09d-b5088780672d', 'Диван-кровать Элегантный Стиль угол универсальный Комо, серый', 18490, '93-a129acf0-d743-4af6-b09d-b5088780672d.jpg', 'Самый лучший среди товаров на рынке Диван-кровать Элегантный Стиль угол универсальный Комо, серый', 4.2, 93),
	('8c2790aa-81c3-4dca-a72b-7434b5cf37c7', 'Диван-кровать NRAVA Fabi 1420х770х810 Vivaldi 13 синий', 28245, '93-8c2790aa-81c3-4dca-a72b-7434b5cf37c7.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать NRAVA Fabi 1420х770х810 Vivaldi 13 синий', 4.17, 93),
	('8b0be6bf-7035-455a-8312-46b1cf4dfb8c', 'Диван-кровать еврокнижка МногоМеб Севилья рогожка серая', 15990, '93-8b0be6bf-7035-455a-8312-46b1cf4dfb8c.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать еврокнижка МногоМеб Севилья рогожка серая', 4.2, 93),
	('b1de41d9-ce46-48fb-9bfe-92389f25a383', 'Угловой диван Лига Диванов Атланта лайт правый угол', 36990, '93-b1de41d9-ce46-48fb-9bfe-92389f25a383.jpg', 'Самый лучший среди товаров на рынке Угловой диван Лига Диванов Атланта лайт правый угол', 4.4, 93),
	('53ce80d5-a931-4210-b218-2c95487d8ba7', 'Угловой диван Лига Диванов Атланта лайт левый угол', 36990, '93-53ce80d5-a931-4210-b218-2c95487d8ba7.jpg', 'Самый лучший среди товаров на рынке Угловой диван Лига Диванов Атланта лайт левый угол', 4.1, 93),
	('a8c2b26a-755e-40ae-844a-e2d6493322fe', 'Диван-кровать Многомеб Книжка в велюре Астра светло-серый 120 ППУ', 10200, '93-a8c2b26a-755e-40ae-844a-e2d6493322fe.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать Многомеб Книжка в велюре Астра светло-серый 120 ППУ', 4.1, 93),
	('83a17a06-ee91-4933-9bd8-8ef9fc3fc76a', 'Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 20', 35990, '93-83a17a06-ee91-4933-9bd8-8ef9fc3fc76a.jpeg', 'Самый лучший среди товаров на рынке Диван прямой ОБСТАНОВКА.РУ Босс Фабио Анатомик Melissa 20', 4.5, 93),
	('4d43e38d-1e93-40c2-a0cb-d0ab1af86372', 'Диван угловой Остин, Зара Крем', 55990, '93-4d43e38d-1e93-40c2-a0cb-d0ab1af86372.jpeg', 'Самый лучший среди товаров на рынке Диван угловой Остин, Зара Крем', 4.1, 93),
	('2963081a-7f68-419c-8848-006205acab8f', 'Диван прямой Dihall Сиэтл King, Зара Грей', 42290, '93-2963081a-7f68-419c-8848-006205acab8f.jpeg', 'Самый лучший среди товаров на рынке Диван прямой Dihall Сиэтл King, Зара Грей', 4.3, 93),
	('f77c9f06-c3e2-4a31-ac8a-91135d53901a', 'Диван-кровать еврокнижка МногоМеб Леон велюр принт Москва', 16780, '93-f77c9f06-c3e2-4a31-ac8a-91135d53901a.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать еврокнижка МногоМеб Леон велюр принт Москва', 4.13, 93),
	('6f66fb05-b7ac-4199-a546-42ca05bb9d57', 'Диван для кухни ТриЯ Форест угловой Дуб крафт белый/Велюр серый', 32599, '93-6f66fb05-b7ac-4199-a546-42ca05bb9d57.jpg', 'Самый лучший среди товаров на рынке Диван для кухни ТриЯ Форест угловой Дуб крафт белый/Велюр серый', 4.1, 93),
	('674caef7-1464-449a-8a8b-883f128ad4da', 'Диван-кровать МногоМеб Стелла синий велюр, еврокнижка', 18520, '93-674caef7-1464-449a-8a8b-883f128ad4da.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать МногоМеб Стелла синий велюр, еврокнижка', 4.8, 93),
	('17d4dc20-7937-466f-990f-5a6a2c474717', 'Диван MONOFIX, БАФФ экокожа серый', 9006, '93-17d4dc20-7937-466f-990f-5a6a2c474717.jpeg', 'Самый лучший среди товаров на рынке Диван MONOFIX, БАФФ экокожа серый', 4.3, 93),
	('fb2807e7-4567-41dc-ba15-3f717c59d595', 'Диван-кровать NRAVA Fabi 1420х770х810 Vivaldi 7 серый', 28245, '93-fb2807e7-4567-41dc-ba15-3f717c59d595.jpeg', 'Самый лучший среди товаров на рынке Диван-кровать NRAVA Fabi 1420х770х810 Vivaldi 7 серый', 4.6, 93),
	('3263435e-edcf-43e3-af32-bd182bf40d4d', 'Диван MONOFIX АММА, серый, микровелюр', 13172, '93-3263435e-edcf-43e3-af32-bd182bf40d4d.jpeg', 'Самый лучший среди товаров на рынке Диван MONOFIX АММА, серый, микровелюр', 4.4, 93),
	('3ce40a85-acf0-4633-be1f-e53a55ba5245', 'Диван MONOFIX, БАФФ экокожа черный', 9006, '93-3ce40a85-acf0-4633-be1f-e53a55ba5245.jpeg', 'Самый лучший среди товаров на рынке Диван MONOFIX, БАФФ экокожа черный', 4.5, 93),
	('28238713-0c0b-4b05-a52b-c7a14233b2be', 'Диван, SALON TRON, Монреаль чёрный, 120х54х75', 15990, '93-28238713-0c0b-4b05-a52b-c7a14233b2be.jpeg', 'Самый лучший среди товаров на рынке Диван, SALON TRON, Монреаль чёрный, 120х54х75', 4.10, 93),
	('00aadf3e-bf85-439d-a376-78de2b475ffc', 'Диван MONOFIX, БАФФ экокожа серый', 6450, '93-00aadf3e-bf85-439d-a376-78de2b475ffc.jpeg', 'Самый лучший среди товаров на рынке Диван MONOFIX, БАФФ экокожа серый', 4.3, 93),
	('1adff586-4f77-4c1b-9de2-1397388274a1', 'Диван прямой ОБСТАНОВКА.РУ Босс Фабио Melissa 25', 34790, '93-1adff586-4f77-4c1b-9de2-1397388274a1.jpeg', 'Самый лучший среди товаров на рынке Диван прямой ОБСТАНОВКА.РУ Босс Фабио Melissa 25', 4.1, 93);


INSERT INTO product (id, name, price, imgsrc, description, rating, category_id)
	VALUES

	('1bbddb00-0eef-4982-8695-9c0f87ce2f18', 'Холодильник Indesit DS4180W White', 44590, '6-1bbddb00-0eef-4982-8695-9c0f87ce2f18.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit DS4180W White', 4.16, 6),
	('b76b1965-6d0f-4819-ab1e-23fda85366bd', 'Холодильник Indesit DS4200W White', 46854, '6-b76b1965-6d0f-4819-ab1e-23fda85366bd.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit DS4200W White', 4.41, 6),
	('8a65166f-56ee-4883-b773-01039d52e0ea', 'Холодильник Indesit ITR 5200 W', 63805, '6-8a65166f-56ee-4883-b773-01039d52e0ea.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit ITR 5200 W', 4.28, 6),
	('7d0b2651-10e3-4119-8768-7bc2951d8e06', 'Холодильник Бирюса М90', 24490, '6-7d0b2651-10e3-4119-8768-7bc2951d8e06.jpg', 'Самый лучший среди товаров на рынке Холодильник Бирюса М90', 4.49, 6),
	('e8d80216-e6a0-4b11-bdeb-68ee3d226a77', 'Холодильник Indesit DS4180W White', 30408, '6-e8d80216-e6a0-4b11-bdeb-68ee3d226a77.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit DS4180W White', 4.16, 6),
	('9b3cc994-c21d-4889-b853-121152131329', 'Холодильник Бирюса 50 White', 8170, '6-9b3cc994-c21d-4889-b853-121152131329.jpg', 'Самый лучший среди товаров на рынке Холодильник Бирюса 50 White', 4.47, 6),
	('34706d43-49f3-4008-ba32-54ca2243ee4c', 'Холодильник Бирюса Б-70 White', 22990, '6-34706d43-49f3-4008-ba32-54ca2243ee4c.jpg', 'Самый лучший среди товаров на рынке Холодильник Бирюса Б-70 White', 4.99, 6),
	('0230c03d-2e70-4ab3-9e78-bcfd985e8e6a', 'Холодильник Beko RCNK335E20VW', 63773, '6-0230c03d-2e70-4ab3-9e78-bcfd985e8e6a.jpg', 'Самый лучший среди товаров на рынке Холодильник Beko RCNK335E20VW', 4.14, 6),
	('324fa260-74a8-434b-b295-97ef6333d3c0', 'Холодильник Gorenje RK4181PS4', 52180, '6-324fa260-74a8-434b-b295-97ef6333d3c0.jpg', 'Самый лучший среди товаров на рынке Холодильник Gorenje RK4181PS4', 4.51, 6),
	('b56721c9-b3b8-4590-972b-0bbcf015d5a1', 'Холодильник Beko RCSK 270M20 W White', 43560, '6-b56721c9-b3b8-4590-972b-0bbcf015d5a1.jpg', 'Самый лучший среди товаров на рынке Холодильник Beko RCSK 270M20 W White', 4.63, 6),
	('fa4a1ad9-7046-4b7f-9166-33033291a593', 'Холодильник Бирюса М90', 17248, '6-fa4a1ad9-7046-4b7f-9166-33033291a593.jpg', 'Самый лучший среди товаров на рынке Холодильник Бирюса М90', 4.49, 6),
	('122dd32c-764c-45d9-878b-d8d98a26016e', 'Холодильник Indesit ITR 4180 W White', 40850, '6-122dd32c-764c-45d9-878b-d8d98a26016e.png', 'Самый лучший среди товаров на рынке Холодильник Indesit ITR 4180 W White', 4.82, 6),
	('169ee350-fbc6-4910-acce-801ad0c737ea', 'Холодильник Indesit ITR 4200 S', 57800, '6-169ee350-fbc6-4910-acce-801ad0c737ea.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit ITR 4200 S', 4.56, 6),
	('087e79af-3b9c-44f4-94ca-224b5a66021d', 'Холодильник Indesit DS4180W White', 34270, '6-087e79af-3b9c-44f4-94ca-224b5a66021d.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit DS4180W White', 4.16, 6),
	('e1e4ba83-7f7e-4d3b-9e80-5c11d2dd6e9a', 'Холодильник Бирюса 50 White', 17610, '6-e1e4ba83-7f7e-4d3b-9e80-5c11d2dd6e9a.jpg', 'Самый лучший среди товаров на рынке Холодильник Бирюса 50 White', 4.47, 6),
	('70742d73-d935-4f48-b372-5015e97b648d', 'Холодильник Indesit ITR 5200 S', 46360, '6-70742d73-d935-4f48-b372-5015e97b648d.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit ITR 5200 S', 4.65, 6),
	('47c3be90-e223-4cf1-b4aa-4781d45a068a', 'Холодильник Indesit ITR 5180 W', 54220, '6-47c3be90-e223-4cf1-b4aa-4781d45a068a.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit ITR 5180 W', 4.6, 6),
	('e509652c-de7f-466a-89e8-b6ca5f6284d5', 'Холодильник Indesit DS4200W White', 39500, '6-e509652c-de7f-466a-89e8-b6ca5f6284d5.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit DS4200W White', 4.41, 6),
	('7e45e06f-feb6-4fbc-b8a6-c43dc7ec3ded', 'Холодильник Beko RCNK335E20VW', 59970, '6-7e45e06f-feb6-4fbc-b8a6-c43dc7ec3ded.jpg', 'Самый лучший среди товаров на рынке Холодильник Beko RCNK335E20VW', 4.14, 6),
	('f2af1bd8-6edb-4403-a8a5-c309561dff55', 'Холодильник Indesit TIA16 White', 38110, '6-f2af1bd8-6edb-4403-a8a5-c309561dff55.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit TIA16 White', 4.6, 6),
	('4e665c5a-515a-45ef-a441-ad47a2fefe69', 'Холодильник Indesit ITR 4200 W White', 53930, '6-4e665c5a-515a-45ef-a441-ad47a2fefe69.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit ITR 4200 W White', 4.66, 6),
	('15218fff-2851-4b94-947a-369760886cae', 'Холодильник Gorenje RK4181PS4', 42670, '6-15218fff-2851-4b94-947a-369760886cae.jpg', 'Самый лучший среди товаров на рынке Холодильник Gorenje RK4181PS4', 4.51, 6),
	('de0e6fbc-4715-4fe7-a195-69fa87d6c283', 'Холодильник Haier C2F636CFRG Silver', 103340, '6-de0e6fbc-4715-4fe7-a195-69fa87d6c283.jpg', 'Самый лучший среди товаров на рынке Холодильник Haier C2F636CFRG Silver', 4.58, 6),
	('02eed52f-e3e8-4be3-9e17-877f85949917', 'Холодильник Indesit DS4160W White', 32300, '6-02eed52f-e3e8-4be3-9e17-877f85949917.jpg', 'Самый лучший среди товаров на рынке Холодильник Indesit DS4160W White', 4.15, 6);

INSERT INTO product (id, name, price, imgsrc, description, rating, category_id)
	VALUES

	('13f45146-6cb5-4302-b2e0-a527f4d8db6d', 'Кресло-мешок Happy-puff Груша XXXXL-Комфорт W35', 3999, '94-13f45146-6cb5-4302-b2e0-a527f4d8db6d.png', 'Самый лучший среди товаров на рынке Кресло-мешок Happy-puff Груша XXXXL-Комфорт W35', 4.29, 94),
	('827cb9ed-1012-43f5-be39-c3ca273b46cd', 'Кресло Кресло для кухни со спинкой обеденное кухонное Tetchair BREMO', 6640, '94-827cb9ed-1012-43f5-be39-c3ca273b46cd.jpg', 'Самый лучший среди товаров на рынке Кресло Кресло для кухни со спинкой обеденное кухонное Tetchair BREMO', 4.95, 94),
	('5e53dfbe-27a1-4430-96ea-d5365d6fca96', 'Кресло мешок PUFOFF XXXL Kittens', 3990, '94-5e53dfbe-27a1-4430-96ea-d5365d6fca96.jpeg', 'Самый лучший среди товаров на рынке Кресло мешок PUFOFF XXXL Kittens', 4.29, 94),
	('12921ee4-b8bd-4520-b73d-496f1384f3f4', 'Кресло-мешок ONPUFF пуфик груша, размер XXXL, серый оксфорд', 3490, '94-12921ee4-b8bd-4520-b73d-496f1384f3f4.jpg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXL, серый оксфорд', 4.8, 94),
	('f0b12213-46bb-4749-8476-f558736c89b2', 'Кресло Оскар Velour 8', 8890, '94-f0b12213-46bb-4749-8476-f558736c89b2.jpeg', 'Самый лучший среди товаров на рынке Кресло Оскар Velour 8', 4.4, 94),
	('415215e9-f069-454a-876d-5ba8989b2f55', 'Кресло-мешок ONPUFF пуфик груша, размер XXXXL, серый оксфорд', 3990, '94-415215e9-f069-454a-876d-5ba8989b2f55.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXXL, серый оксфорд', 4.7, 94),
	('10c06d62-773b-4210-97dd-5eb9e0879525', 'Кресло-мешок ONPUFF пуфик груша, размер XXXXL, бирюзовый оксфорд', 3990, '94-10c06d62-773b-4210-97dd-5eb9e0879525.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXXL, бирюзовый оксфорд', 4.4, 94),
	('e7b83918-fc70-4284-a717-020d02dbc2c7', 'Кресло-мешок ONPUFF пуфик груша, размер XXXXL, серый рогожка', 5990, '94-e7b83918-fc70-4284-a717-020d02dbc2c7.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXXL, серый рогожка', 4.11, 94),
	('ebc5269b-0c61-4b45-8313-38bc3ad6b80f', 'Кресло-мешок PUFON Груша XXXL, серый', 3144, '94-ebc5269b-0c61-4b45-8313-38bc3ad6b80f.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок PUFON Груша XXXL, серый', 4.6, 94),
	('d33a067d-ec0b-4906-ad53-c4a88e4772f8', 'Кресло-качалка с подножкой Glider Аоста', 11197, '94-d33a067d-ec0b-4906-ad53-c4a88e4772f8.jpeg', 'Самый лучший среди товаров на рынке Кресло-качалка с подножкой Glider Аоста', 4.6, 94),
	('66e6b262-41d9-4b67-9be7-cadb85e5bd88', 'Кресло-мешок ONPUFF пуфик груша, размер XXXXL, серый оксфорд', 3990, '94-66e6b262-41d9-4b67-9be7-cadb85e5bd88.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXXL, серый оксфорд', 4.7, 94),
	('962ca7c5-bac1-443b-a128-218601d4b872', 'Кресло-мешок ONPUFF пуфик груша, размер XXXXL, коричневый оксфорд', 3990, '94-962ca7c5-bac1-443b-a128-218601d4b872.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXXL, коричневый оксфорд', 4.2, 94),
	('27f7175e-93aa-47bf-b517-0a283c9e7a5f', 'Кресло-мешок ONPUFF пуфик груша, размер XXXXL, бирюзовый оксфорд', 3990, '94-27f7175e-93aa-47bf-b517-0a283c9e7a5f.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXXL, бирюзовый оксфорд', 4.4, 94),
	('c0b59bd3-ca55-465c-abc9-be5dbd6e44f1', 'Кресло-мешок PUFON Груша XL, черный', 1784, '94-c0b59bd3-ca55-465c-abc9-be5dbd6e44f1.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок PUFON Груша XL, черный', 4.10, 94),
	('39df1714-769e-4f34-a5bc-a08ef33c4141', 'Кресло-мешок PUFLOVE пуфик груша, размер XXXXL, серый велюр', 5990, '94-39df1714-769e-4f34-a5bc-a08ef33c4141.jpg', 'Самый лучший среди товаров на рынке Кресло-мешок PUFLOVE пуфик груша, размер XXXXL, серый велюр', 4.6, 94),
	('9760c021-05b7-451b-bc5b-f227355ab79a', 'Кресло-мешок ONPUFF пуфик груша, размер XXXXL, серый рогожка', 5990, '94-9760c021-05b7-451b-bc5b-f227355ab79a.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXXL, серый рогожка', 4.11, 94),
	('d2db2181-1301-4303-b598-09dccffda092', 'Кресло-мешок PUFON Груша XXXL, серый', 3699, '94-d2db2181-1301-4303-b598-09dccffda092.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок PUFON Груша XXXL, серый', 4.6, 94),
	('6e7ed2aa-c188-4359-a617-efdb4f6cc116', 'Кресло Оскар Crash beige 33', 10290, '94-6e7ed2aa-c188-4359-a617-efdb4f6cc116.jpeg', 'Самый лучший среди товаров на рынке Кресло Оскар Crash beige 33', 4.4, 94),
	('f0510c84-61de-4d79-94b0-be5a7e0098c8', 'Кресло Оскар Dream Lux 19', 8690, '94-f0510c84-61de-4d79-94b0-be5a7e0098c8.jpeg', 'Самый лучший среди товаров на рынке Кресло Оскар Dream Lux 19', 4.3, 94),
	('07013117-1006-4e8f-bfb0-b8669701023f', 'Кресло-мешок ONPUFF пуфик груша, размер XXXXL, коричневый оксфорд', 3990, '94-07013117-1006-4e8f-bfb0-b8669701023f.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXXL, коричневый оксфорд', 4.2, 94),
	('26698202-6008-462a-b531-b269ecec2ec8', 'Кресло-мешок PUFON Груша XL, черный', 2099, '94-26698202-6008-462a-b531-b269ecec2ec8.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок PUFON Груша XL, черный', 4.10, 94),
	('5a0b33b7-6810-49b5-803e-b7174c9cc249', 'Кресло-мешок ONPUFF пуфик груша, размер XXXL, бирюзовый оксфорд', 3490, '94-5a0b33b7-6810-49b5-803e-b7174c9cc249.jpg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXL, бирюзовый оксфорд', 4.5, 94),
	('db1c8979-079a-495d-a7aa-7fcd6433fb47', 'Кресло-мешок «Мяч», XL (95x95), оксфорд, Белый и черный', 2244, '94-db1c8979-079a-495d-a7aa-7fcd6433fb47.jpg', 'Самый лучший среди товаров на рынке Кресло-мешок «Мяч», XL (95x95), оксфорд, Белый и черный', 4.12, 94),
	('957b2a4c-95ab-4668-9ce6-c736acb19569', 'Кресло-качалка с подножкой, мятниковый механизм Glider Стронг', 16395, '94-957b2a4c-95ab-4668-9ce6-c736acb19569.jpeg', 'Самый лучший среди товаров на рынке Кресло-качалка с подножкой, мятниковый механизм Glider Стронг', 4.4, 94),
	('a3a6103a-bb86-4c71-b8b5-433cabfb76bf', 'Кресло-мешок ONPUFF пуфик груша, размер XXXL, коричневый оксфорд', 3490, '94-a3a6103a-bb86-4c71-b8b5-433cabfb76bf.jpg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXL, коричневый оксфорд', 4.1, 94),
	('3a7b66fd-78b0-41c1-a588-a6079e9835bd', 'Кресло-качалка с маятниковым механизмом Glider 68', 10990, '94-3a7b66fd-78b0-41c1-a588-a6079e9835bd.jpeg', 'Самый лучший среди товаров на рынке Кресло-качалка с маятниковым механизмом Glider 68', 4.4, 94),
	('e4f6d489-c19d-4a66-902d-5fc2cb7863b6', 'Кресло-мешок ONPUFF пуфик груша, размер XXXL, бирюзовый оксфорд', 3490, '94-e4f6d489-c19d-4a66-902d-5fc2cb7863b6.jpg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXL, бирюзовый оксфорд', 4.5, 94),
	('62bfdd84-10f0-4817-82a5-1defdb5aaa3c', 'Кресло-мешок PUFON Груша XXXXL, черный', 3654, '94-62bfdd84-10f0-4817-82a5-1defdb5aaa3c.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок PUFON Груша XXXXL, черный', 4.6, 94),
	('122aedc7-77e4-4485-9b58-04bc00350750', 'Кресло Оскар Crash silver 10', 10290, '94-122aedc7-77e4-4485-9b58-04bc00350750.jpeg', 'Самый лучший среди товаров на рынке Кресло Оскар Crash silver 10', 4.2, 94),
	('320c6b9b-2baf-445e-adbc-c1cc2f500913', 'Кресло-мешок «Груша», L (85x70), оксфорд, Бежевый', 1318, '94-320c6b9b-2baf-445e-adbc-c1cc2f500913.jpg', 'Самый лучший среди товаров на рынке Кресло-мешок «Груша», L (85x70), оксфорд, Бежевый', 4.8, 94),
	('09664716-6548-4f1f-8476-78450b728908', 'Кресло-качалка Glider Экси', 11934, '94-09664716-6548-4f1f-8476-78450b728908.jpeg', 'Самый лучший среди товаров на рынке Кресло-качалка Glider Экси', 4.4, 94),
	('e1f654db-ac78-476b-8161-548bdff2f4e2', 'Кресло-мешок папа пуф оксфорд коричневый 3xl 150x100', 2990, '94-e1f654db-ac78-476b-8161-548bdff2f4e2.jpeg', 'Самый лучший среди товаров на рынке Кресло-мешок папа пуф оксфорд коричневый 3xl 150x100', 4.18, 94),
	('483524f2-c6f0-4022-bbdc-1fdab691191c', 'Кресло-мешок ONPUFF пуфик груша, размер XXXL, коричневый оксфорд', 3490, '94-483524f2-c6f0-4022-bbdc-1fdab691191c.jpg', 'Самый лучший среди товаров на рынке Кресло-мешок ONPUFF пуфик груша, размер XXXL, коричневый оксфорд', 4.1, 94);


INSERT INTO product (id, name, price, imgsrc, description, rating, category_id)
	VALUES

	('0bf91a38-a463-4323-b0e3-971fa11a154a', 'Стол игровой Defender Electro RGB с подставкой под кружку и подвесом под гарнитуру, чёрный', 10765, '92-0bf91a38-a463-4323-b0e3-971fa11a154a.jpg', 'Самый лучший среди товаров на рынке Стол игровой Defender Electro RGB с подставкой под кружку и подвесом под гарнитуру, чёрный', 4.37, 92),
	('5724a3e1-8191-4606-815e-94af8c523a2a', 'Стол письменный Империал Гарвард угловой Полка-Надставка 120 ПР Дуб сонома/Белый', 22999, '92-5724a3e1-8191-4606-815e-94af8c523a2a.jpeg', 'Самый лучший среди товаров на рынке Стол письменный Империал Гарвард угловой Полка-Надставка 120 ПР Дуб сонома/Белый', 4.95, 92),
	('0ff6af49-c715-4ac3-bd6d-90870c666a99', 'Стол игровой Hiper Titan HGTBL001, с подсветкой, черный', 13374, '92-0ff6af49-c715-4ac3-bd6d-90870c666a99.jpg', 'Самый лучший среди товаров на рынке Стол игровой Hiper Titan HGTBL001, с подсветкой, черный', 4.27, 92),
	('2846bca6-2066-4310-ba5e-a8427c5a0496', 'Стол Уютная логика МС-7 Квазар 3 ящика, 130*78,4*60 см, белый', 8400, '92-2846bca6-2066-4310-ba5e-a8427c5a0496.png', 'Самый лучший среди товаров на рынке Стол Уютная логика МС-7 Квазар 3 ящика, 130*78,4*60 см, белый', 4.2, 92),
	('fa2f7089-cb5a-4770-a387-8f5b4cc72b6e', 'Стол компьютерный с ящиками белый Frenesie, ЛДСП, 90х45х77 см.', 3799, '92-fa2f7089-cb5a-4770-a387-8f5b4cc72b6e.png', 'Самый лучший среди товаров на рынке Стол компьютерный с ящиками белый Frenesie, ЛДСП, 90х45х77 см.', 4.95, 92),
	('0c7c7c4b-9dd1-4c05-ae5f-c0d8e043176e', 'Стол игровой Defender Space с подставкой под кружку и подвесом под гарнитуру, чёрный', 8484, '92-0c7c7c4b-9dd1-4c05-ae5f-c0d8e043176e.jpg', 'Самый лучший среди товаров на рынке Стол игровой Defender Space с подставкой под кружку и подвесом под гарнитуру, чёрный', 4.24, 92),
	('caaaf490-0b02-4d72-9842-71e27652be35', 'Стол компьютерный игровой LevelUP 1400 Черный, 140*74 см.', 7200, '92-caaaf490-0b02-4d72-9842-71e27652be35.jpg', 'Самый лучший среди товаров на рынке Стол компьютерный игровой LevelUP 1400 Черный, 140*74 см.', 4.33, 92),
	('5287eb68-db0d-4cf1-b184-ef2b9e79c823', 'Компьютерный, письменный стол СТОЛПЛИТ Альфа СБ-2955', 5280, '92-5287eb68-db0d-4cf1-b184-ef2b9e79c823.jpeg', 'Самый лучший среди товаров на рынке Компьютерный, письменный стол СТОЛПЛИТ Альфа СБ-2955', 4.2, 92),
	('b0cc9a20-6818-483e-aa4c-11c92f147c76', 'Стол письменный Домотека Мартин БЛ 71 БЛ белый 110х57х75', 6900, '92-b0cc9a20-6818-483e-aa4c-11c92f147c76.jpeg', 'Самый лучший среди товаров на рынке Стол письменный Домотека Мартин БЛ 71 БЛ белый 110х57х75', 4.29, 92),
	('894af921-2bd1-4fe8-8992-c3c15faeab4e', 'Стол письменный ТОП-мебель Фрея 4 ящика, 110х75х50см дуб сонома', 4650, '92-894af921-2bd1-4fe8-8992-c3c15faeab4e.jpeg', 'Самый лучший среди товаров на рынке Стол письменный ТОП-мебель Фрея 4 ящика, 110х75х50см дуб сонома', 4.6, 92),
	('43d7584a-3c16-4d2f-87b9-df5e75a27ae5', 'Стол компьютерный, стол письменный Ascetic 1200 Черный/Красный, 120*71,6 см.', 4790, '92-43d7584a-3c16-4d2f-87b9-df5e75a27ae5.jpg', 'Самый лучший среди товаров на рынке Стол компьютерный, стол письменный Ascetic 1200 Черный/Красный, 120*71,6 см.', 4.36, 92),
	('94ac6700-a707-451a-8e1d-da9a0e75820a', 'Стол письменный Домотека Мартин ДК 71 ЧР темно бежевый/черный 110х57х75', 6900, '92-94ac6700-a707-451a-8e1d-da9a0e75820a.jpeg', 'Самый лучший среди товаров на рынке Стол письменный Домотека Мартин ДК 71 ЧР темно бежевый/черный 110х57х75', 4.40, 92),
	('2a053227-bce9-4fcd-9f49-566633c9fee1', 'Компьютерный, письменный стол СТОЛПЛИТ Альфа СБ-2955 Диамант Серый', 5580, '92-2a053227-bce9-4fcd-9f49-566633c9fee1.jpeg', 'Самый лучший среди товаров на рынке Компьютерный, письменный стол СТОЛПЛИТ Альфа СБ-2955 Диамант Серый', 4.8, 92),
	('7ebb673a-1e29-4386-a091-be39ac31d58e', 'Стол письменный, компьютерный с ящиками Марио-1 БЛ 71 М БЛ-2 (120х60х75) белый', 9900, '92-7ebb673a-1e29-4386-a091-be39ac31d58e.jpeg', 'Самый лучший среди товаров на рынке Стол письменный, компьютерный с ящиками Марио-1 БЛ 71 М БЛ-2 (120х60х75) белый', 4.29, 92),
	('8bed39b9-4e2f-427e-bd8b-bcdcc683a0a1', 'Стол письменный ТриЯ тип 4 Дуб Сонома/Белый Ясень', 5999, '92-8bed39b9-4e2f-427e-bd8b-bcdcc683a0a1.jpg', 'Самый лучший среди товаров на рынке Стол письменный ТриЯ тип 4 Дуб Сонома/Белый Ясень', 4.24, 92),
	('b1828ac1-0c13-4c7c-aabc-6b2051f7bd3d', 'Стол компьютерный ZONE 51 PLATFORM Ambilight 120', 22840, '92-b1828ac1-0c13-4c7c-aabc-6b2051f7bd3d.png', 'Самый лучший среди товаров на рынке Стол компьютерный ZONE 51 PLATFORM Ambilight 120', 4.10, 92),
	('477bd96f-c0c6-4bc4-9672-e962aa04c32c', 'Стол игровой Defender Jupiter RGB с подставкой под кружку и подвесом под гарнитуру, чёрный', 9563, '92-477bd96f-c0c6-4bc4-9672-e962aa04c32c.jpg', 'Самый лучший среди товаров на рынке Стол игровой Defender Jupiter RGB с подставкой под кружку и подвесом под гарнитуру, чёрный', 4.9, 92),
	('366b4ece-177b-4966-bfbf-26d09ccb1d25', 'Стол письменный, компьютерный с ящиками Марио БЛ 71 М БЛ-2 (110х55х75) белый', 9400, '92-366b4ece-177b-4966-bfbf-26d09ccb1d25.jpeg', 'Самый лучший среди товаров на рынке Стол письменный, компьютерный с ящиками Марио БЛ 71 М БЛ-2 (110х55х75) белый', 4.29, 92),
	('795ecbd8-2b1e-4375-92b4-51417ccccedf', 'Стол компьютерный игровой LevelUP 1400 Белый/Серый, 140*74 см.', 7490, '92-795ecbd8-2b1e-4375-92b4-51417ccccedf.jpg', 'Самый лучший среди товаров на рынке Стол компьютерный игровой LevelUP 1400 Белый/Серый, 140*74 см.', 4.21, 92),
	('35424fee-5f03-4bc9-9af3-62f47d4bdb40', 'Стол компьютерный игровой LevelUP 1400 Дуб Сонома, 140*74 см.', 7490, '92-35424fee-5f03-4bc9-9af3-62f47d4bdb40.jpg', 'Самый лучший среди товаров на рынке Стол компьютерный игровой LevelUP 1400 Дуб Сонома, 140*74 см.', 4.7, 92),
	('ccbc22d2-ffc5-402b-8f4a-733aee2c0174', 'Стол компьютерный ZONE 51 PLATFORM Ambilight 120', 19690, '92-ccbc22d2-ffc5-402b-8f4a-733aee2c0174.png', 'Самый лучший среди товаров на рынке Стол компьютерный ZONE 51 PLATFORM Ambilight 120', 4.10, 92),
	('44e69281-d026-48bc-bebe-399f0a586dff', 'Стол письменный ТриЯ тип 4 Дуб Сонома/Белый Ясень', 8089, '92-44e69281-d026-48bc-bebe-399f0a586dff.jpg', 'Самый лучший среди товаров на рынке Стол письменный ТриЯ тип 4 Дуб Сонома/Белый Ясень', 4.24, 92),
	('ffa4fc42-d491-45a9-a8f2-54a1056ac554', 'Маленький компьютерный стол Олмеко "Костер-4", белый', 2218, '92-ffa4fc42-d491-45a9-a8f2-54a1056ac554.jpeg', 'Самый лучший среди товаров на рынке Маленький компьютерный стол Олмеко "Костер-4", белый', 4.33, 92),
	('1086bb26-df54-468c-8be1-b5ba4ecf8fcd', 'Атмосфера Найс стол письменный исп 2 00-00017838', 12024, '92-1086bb26-df54-468c-8be1-b5ba4ecf8fcd.jpg', 'Самый лучший среди товаров на рынке Атмосфера Найс стол письменный исп 2 00-00017838', 4.13, 92),
	('71437f76-6ea6-484a-80a6-502fabaa535e', 'Стол компьютерный ZONE 51 PLATFORM Ambilight 120', 18046, '92-71437f76-6ea6-484a-80a6-502fabaa535e.png', 'Самый лучший среди товаров на рынке Стол компьютерный ZONE 51 PLATFORM Ambilight 120', 4.10, 92),
	('f120c842-7df6-448e-80b4-59162c7ec1b2', 'Стол письменный, компьютерный Домотека Нобель 1 (120х60х75) СБ 71 ЧР, серый, черный', 5300, '92-f120c842-7df6-448e-80b4-59162c7ec1b2.jpeg', 'Самый лучший среди товаров на рынке Стол письменный, компьютерный Домотека Нобель 1 (120х60х75) СБ 71 ЧР, серый, черный', 4.12, 92),
	('26febc34-f64c-46a3-922c-862e137cd1d7', 'Стол письменный прямой Skyland Imago/Белый, 1600х720х755, СП-4', 5878, '92-26febc34-f64c-46a3-922c-862e137cd1d7.jpg', 'Самый лучший среди товаров на рынке Стол письменный прямой Skyland Imago/Белый, 1600х720х755, СП-4', 4.16, 92),
	('e930f487-ad9b-4c39-af4b-44d55d21b8b6', 'Компьютерный стол LOFTWELL Modern Plus Country', 16000, '92-e930f487-ad9b-4c39-af4b-44d55d21b8b6.jpeg', 'Самый лучший среди товаров на рынке Компьютерный стол LOFTWELL Modern Plus Country', 4.6, 92),
	('a052c2df-39b6-4168-8142-b8ee10f73f92', 'Стол компьютерный игровой LevelUP 1400 Черный/Красный, 140*74 см.', 7490, '92-a052c2df-39b6-4168-8142-b8ee10f73f92.jpg', 'Самый лучший среди товаров на рынке Стол компьютерный игровой LevelUP 1400 Черный/Красный, 140*74 см.', 4.26, 92),
	('0e565034-e3c3-4785-8824-db14243b71cf', 'Атмосфера Найс стол письменный исп 2 00-00017838', 4930, '92-0e565034-e3c3-4785-8824-db14243b71cf.jpg', 'Самый лучший среди товаров на рынке Атмосфера Найс стол письменный исп 2 00-00017838', 4.13, 92),
	('2e5148f8-dfea-474c-819f-248d2c4cef67', 'Атмосфера Найс стол письменный исп 2 00-00017838', 12024, '92-2e5148f8-dfea-474c-819f-248d2c4cef67.jpg', 'Самый лучший среди товаров на рынке Атмосфера Найс стол письменный исп 2 00-00017838', 4.13, 92),
	('6b9edd84-d614-4c49-baba-dc03e304e7ad', 'Стол компьютерный ZONE 51 PLATFORM Ambilight 120', 17990, '92-6b9edd84-d614-4c49-baba-dc03e304e7ad.png', 'Самый лучший среди товаров на рынке Стол компьютерный ZONE 51 PLATFORM Ambilight 120', 4.10, 92),
	('0d943410-42d4-4ec6-965a-e04ccc37353f', 'Стол компьютерный игровой LevelUP 1400 Венге, 140*74 см.', 7490, '92-0d943410-42d4-4ec6-965a-e04ccc37353f.jpg', 'Самый лучший среди товаров на рынке Стол компьютерный игровой LevelUP 1400 Венге, 140*74 см.', 4.11, 92),
	('59ae93f5-4598-474d-9d47-e0024a1ea1f1', 'Письменный стол, компьютерный стол, офисный стол FLAT "Пайн" 120х60х75 см', 10500, '92-59ae93f5-4598-474d-9d47-e0024a1ea1f1.jpeg', 'Самый лучший среди товаров на рынке Письменный стол, компьютерный стол, офисный стол FLAT "Пайн" 120х60х75 см', 4.14, 92),
	('c3de189b-2ba1-40b2-9331-e503fa8c0930', 'Стол письменный ТриЯ тип 4 Дуб Крафт Белый', 6299, '92-c3de189b-2ba1-40b2-9331-e503fa8c0930.jpg', 'Самый лучший среди товаров на рынке Стол письменный ТриЯ тип 4 Дуб Крафт Белый', 4.19, 92),
	('cc46f69b-6ae4-4738-8c62-15fb3d45c610', 'Компьютерный стол, письменный стол LOFTWELL в стиле ЛОФТ Modern Plus Country', 16000, '92-cc46f69b-6ae4-4738-8c62-15fb3d45c610.jpeg', 'Самый лучший среди товаров на рынке Компьютерный стол, письменный стол LOFTWELL в стиле ЛОФТ Modern Plus Country', 4.1, 92),
	('17037012-557f-42d0-954e-dc36af7c5d10', 'Стол письменный Комфорт-S АГАТА М16 белый', 4390, '92-17037012-557f-42d0-954e-dc36af7c5d10.jpeg', 'Самый лучший среди товаров на рынке Стол письменный Комфорт-S АГАТА М16 белый', 4.1, 92),
	('391998e4-e27d-4872-970f-0bb659ee4f34', 'Стол компьютерный, стол письменный Ascetic 1200 Венге, 120*71,6 см.', 4790, '92-391998e4-e27d-4872-970f-0bb659ee4f34.jpg', 'Самый лучший среди товаров на рынке Стол компьютерный, стол письменный Ascetic 1200 Венге, 120*71,6 см.', 4.23, 92),
	('b623c7ff-a7e5-4df4-a535-a6d086400dce', 'Геймерский стол компьютерный, стол письменный Jedi 1400 Венге, 140*71,6 см.', 6990, '92-b623c7ff-a7e5-4df4-a535-a6d086400dce.jpg', 'Самый лучший среди товаров на рынке Геймерский стол компьютерный, стол письменный Jedi 1400 Венге, 140*71,6 см.', 4.6, 92),
	('9159539d-2ee0-4abb-aa67-473dd7ef2d3f', 'Стол Cactus CS-EGD-BBK, стекло, черный', 60195, '92-9159539d-2ee0-4abb-aa67-473dd7ef2d3f.jpg', 'Самый лучший среди товаров на рынке Стол Cactus CS-EGD-BBK, стекло, черный', 4.9, 92),
	('84be5724-c616-4f52-80d5-688022bdec07', 'Атмосфера Найс стол письменный исп 2 00-00017838', 6029, '92-84be5724-c616-4f52-80d5-688022bdec07.jpg', 'Самый лучший среди товаров на рынке Атмосфера Найс стол письменный исп 2 00-00017838', 4.13, 92),
	('9f18a466-0e2d-4d06-a8c3-96ddc09e9262', 'Стол компьютерный игровой LevelUP 1400 Венге/Белый, 140*74 см.', 7490, '92-9f18a466-0e2d-4d06-a8c3-96ddc09e9262.jpg', 'Самый лучший среди товаров на рынке Стол компьютерный игровой LevelUP 1400 Венге/Белый, 140*74 см.', 4.22, 92),
	('6eda5c21-ef9d-4ec6-883b-6740716af7a5', 'Стол компьютерный ZONE 51 PLATFORM PRO 120, черный', 37870, '92-6eda5c21-ef9d-4ec6-883b-6740716af7a5.jpeg', 'Самый лучший среди товаров на рынке Стол компьютерный ZONE 51 PLATFORM PRO 120, черный', 4.2, 92),
	('006ba728-b33c-4676-83eb-956fad463aff', 'Письменный стол, Компьютерный стол LOFTWELL India, 112х64х75 см', 14000, '92-006ba728-b33c-4676-83eb-956fad463aff.jpeg', 'Самый лучший среди товаров на рынке Письменный стол, Компьютерный стол LOFTWELL India, 112х64х75 см', 4.22, 92),
	('1a12a8b2-1f17-42ba-b308-9ff0e2f55b1a', 'Компьютерный стол, письменный стол LOFTWELL в стиле ЛОФТ Modern Plus Country', 16000, '92-1a12a8b2-1f17-42ba-b308-9ff0e2f55b1a.jpeg', 'Самый лучший среди товаров на рынке Компьютерный стол, письменный стол LOFTWELL в стиле ЛОФТ Modern Plus Country', 4.3, 92),
	('d6f8077b-656e-43c1-a41f-2cdef0cf6b25', 'Компьютерный стол, письменный стол LOFTWELL в стиле ЛОФТ Modern Plus India', 16000, '92-d6f8077b-656e-43c1-a41f-2cdef0cf6b25.jpeg', 'Самый лучший среди товаров на рынке Компьютерный стол, письменный стол LOFTWELL в стиле ЛОФТ Modern Plus India', 4.11, 92);

INSERT INTO product (id, name, price, imgsrc, description, rating, category_id)
	VALUES

    ('9d82fa8a-62b1-4800-b886-cd0b7ac0031c', 'Комплект стульев для кухни Ridberg Лори Velour grey 2 шт', 8090, '91-9d82fa8a-62b1-4800-b886-cd0b7ac0031c.jpeg', 'Самый лучший среди товаров на рынке Комплект стульев для кухни Ridberg Лори Velour grey 2 шт', 4.30, 91),
    ('efd17547-164d-450c-b687-b10b13afdbaf', 'Стул для кухни обеденный DSW Style белый (комплект 4 стула)', 7290, '91-efd17547-164d-450c-b687-b10b13afdbaf.jpg', 'Самый лучший среди товаров на рынке Стул для кухни обеденный DSW Style белый (комплект 4 стула)', 4.13, 91),
    ('f0be1355-f220-4938-b0d6-4d34d89f2e4b', 'Комплект складных стульев Stool Group SUPER LITE N банкетный 4 шт белый', 9840, '91-f0be1355-f220-4938-b0d6-4d34d89f2e4b.jpg', 'Самый лучший среди товаров на рынке Комплект складных стульев Stool Group SUPER LITE N банкетный 4 шт белый', 4.1, 91),
    ('271a010a-6fd1-4c42-8892-1f29dc4d7699', 'Стул для кухни обеденный Одди велюр светло-серый (комплект 4 стула)', 18990, '91-271a010a-6fd1-4c42-8892-1f29dc4d7699.jpg', 'Самый лучший среди товаров на рынке Стул для кухни обеденный Одди велюр светло-серый (комплект 4 стула)', 4.5, 91),
    ('eeafd5df-d53a-4fc8-9ecc-a5b98cc34037', 'Стул TetChair CHILLY (mod. 7095-1) ткань, металл, серый/черный', 8448, '91-eeafd5df-d53a-4fc8-9ecc-a5b98cc34037.jpeg', 'Самый лучший среди товаров на рынке Стул TetChair CHILLY (mod. 7095-1) ткань, металл, серый/черный', 4.19, 91),
    ('1f982c8c-31ae-4099-8956-3ffe5af62450', 'Стул byROOM Gokotta A276-3-E, изумрудный', 3625, '91-1f982c8c-31ae-4099-8956-3ffe5af62450.jpeg', 'Самый лучший среди товаров на рынке Стул byROOM Gokotta A276-3-E, изумрудный', 4.23, 91),
    ('fa78e50f-f2ac-4bca-82de-73f2ca60dc8b', 'Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 1564, '91-fa78e50f-f2ac-4bca-82de-73f2ca60dc8b.jpg', 'Самый лучший среди товаров на рынке Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 4.70, 91),
    ('05e71eff-aad5-495c-b46c-6e824efe7ec4', 'Комплект стульев 2 шт. LEON GROUP для кухни в стиле EAMES DSW, белый', 3900, '91-05e71eff-aad5-495c-b46c-6e824efe7ec4.jpeg', 'Самый лучший среди товаров на рынке Комплект стульев 2 шт. LEON GROUP для кухни в стиле EAMES DSW, белый', 4.23, 91),
    ('e9da00a1-665b-40c6-b5fa-67b113031c71', 'Стул Sephi, черный', 2699, '91-e9da00a1-665b-40c6-b5fa-67b113031c71.jpg', 'Самый лучший среди товаров на рынке Стул Sephi, черный', 4.8, 91),
    ('bc0fb5ca-32db-49dd-9031-e163ac216c63', 'Стул TetChair CHILLY (mod. 7095-1) ткань, металл, серый/черный', 8448, '91-bc0fb5ca-32db-49dd-9031-e163ac216c63.jpeg', 'Самый лучший среди товаров на рынке Стул TetChair CHILLY (mod. 7095-1) ткань, металл, серый/черный', 4.19, 91),
    ('80c78d3a-41a7-421b-b85f-f9ccc8895aa7', 'Комплект стульев 4 шт. LEON GROUP для кухни в стиле EAMES DSW, белый', 6450, '91-80c78d3a-41a7-421b-b85f-f9ccc8895aa7.jpeg', 'Самый лучший среди товаров на рынке Комплект стульев 4 шт. LEON GROUP для кухни в стиле EAMES DSW, белый', 4.35, 91),
    ('09875361-3a37-49f4-8010-743aae29a1f4', 'Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 918, '91-09875361-3a37-49f4-8010-743aae29a1f4.jpg', 'Самый лучший среди товаров на рынке Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 4.70, 91),
    ('6cb9c91a-2b20-48b7-b8ce-d284039d1562', 'Стул TetChair CHILLY (mod. 7095-1) ткань, металл, темно-серый', 7611, '91-6cb9c91a-2b20-48b7-b8ce-d284039d1562.jpeg', 'Самый лучший среди товаров на рынке Стул TetChair CHILLY (mod. 7095-1) ткань, металл, темно-серый', 4.8, 91),
    ('ac769c80-b5bf-4733-a313-00d74a97ae69', 'Стул складной "Ника 1", цвет слоновая кость', 3188, '91-ac769c80-b5bf-4733-a313-00d74a97ae69.jpg', 'Самый лучший среди товаров на рынке Стул складной "Ника 1", цвет слоновая кость', 4.18, 91),
    ('2cb81d7b-9409-42b9-b292-a7418fed026c', 'Стул TetChair CHILLY (mod. 7095-1) ткань, металл, серый/черный', 8448, '91-2cb81d7b-9409-42b9-b292-a7418fed026c.jpeg', 'Самый лучший среди товаров на рынке Стул TetChair CHILLY (mod. 7095-1) ткань, металл, серый/черный', 4.19, 91),
    ('ad6be4bc-689e-4640-98b4-6c4107ac2351', 'Стул для кухни обеденный DSW Style темно-серый (комплект 4 стула)', 10190, '91-ad6be4bc-689e-4640-98b4-6c4107ac2351.jpg', 'Самый лучший среди товаров на рынке Стул для кухни обеденный DSW Style темно-серый (комплект 4 стула)', 4.29, 91),
    ('b861b238-2fdc-4870-b4d0-7f0eedd29f39', 'Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 1087, '91-b861b238-2fdc-4870-b4d0-7f0eedd29f39.jpg', 'Самый лучший среди товаров на рынке Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 4.70, 91),
    ('e024f9f0-56e8-47ad-85e7-b56c2a95d65a', 'Стул TetChair CHILLY (mod. 7095-1) ткань, металл, темно-серый', 7611, '91-e024f9f0-56e8-47ad-85e7-b56c2a95d65a.jpeg', 'Самый лучший среди товаров на рынке Стул TetChair CHILLY (mod. 7095-1) ткань, металл, темно-серый', 4.8, 91),
    ('bba02d14-39f8-47c1-8bf4-22b4aae7282d', 'Стул Ridberg БРУКЛИН (Light Grey)', 6390, '91-bba02d14-39f8-47c1-8bf4-22b4aae7282d.jpeg', 'Самый лучший среди товаров на рынке Стул Ridberg БРУКЛИН (Light Grey)', 4.12, 91),
    ('da113290-3b7e-411a-a3fa-e9db1d28c50d', 'Стул складной "Ника 1", цвет слоновая кость', 1875, '91-da113290-3b7e-411a-a3fa-e9db1d28c50d.jpg', 'Самый лучший среди товаров на рынке Стул складной "Ника 1", цвет слоновая кость', 4.18, 91),
    ('5231e9d4-e152-42c3-a434-6b36c94c913c', 'Стул Woodville 15110, gray/black', 5888, '91-5231e9d4-e152-42c3-a434-6b36c94c913c.jpg', 'Самый лучший среди товаров на рынке Стул Woodville 15110, gray/black', 4.1, 91),
    ('aedbd529-bc09-4924-95c2-3ea6dc0fb6c0', 'Стул TetChair CHILLY (mod. 7095-1) ткань, металл, серый/черный', 9504, '91-aedbd529-bc09-4924-95c2-3ea6dc0fb6c0.jpeg', 'Самый лучший среди товаров на рынке Стул TetChair CHILLY (mod. 7095-1) ткань, металл, серый/черный', 4.19, 91),
    ('8fc1d781-a348-40a2-9a01-b9c39641d93b', 'Стул Экспресс офис Венский CH 56, молочный', 1505, '91-8fc1d781-a348-40a2-9a01-b9c39641d93b.jpg', 'Самый лучший среди товаров на рынке Стул Экспресс офис Венский CH 56, молочный', 4.6, 91),
    ('d8beb2dd-f0e8-4086-b258-498ae24f44c1', 'Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 1380, '91-d8beb2dd-f0e8-4086-b258-498ae24f44c1.jpg', 'Самый лучший среди товаров на рынке Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 4.70, 91),
    ('248a2e2e-08bc-4a12-96cb-a45a11f137d7', 'Стул Woodville Gabi 1 Dark beige/Black', 5888, '91-248a2e2e-08bc-4a12-96cb-a45a11f137d7.jpg', 'Самый лучший среди товаров на рынке Стул Woodville Gabi 1 Dark beige/Black', 4.10, 91),
    ('3d5844b1-d667-4965-9a6e-2aa1357d8636', 'Деревянный стул Гольф черный/ массив березы М-трейд Гольф Z-09 ИнгZ-09 черн', 4360, '91-3d5844b1-d667-4965-9a6e-2aa1357d8636.jpg', 'Самый лучший среди товаров на рынке Деревянный стул Гольф черный/ массив березы М-трейд Гольф Z-09 ИнгZ-09 черн', 4.3, 91),
    ('b7d08a5e-7f41-4568-bc9a-1e56150e54d3', 'Стул TetChair CHILLY (mod. 7095-1) ткань, металл, темно-серый', 8562, '91-b7d08a5e-7f41-4568-bc9a-1e56150e54d3.jpeg', 'Самый лучший среди товаров на рынке Стул TetChair CHILLY (mod. 7095-1) ткань, металл, темно-серый', 4.8, 91),
    ('dedef5fc-e7dd-4f4a-8f6b-30261b8eea82', 'Стул складной "Ника 1", цвет слоновая кость', 1712, '91-dedef5fc-e7dd-4f4a-8f6b-30261b8eea82.jpg', 'Самый лучший среди товаров на рынке Стул складной "Ника 1", цвет слоновая кость', 4.18, 91),
    ('5a8c5098-5889-412b-af0c-a3df9d649144', 'Табурет Nika ТЭ2/С 32х32х46,5 см, серый', 1175, '91-5a8c5098-5889-412b-af0c-a3df9d649144.jpg', 'Самый лучший среди товаров на рынке Табурет Nika ТЭ2/С 32х32х46,5 см, серый', 4.66, 91),
    ('62a8dff9-4c26-4412-a04e-893ae7866886', 'Комплект стульев для кухни Ridberg Лори Velour grey 2 шт', 8090, '91-62a8dff9-4c26-4412-a04e-893ae7866886.jpeg', 'Самый лучший среди товаров на рынке Комплект стульев для кухни Ridberg Лори Velour grey 2 шт', 4.30, 91),
    ('e38b50e6-4730-4f42-b0f6-b47268d37be4', 'Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 681, '91-e38b50e6-4730-4f42-b0f6-b47268d37be4.jpg', 'Самый лучший среди товаров на рынке Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 4.70, 91),
    ('73fea4ff-7259-4ebf-8bfa-b8f6e80ad651', 'Стул Woodville Gabi 1 Dark beige/Black', 4995, '91-73fea4ff-7259-4ebf-8bfa-b8f6e80ad651.jpg', 'Самый лучший среди товаров на рынке Стул Woodville Gabi 1 Dark beige/Black', 4.10, 91),
    ('a4969a16-73f4-4771-ab42-6927a73b5cc4', 'Комплект стульев 2 шт. RIDBERG Лондон RLOBE2, beige', 9459, '91-a4969a16-73f4-4771-ab42-6927a73b5cc4.jpeg', 'Самый лучший среди товаров на рынке Комплект стульев 2 шт. RIDBERG Лондон RLOBE2, beige', 4.69, 91),
    ('1378b6d7-d73c-43c7-a2f9-7bbd09e66a32', 'Табурет Nika тб1 Эконом', 459, '91-1378b6d7-d73c-43c7-a2f9-7bbd09e66a32.jpg', 'Самый лучший среди товаров на рынке Табурет Nika тб1 Эконом', 4.21, 91),
    ('6cee7a3a-47db-46d7-9b44-d7bc6ab97b21', 'Стул кухонный Ridberg CIRCUS (Beige)', 4190, '91-6cee7a3a-47db-46d7-9b44-d7bc6ab97b21.jpeg', 'Самый лучший среди товаров на рынке Стул кухонный Ridberg CIRCUS (Beige)', 4.10, 91),
    ('0dccc97b-2e8f-49c0-bd64-923226273c39', 'Табурет Nika ТЭ2/С 32х32х46,5 см, серый', 1170, '91-0dccc97b-2e8f-49c0-bd64-923226273c39.jpg', 'Самый лучший среди товаров на рынке Табурет Nika ТЭ2/С 32х32х46,5 см, серый', 4.66, 91),
    ('3774d1c9-07af-4725-a989-bf774f498fae', 'Табурет Nika Фабрик 1, серый', 1206, '91-3774d1c9-07af-4725-a989-bf774f498fae.jpg', 'Самый лучший среди товаров на рынке Табурет Nika Фабрик 1, серый', 4.48, 91),
    ('270428bb-57aa-49fd-9d08-09b412c3c8f8', 'Табурет кухонный Ника Эконом слоновая кость', 1039, '91-270428bb-57aa-49fd-9d08-09b412c3c8f8.jpg', 'Самый лучший среди товаров на рынке Табурет кухонный Ника Эконом слоновая кость', 4.6, 91),
    ('f6fdf188-38d1-422c-801f-7fdc01a3c757', 'Стул для кухни обеденный DSW Style коричневый (Комплект 4 стула)', 10190, '91-f6fdf188-38d1-422c-801f-7fdc01a3c757.jpeg', 'Самый лучший среди товаров на рынке Стул для кухни обеденный DSW Style коричневый (Комплект 4 стула)', 4.13, 91),
    ('375ddd90-0d55-43cd-af79-0dfb53a0711f', 'Табурет Nika Фабрик 2, с мягким сиденьем, 32 см, серый', 1492, '91-375ddd90-0d55-43cd-af79-0dfb53a0711f.jpeg', 'Самый лучший среди товаров на рынке Табурет Nika Фабрик 2, с мягким сиденьем, 32 см, серый', 4.32, 91),
    ('79870c37-4658-4b7c-bf3a-92e1f9155b00', 'Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 768, '91-79870c37-4658-4b7c-bf3a-92e1f9155b00.jpg', 'Самый лучший среди товаров на рынке Табурет Nika ТЭ2/СК 32х32х46,5 см, слоновая кость/серый', 4.70, 91),
    ('cae0584e-8817-466f-b498-9b74812c757d', 'Стул Woodville Gabi 1 Dark beige/Black', 5490, '91-cae0584e-8817-466f-b498-9b74812c757d.jpg', 'Самый лучший среди товаров на рынке Стул Woodville Gabi 1 Dark beige/Black', 4.10, 91),
    ('65f4cf4a-9af9-4e74-85f9-48471914a83e', 'Комплект стульев 2 шт. RIDBERG Лондон RLOBE2, beige', 8090, '91-65f4cf4a-9af9-4e74-85f9-48471914a83e.jpeg', 'Самый лучший среди товаров на рынке Комплект стульев 2 шт. RIDBERG Лондон RLOBE2, beige', 4.69, 91),
    ('98f12385-1c0f-4f1e-9fa7-c1b2fe479724', 'Комплект мягких стульев (4 шт.) М-ТРЕЙД Ruчu-сер-04Ч Ruчu-сер- 4чер', 16470, '91-98f12385-1c0f-4f1e-9fa7-c1b2fe479724.jpeg', 'Самый лучший среди товаров на рынке Комплект мягких стульев (4 шт.) М-ТРЕЙД Ruчu-сер-04Ч Ruчu-сер- 4чер', 4.12, 91),
    ('257a44e0-4f18-41f2-ade7-f2d322668cca', 'Табурет Nika  пластмассовое сидение черный ТП01 1', 648, '91-257a44e0-4f18-41f2-ade7-f2d322668cca.jpeg', 'Самый лучший среди товаров на рынке Табурет Nika  пластмассовое сидение черный ТП01 1', 4.9, 91),
    ('7a6d4b41-3288-4e67-95ca-3f1c6a0d1031', 'Стул Ridberg DAKLINE (Grey)', 5190, '91-7a6d4b41-3288-4e67-95ca-3f1c6a0d1031.jpeg', 'Самый лучший среди товаров на рынке Стул Ridberg DAKLINE (Grey)', 4.3, 91),
    ('7668905c-5397-420d-abbd-727ebb2dc3f4', 'Табурет Nika Фабрик 1, серый', 953, '91-7668905c-5397-420d-abbd-727ebb2dc3f4.jpg', 'Самый лучший среди товаров на рынке Табурет Nika Фабрик 1, серый', 4.48, 91),
    ('4fed4431-1294-4f33-a52c-60a274429969', 'Табурет кухонный Ника Эконом слоновая кость', 908, '91-4fed4431-1294-4f33-a52c-60a274429969.jpg', 'Самый лучший среди товаров на рынке Табурет кухонный Ника Эконом слоновая кость', 4.6, 91);
