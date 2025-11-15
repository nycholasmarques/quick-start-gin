-- Create "users" table
CREATE TABLE "public"."users" (
  "user_id" bigserial NOT NULL,
  "public_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "first_name" character varying(150) NOT NULL,
  "last_name" character varying(150) NOT NULL,
  "email" character varying(255) NOT NULL,
  "password_hash" character varying(255) NOT NULL,
  "phone" character varying(20) NOT NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "email_verified" boolean NOT NULL DEFAULT false,
  "phone_verified" boolean NOT NULL DEFAULT false,
  "last_login_at" timestamptz NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("user_id"),
  CONSTRAINT "users_email_key" UNIQUE ("email"),
  CONSTRAINT "users_phone_key" UNIQUE ("phone")
);
-- Create index "idx_users_email" to table: "users"
CREATE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create index "idx_users_is_active" to table: "users"
CREATE INDEX "idx_users_is_active" ON "public"."users" ("is_active");
-- Create index "idx_users_public_id" to table: "users"
CREATE INDEX "idx_users_public_id" ON "public"."users" ("public_id");
