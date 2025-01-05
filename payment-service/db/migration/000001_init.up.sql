CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Required for gen_random_uuid()

CREATE TABLE "stripe_customers" (
  "id" varchar NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "orders" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_name" varchar NOT NULL,
  "buyer_id" uuid NOT NULL,
  "seller_id" uuid NOT NULL,
  "total_price" float NOT NULL,
  "delivery_address" varchar NOT NULL,
  "country" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT 'processing',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "order_items" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "item_name" varchar NOT NULL,
  "item_sub_total" float NOT NULL,
  "quantity" bigint NOT NULL,
  "item_id" varchar NOT NULL,
  "order_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "payments" (
  "id" varchar PRIMARY KEY,
  "amount" float NOT NULL,
  "currency" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT 'processing',
  "user_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");