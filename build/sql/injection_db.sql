CREATE TABLE IF NOT EXISTS public.profiles
(
    id uuid NOT NULL,
    login text,
    description text,
    imgsrc text,
    passwordhash text,
    CONSTRAINT "ProfileId_pkey" PRIMARY KEY (id),
    CONSTRAINT "ProfileLogin_pkey" UNIQUE (login)
)