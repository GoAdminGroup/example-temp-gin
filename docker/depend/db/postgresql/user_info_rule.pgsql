CREATE TABLE IF NOT EXISTS user_info_rule
(
    id            SERIAL,
    user_info_id  integer,                      -- task config id
    rule_id       integer,                      -- task rule id
    is_activation int                DEFAULT 0, -- 是否激活 默认否 0
    created_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    timestamp NULL     DEFAULT NULL,
    PRIMARY KEY (id)
);
