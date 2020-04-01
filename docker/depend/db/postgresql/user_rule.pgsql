CREATE TABLE IF NOT EXISTS user_rule
(
    id            SERIAL,
    rule_name     varchar(84) NOT NULL UNIQUE,       -- 名称
    is_activation int                  DEFAULT 1,    -- 是否激活 默认是
    command       text                 DEFAULT NULL, -- 命令
    created_at    timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    timestamp   NULL     DEFAULT NULL,
    PRIMARY KEY (id)
);