CREATE TABLE "shopping_cart" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "cart_items" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "cart_id" varchar NOT NULL,
  "product_id" varchar NOT NULL,
  "quantity" bigint NOT NULL
);

CREATE TABLE "orders" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" varchar NOT NULL,
  "payment_method" varchar NOT NULL,
  "total_cost" bigint NOT NULL,
  "order_date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "order_id" varchar NOT NULL,
  "product_id" varchar NOT NULL,
  "quantity" bigint NOT NULL,
  "price_at_purchase" bigint NOT NULL
);

DROP TABLE IF EXISTS purchases;

CREATE INDEX ON "products" ("user_id");

CREATE INDEX ON "shopping_cart" ("user_id");

CREATE INDEX ON "cart_items" ("cart_id");

CREATE INDEX ON "cart_items" ("product_id");

CREATE INDEX ON "orders" ("user_id");

CREATE INDEX ON "order_items" ("order_id");

CREATE INDEX ON "order_items" ("product_id");

ALTER TABLE "shopping_cart" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("cart_id") REFERENCES "shopping_cart" ("id");

ALTER TABLE "cart_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");