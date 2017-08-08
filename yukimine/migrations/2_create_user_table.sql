-- +migrate Up
CREATE TABLE user (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);

-- +migrate Down
DROP TABLE user;
