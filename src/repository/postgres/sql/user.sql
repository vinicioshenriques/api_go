CREATE TABLE IF NOT EXISTS users
(
    id bigserial NOT NULL PRIMARY KEY,
    name text NOT NULL,
    email text NOT NULL
)