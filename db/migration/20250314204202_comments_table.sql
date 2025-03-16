-- +goose Up
CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "user_id" UUID NOT NULL,
  "ticket_id" bigserial NOT NULL,
  "comments" TEXT NOT NULL,
  "customer_visible" bool NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT(now()),
  "updated_at" timestamp NOT NULL DEFAULT(now())
);
ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id");

-- +goose Down
DROP TABLE IF EXISTS comments;
