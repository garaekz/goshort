CREATE TABLE users
(
  id         VARCHAR PRIMARY KEY,
  email       VARCHAR NOT NULL UNIQUE,
  password       VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT 'now()',
  updated_at TIMESTAMP NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS shorts (
  code VARCHAR PRIMARY KEY,
	original_url TEXT NOT NULL,
  visits INTEGER NOT NULL DEFAULT 0,
	user_id VARCHAR NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  creator_ip VARCHAR NOT NULL,
	created_at TIMESTAMP NULL DEFAULT 'now()',
	updated_at TIMESTAMP NULL DEFAULT NULL,
	deleted_at TIMESTAMP NULL DEFAULT NULL
);

CREATE TABLE keys
(
  key         VARCHAR PRIMARY KEY,
  user_id     VARCHAR NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT 'now()',
  updated_at TIMESTAMP NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT true
);


CREATE TABLE configs
(
  name         VARCHAR PRIMARY KEY,
  value        VARCHAR NOT NULL
);

CREATE TABLE roles
(
  id         VARCHAR PRIMARY KEY,
  name       VARCHAR NOT NULL
);

CREATE TABLE role_users
(
  role_id    VARCHAR NOT NULL REFERENCES roles (id) ON DELETE CASCADE,
  user_id    VARCHAR NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  PRIMARY KEY (role_id, user_id)
);