
DROP TABLE IF EXISTS "urls";
CREATE TABLE IF NOT EXISTS "urls" (
	"id" VARCHAR(36) NOT NULL,
	"created_at" TIMESTAMP NULL DEFAULT 'now()',
	"updated_at" TIMESTAMP NULL DEFAULT NULL,
	"deleted_at" TIMESTAMP NULL DEFAULT NULL,
	"code" VARCHAR NOT NULL UNIQUE,
	"original_url" TEXT NOT NULL,
	"user_id" VARCHAR(36) NULL DEFAULT NULL,
	PRIMARY KEY ("id")
);


DROP TABLE IF EXISTS "users";
CREATE TABLE IF NOT EXISTS "users" (
	"id" VARCHAR(36) NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"password" VARCHAR(255) NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"updated_at" TIMESTAMP NULL DEFAULT NULL,
	"is_active" BOOLEAN NOT NULL,
	PRIMARY KEY ("id")
);

INSERT INTO users (id, email, password, created_at, is_active) 
VALUES ('3d7bd6e8-983e-4da5-b6c2-5c483d3ea35f', 
't@t.io', 
'$2a$10$4TUSdE1mVKGndeE4n7gF5uhAJgB64uaoe0AYfDoYidqXr89Jg5Z4q', 
'2019-01-01 00:00:00', true);

ALTER TABLE urls ADD CONSTRAINT "urls_user_id_foreign" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE CASCADE
