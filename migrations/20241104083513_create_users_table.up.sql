-- Create user table
CREATE TABLE IF NOT EXISTS "users" (
    "user_id" SERIAL,
    "login_id" VARCHAR(100) NOT NULL,
    "password" TEXT NOT NULL,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "session_id" VARCHAR(100),
    "is_active" INT NOT NULL DEFAULT 1,
    "role_id" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" INT,
    "updated_at" TIMESTAMP NULL,
    "updated_by" INT,
    "deleted_at" TIMESTAMPTZ,
    "deleted_by" INT,
	CONSTRAINT "pk_users_user_id" PRIMARY KEY ("user_id")
);

-- Create roles table
CREATE TABLE IF NOT EXISTS "roles" (
    "role_id" SERIAL,
    "name" VARCHAR(100) NOT NULL,
    "description" TEXT,
    "is_active" INT NOT NULL DEFAULT 1,
    "order" INT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" INT,
    "updated_at" TIMESTAMP NULL,
    "updated_by" INT,
    "deleted_at" TIMESTAMPTZ,
    "deleted_by" INT,
	CONSTRAINT "pk_roles_role_id" PRIMARY KEY ("role_id")
);

ALTER TABLE "users"
ADD CONSTRAINT "fk_users_role_id" FOREIGN KEY ("role_id")
REFERENCES "roles" ("role_id") ON DELETE NO ACTION ON UPDATE NO ACTION;