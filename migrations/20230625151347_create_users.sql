-- migrate:up
CREATE TABLE users (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),

    username VARCHAR(64) NOT NULL,
    full_name VARCHAR(256) NOT NULL,
    email VARCHAR(128) NOT NULL,
    password VARCHAR(128) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (id)
);

CREATE INDEX ON users (deleted_at);

-- migrate:down
DROP TABLE users;
