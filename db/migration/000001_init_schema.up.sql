CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" varchar NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" bigint NOT NULL,
  "category_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "purchases" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "product_id" varchar NOT NULL,
  "quantity" bigint NOT NULL,
  "total_price" bigint NOT NULL,
  "is_paid" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE books (
  "id" SERIAL PRIMARY KEY,
  "data" jsonb
);

CREATE INDEX ON "products" ("category_id");

CREATE INDEX ON "purchases" ("product_id");

CREATE UNIQUE INDEX ON "purchases" ("user_id", "product_id");

ALTER TABLE "products" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "purchases" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "purchases" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

CREATE VIEW "categories_with_products" AS
SELECT
    "categories".*,
    JSONB_AGG("products".*) AS "products"
FROM
    "categories"
LEFT JOIN "products" ON "products"."category_id" = "categories"."id"
GROUP BY
    "categories"."id";

CREATE VIEW "users_with_products" AS
SELECT
    "users".*,
    JSONB_AGG("products".*) AS "products"
FROM
    "users"
LEFT JOIN "products" ON "products"."user_id" = "users"."username"
GROUP BY
    "users"."username";