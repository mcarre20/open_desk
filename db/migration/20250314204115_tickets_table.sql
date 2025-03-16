-- +goose Up
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
ALTER TABLE "tickets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "tickets" ADD FOREIGN KEY ("assigned_to") REFERENCES "users" ("id");

-- +goose Down
DROP TABLE IF EXISTS tickets;
