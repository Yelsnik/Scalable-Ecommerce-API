CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Required for gen_random_uuid()

CREATE TABLE "cartitems" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "cart" uuid  NOT NULL,
  "product" varchar  NOT NULL,
  "quantity" bigint NOT NULL,
  "price" float NOT NULL,
  "currency" varchar NOT NULL,
  "sub_total" float NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "carts" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "total_price" float NOT NULL
);

COMMENT ON COLUMN "cartitems"."price" IS 'must be positive';

ALTER TABLE "cartitems" ADD FOREIGN KEY ("cart") REFERENCES "carts" ("id");