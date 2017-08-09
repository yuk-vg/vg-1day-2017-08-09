-- +migrate Up
CREATE TABLE user (
    id INTEGER NOT NULL PRIMARY KEY,
    username TEXT NOT NULL DEFAULT "",
    sex TEXT NOT NULL DEFAULT "",
    age INTEGER NOT NULL DEFAULT 0
);

-- +migrate Down
DROP TABLE user;
