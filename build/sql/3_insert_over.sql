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
