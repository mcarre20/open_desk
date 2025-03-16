-- +goose Up
CREATE TABLE "caterogies" (
  "id" smallserial PRIMARY KEY,
  "category" varchar(60) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT(now())
);

-- +goose Down
DROP TABLE IF EXISTS caterogies;
