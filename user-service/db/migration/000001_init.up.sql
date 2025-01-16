CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Required for gen_random_uuid()

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL,
  "is_email_verified" boolean NOT NULL DEFAULT false,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "password_reset_tokens" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "token" TEXT NOT NULL UNIQUE,
  "expires_at" TIMESTAMP NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "stripe_accounts" (
  "id" varchar NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "password_reset_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "stripe_accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
