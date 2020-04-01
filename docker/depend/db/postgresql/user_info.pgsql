CREATE TABLE IF NOT EXISTS user_info
(
    id         SERIAL,
    name       varchar(50)        DEFAULT NULL,
    gender     int                DEFAULT NULL,
    city       varchar(20)        DEFAULT NULL,
    ip         varchar(20)        DEFAULT NULL,
    phone      varchar(20)        DEFAULT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL     DEFAULT NULL,
    PRIMARY KEY (id)
);