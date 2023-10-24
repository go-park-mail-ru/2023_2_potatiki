# Фунциональные зависимости
- Таблица profiles
   - {id} -> login, description, imgsrc, passwordhash
   - {login} -> id, description, imgsrc, passwordhash

- Таблица products
    - {id} -> name, description, price, imgsrc, rating, categoryId
    - {name} -> id, description, price, imgsrc, rating, categoryId

- Таблица orders
    - {id} ->  profileId, promocodeId, status, deliveryDate, creationDate

- Таблица order_items
    - {id} -> quantity, orderId, productId
    - {orderId, productId} -> id, quantity

- Таблица favorites
    - {id} -> profileId, productId
    - {profileId, productId} -> id

- Таблица addresses
    - {id} -> profileId, city, street, house, flat, isCurrent
    - {profileId, city, street, house, flat} -> id, isCurrent

- Таблица categories
    - {id} -> name
    - {name} -> id

- Таблица category_refferences
    - {categoryId} -> parentCategoryId

- Таблица shopping_cart_items
    - {id} ->  quantity, profileId, productId
    - {profileId, productId} - id, quantity

- Таблица promocodes
    - {id} -> name, discount
    - {name} -> id, discount
```mermaid
erDiagram
    PROFILES ||--o{ ORDERS : includes
    PROFILES ||--o{ ADDRESSES : includes
    PROFILES ||--o{ SHOPPING_CART_ITEMS : includes
    PROFILES ||--o{ FAVORITES : includes
    
    PROMOCODES ||--o{ ORDERS : includes
    ORDERS ||--o{ ORDER_ITEMS : includes

    PRODUCTS ||--o{ ORDER_ITEMS : includes
    PRODUCTS ||--o{ SHOPPING_CART_ITEMS : includes
    CATEGORIES ||--|{ PRODUCTS : includes
    PRODUCTS ||--o{ FAVORITES : includes

    CATEGORIES ||--|{ CATEGORY_REFERENCES : includes

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

    FAVORITES {
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

    ORDER_ITEMS {
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

    CATEGORY_REFERENCES {
        uuid categoryId PK, FK
        uuid parentCategoryId FK
    }

    SHOPPING_CART_ITEMS {
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
