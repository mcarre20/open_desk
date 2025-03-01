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

CREATE TABLE "tickets" (
  "id" bigserial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "assigned_to" UUID,
  "description" TEXT NOT NULL,
  "status" int NOT NULL DEFAULT(0),
  "priority" int NOT NULL DEFAULT(0),
  "category_id" UUID,
  "updated_at" timestamp NOT NULL DEFAULT(now()),
  "created_at" timestamp NOT NULL DEFAULT(now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "ticket_id" bigserial NOT NULL,
  "comments" TEXT NOT NULL,
  "customer_visible" bool NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT(now()),
  "updated_at" timestamp NOT NULL DEFAULT(now())
);

CREATE TABLE "caterogies" (
  "id" smallserial PRIMARY KEY,
  "category" varchar(60) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT(now())
);

ALTER TABLE "tickets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tickets" ADD FOREIGN KEY ("assigned_to") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id");