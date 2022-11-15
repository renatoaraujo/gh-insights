-- Yeah, I want to create migrations for this, but I will do it later.

DROP TABLE IF EXISTS public.repositories;
CREATE TABLE IF NOT EXISTS public.repositories
(
    id         INT PRIMARY KEY,
    name       VARCHAR(50) UNIQUE  NOT NULL,
    url        VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS public.issues;
CREATE TABLE IF NOT EXISTS public.issues
(
    id            INT PRIMARY KEY,
    repository_id INT       NOT NULL,
    title         VARCHAR,
    number        INT,
    state         VARCHAR,
    opened_at     TIMESTAMP NOT NULL,
    closed_at     TIMESTAMP,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS public.pulls;
CREATE TABLE IF NOT EXISTS public.pulls
(
    id            INT PRIMARY KEY,
    repository_id INT       NOT NULL,
    title         VARCHAR,
    number        INT,
    state         VARCHAR,
    opened_at     TIMESTAMP NOT NULL,
    closed_at     TIMESTAMP,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);