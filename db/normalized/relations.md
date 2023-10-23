# Фунциональные зависимости
- Таблица profile
   - {id} -> login, description, imgsrc, passwordhash
   - {login} -> id, description, imgsrc, passwordhash

- Таблица product
    - {id} -> name, description, price, imgsrc, rating, categoryId
    - {name} -> id, description, price, imgsrc, rating, categoryId

- Таблица order_info
    - {id} ->  profileId, promocodeId, status, deliveryDate, creationDate

- Таблица order_item
    - {id} -> quantity, orderId, productId
    - {orderId, productId} -> id, quantity

- Таблица favorite
    - {id} -> profileId, productId
    - {profileId, productId} -> id

- Таблица address
    - {id} -> profileId, city, street, house, flat, isCurrent
    - {profileId, city, street, house, flat} -> id, isCurrent

- Таблица category
    - {id} -> name
    - {name} -> id

- Таблица category_refference
    - {id} -> categoryId, parentCategoryId
    - {categoryId} -> id, parentCategoryId

- Таблица shopping_cart_item
    - {id} ->  quantity, profileId, productId
    - {profileId, productId} - id, quantity

- Таблица promocode
    - {id} -> name, discount
    - {name} -> id, discount
```mermaid
erDiagram
    PROFILES ||--o{ ORDERS : includes
    PROFILES ||--o{ ADDRESSES : includes
    PROFILES ||--o{ SHOPPINGCARTITEMS : includes
    PROFILES ||--o{ FAVOURITES : includes
    
    PROMOCODES ||--o{ ORDERS : includes
    ORDERS ||--o{ ORDERITEMS : includes

    PRODUCTS ||--o{ ORDERITEMS : includes
    PRODUCTS ||--o{ SHOPPINGCARTITEMS : includes
    CATEGORIES ||--|{ PRODUCTS : includes
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
