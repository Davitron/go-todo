CREATE TABLE IF NOT EXISTS users
(
    id
    serial
    PRIMARY
    KEY,
    username
    VARCHAR
(
    50
) UNIQUE NOT NULL,
    password VARCHAR
(
    300
) NOT NULL,
    email VARCHAR
(
    300
) UNIQUE NOT NULL
    );