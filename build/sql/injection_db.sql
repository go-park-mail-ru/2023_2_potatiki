CREATE TABLE IF NOT EXISTS public.profiles
(
    id uuid NOT NULL,
    login text,
    description text,
    imgsrc text,
    passwordhash text,
    CONSTRAINT "ProfileId_pkey" PRIMARY KEY (id),
    CONSTRAINT "ProfileLogin_unique" UNIQUE (login)
)

CREATE TABLE IF NOT EXISTS public.products
(
    id uuid NOT NULL,
    nameProduct text,
    description text, //добавить imgSrc
    price int,
    CONSTRAINT "ProductId_pkey" PRIMARY KEY (id),
    CONSTRAINT "ProductName_unique" UNIQUE (nameProduct)
)