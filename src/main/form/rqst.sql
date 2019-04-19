-- name: create-groups-table
CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    readonly VARCHAR(255),
    readwrite VARCHAR(255)
)

-- name: create-users-table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    pass VARCHAR(255),
    active BOOLEAN NOT NULL,
    groups VARCHAR(255),
    readonly VARCHAR(255),
    readwrite VARCHAR(255)
);

-- name: add-user
INSERT INTO users (name, pass, active, groups, readonly, readwrite) VALUES (?, ?, ?, ?, ?, ?)

--name: del-user
DELETE * FROM users WHERE id = ? 

-- name: get-user-by-name
SELECT * FROM users WHERE name = ? LIMIT 1

-- name: get-user-by-id
SELECT * FROM users WHERE id = ? LIMIT 1

-- name: list-users-safe // no password
SELECT id, name, active, groups, readonly, readwrite FROM users

-- name: add-group
INSERT INTO groups (name, readonly, readwrite) VALUES (?, ?, ?)

-- name: del-group
DELETE * FROM groups WHERE id = ?

--name: get-group-by-id
SELECT * FROM groups WHERE id = ? LIMIT 1

--name: get-group-by-name
SELECT * FROM groups WHERE name = ? LIMIT 1

--name: list-groups
SELECT * FROM groups