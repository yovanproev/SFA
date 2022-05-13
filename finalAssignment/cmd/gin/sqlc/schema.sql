-- schema.sql
CREATE TABLE tasks (
  id   INTEGER  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  text TEXT NOT NULL,
  listId INTEGER NOT NULL,
  userId INTEGER NOT NULL,
  completed BIT NOT NULL
);

CREATE TABLE lists (
  id   INTEGER  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name TEXT NOT NULL,
  userId INTEGER NOT NULL
);

CREATE TABLE users (
  id   INTEGER  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  datestamp TIMESTAMP NOT NULL
);

