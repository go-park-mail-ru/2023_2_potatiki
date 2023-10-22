# Фунциональные зависимости
- Таблица profiles
   - {id} -> login, description, imgsrc, passwordhash

- Таблица products
    - {id} -> name, description, price, imgsrc, rating

- Таблица orders
    - {id} ->  profileId, promocodeId

- Таблица orderitems
    - {id, orderId, productId} -> quantity

- Таблица favourites
    - {id} -> profileId, productId 

- Таблица addresses
    - {id} -> profileId, city, street, house, flat, isCurrent

- Таблица categories
    - {id} -> name

- Таблица categoryrefferences
    - {categoryId} -> parentCategoryId

- Таблица shoppingcartitems
    - {id, profileId, productId} ->  quantity

- Таблица promocodes
    - {id} -> name

```mermaid
erDiagram
    PROFILES ||--o{ ORDERS : includes
    PROFILES ||--o{ ADDRESSES : includes
    PROFILES ||--o{ SHOPPINGCARTITEMS : includes
    PROFILES ||--o{ FAVOURITES : includes
    
    ORDERS ||--o{ PROMOCODES : includes
    ORDERS ||--o{ ORDERITEMS : includes

    PRODUCTS ||--o{ ORDERITEMS : includes
    PRODUCTS ||--o{ SHOPPINGCARTITEMS : includes
    PRODUCTS ||--|{ CATEGORIES : includes
    PRODUCTS ||--o{ FAVOURITES : includes

    CATEGORIES ||--|{ CATEGORYREFERENCES : includes

    PROFILES {
        uuid id PK
        text login UK
        text description
        text imgsrc
        text passwordhash
    }
    
    PRODUCTS {
        uuid id PK
        text name UK
        text description
        int price
        text imgsrc
        number rating
        uuid category FK
    }

    FAVOURITES {
        uuid id PK
        uuid profileId FK
        uuid productId FK
    }

    ORDERS {
        uuid id PK
        uuid profileId FK
        uuid promocodeId FK
        text status
        timestamp creationDate
        timestamp deliveryDate
    }

    ORDERITEMS {
        uuid id PK
        uuid orderId FK
        uuid productId FK
        int quantity
        int price
    }

    ADDRESSES {
        uuid id PK
        uuid profileId FK
        text city
        text street
        text house
        text flat
        bool isCurrent
    }

    CATEGORIES {
        uuid id PK
        text name UK
    }

    CATEGORYREFERENCES {
        uuid categoryId FK
        uuid parentCategoryId FK
    }

    SHOPPINGCARTITEMS {
        uuid id PK
        uuid profileId FK
        uuid productId FK
        int quantity
    }

    PROMOCODES {
        uuid id PK
        text name UK
        int discount
    }

```
