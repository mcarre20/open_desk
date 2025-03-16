-- +goose Up
CREATE TABLE "users" (
  "id" UUID PRIMARY KEY DEFAULT(gen_random_uuid()),
  "username" varchar(30) UNIQUE NOT NULL,
  "first_name" varchar(50) NOT NULL,
  "last_name" varchar(50) NOT NULL,
  "email" varchar(50) NOT NULL,
  "hashed_password" varchar(50) NOT NULL,
  "user_role" int NOT NULL,
  "active" bool NULL DEFAULT(FALSE),
  "created_at" timestamp NOT NULL DEFAULT(now()),
  "updated_at" timestamp NOT NULL NOT NULL DEFAULT(now()),
  "password_updated_at" timestamp NOT NULL NOT NULL DEFAULT(now())
);

-- +goose Down
DROP TABLE IF EXISTS users;