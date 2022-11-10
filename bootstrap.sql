CREATE TABLE IF NOT EXISTS public.repositories (
      id INT PRIMARY KEY,
      name VARCHAR(50) UNIQUE NOT NULL,
      url VARCHAR ( 255 ) UNIQUE NOT NULL,
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS public.issues (
    id INT PRIMARY KEY,
    repository_id INT NOT NULL,
    opened_at TIMESTAMP NOT NULL,
    closed_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO repositories VALUES (12983, 'facebook', 'https://facebook.com');

SELECT * FROM repositories;


