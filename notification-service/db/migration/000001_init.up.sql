CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- Required for gen_random_uuid()

CREATE TABLE "verify_emails" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "email" varchar NOT NULL,
  "user_name" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minute'),
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);