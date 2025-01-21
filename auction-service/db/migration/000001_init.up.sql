CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Required for gen_random_uuid()

CREATE TABLE "auctions" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "product_id" varchar NOT NULL,
  "user_id" uuid NOT NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NOT NULL,
  "starting_price" float NOT NULL,
  "current_price" float NOT NULL,
  "status" varchar NOT NULL,
  "winner_id" uuid NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "bids" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "auction_id" uuid NOT NULL,
  "amount" float NOT NULL,
  "bid_time" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "auctions" ("user_id");

CREATE INDEX ON "bids" ("auction_id");

COMMENT ON COLUMN "auctions"."status" IS 'it can be active, suspended or ended';

ALTER TABLE "bids" ADD FOREIGN KEY ("auction_id") REFERENCES "auctions" ("id");
