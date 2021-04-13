/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Volcando estructura para tabla urls
DROP TABLE IF EXISTS "urls";
CREATE TABLE IF NOT EXISTS "urls" (
	"id" VARCHAR(36) NOT NULL,
	"created_at" TIMESTAMP NULL DEFAULT 'now()',
	"updated_at" TIMESTAMP NULL DEFAULT NULL,
	"deleted_at" TIMESTAMP NULL DEFAULT NULL,
	"code" VARCHAR NOT NULL,
	"original_url" TEXT NOT NULL,
	"user_id" VARCHAR(36) NULL DEFAULT NULL,
	PRIMARY KEY ("id"),
	INDEX "idx_urls_deleted_at" ("deleted_at"),
	UNIQUE INDEX "urls_code_key" ("code"),
	CONSTRAINT "urls_user_id_foreign" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);


-- Volcando estructura para tabla users
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
