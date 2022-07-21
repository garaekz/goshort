CREATE TABLE users
(
    id         VARCHAR PRIMARY KEY,
    email       VARCHAR NOT NULL,
    password       VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT 'now()',
    updated_at TIMESTAMP NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS shorts (
  code VARCHAR NOT NULL UNIQUE,
	original_url TEXT NOT NULL,
  visits INTEGER NOT NULL DEFAULT 0,
	user_id VARCHAR NULL DEFAULT NULL,
  creator_ip VARCHAR NOT NULL,
	created_at TIMESTAMP NULL DEFAULT 'now()',
	updated_at TIMESTAMP NULL DEFAULT NULL,
	deleted_at TIMESTAMP NULL DEFAULT NULL
);

CREATE TABLE keys
(
    id         VARCHAR PRIMARY KEY,
    user_id    VARCHAR NOT NULL,
    key        VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT 'now()',
    updated_at TIMESTAMP NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true
);
