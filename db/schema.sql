CREATE TABLE IF NOT EXISTS todos (
  id          INTEGER  NOT NULL PRIMARY KEY AUTOINCREMENT,
  subject     TEXT     NOT NULL,
  description TEXT     NOT NULL DEFAULT '',
  created_at  DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at  DATETIME NOT NULL DEFAULT (DATETIME('now')),
  CHECK(subject <> '')
);

CREATE TRIGGER IF NOT EXISTS trigger_todos_updated_at AFTER UPDATE ON todos
BEGIN
  UPDATE todos SET updated_at = DATETIME('now') WHERE id == NEW.id;
END;

CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  user_name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
  CHECK(user_name <> '' AND email <> '' AND password <> '')
);

CREATE TRIGGER IF NOT EXISTS trigger_users_updated_at AFTER UPDATE ON users
BEGIN
  UPDATE users SET updated_at = DATETIME('now') WHERE id == NEW.id;
END;

CREATE TABLE IF NOT EXISTS tokens (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  token TEXT NOT NULL UNIQUE,
  created_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  CHECK(token <> '')
);

CREATE TRIGGER IF NOT EXISTS trigger_tokens_updated_at AFTER UPDATE ON tokens
BEGIN
  UPDATE tokens SET updated_at = DATETIME('now') WHERE id == NEW.id;
END;
