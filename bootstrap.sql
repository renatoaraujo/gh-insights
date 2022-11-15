CREATE TABLE IF NOT EXISTS public.repositories
(
    id         INT PRIMARY KEY,
    name       VARCHAR(50) UNIQUE  NOT NULL,
    url        VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP
);

TRUNCATE public.repositories;

DROP TABLE public.issues;
CREATE TABLE IF NOT EXISTS public.issues
(
    id            INT PRIMARY KEY,
    repository_id INT       NOT NULL,
    title         VARCHAR,
    number        INT,
    opened_at     TIMESTAMP NOT NULL,
    closed_at     TIMESTAMP NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

TRUNCATE public.issues;


SELECT *
FROM ISSUES;