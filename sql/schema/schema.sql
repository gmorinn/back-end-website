CREATE EXTENSION pgcrypto;

CREATE TYPE "role" AS ENUM (
  'admin',
  'pro',
  'user'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
  "email" text NOT NULL CONSTRAINT emailchk CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
  "password" text NOT NULL CONSTRAINT passwordchk CHECK (char_length(password) >= 9), 
  "firstname" text CONSTRAINT firstnamehk CHECK (char_length(firstname) >= 2 AND char_length(firstname) <= 20 AND  firstname ~ '^[^0-9]*$') DEFAULT NULL,
  "lastname" text CONSTRAINT lastnamehk CHECK (char_length(lastname) >= 2 AND char_length(lastname) <= 20 AND  lastname ~ '^[^0-9]*$') DEFAULT NULL,
  "role" role NOT NULL DEFAULT 'user'
);

CREATE TABLE "blogs" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamptz,
  "user_id" uuid NOT NULL,
  "title" text NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 2 AND char_length(title) <= 50),
  "content" text NOT NULL CONSTRAINT contentchk CHECK (char_length(content) >= 2 AND char_length(content) <= 500),
  "image" text NOT NULL
)

-- CREATE TABLE "projects" (
--   "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
--   "created_at" timestamptz NOT NULL DEFAULT (NOW()),
--   "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
--   "deleted_at" timestamptz,
--   "user_id" uuid NOT NULL,
--   "title" text NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 2 AND char_length(title) <= 50),
--   "content" text NOT NULL CONSTRAINT contentchk CHECK (char_length(content) >= 2 AND char_length(content) <= 500),
--   "image" text NOT NULL,
--   "language" text NOT NULL CONSTRAINT languagechk CHECK (language ~* '^[A-Za-z]{2}$'),
--   "url" text NOT NULL
-- )

CREATE TABLE "refresh_token" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
  "deleted_at" timestamptz,
  "token" text NOT NULL,
  "ip" text NOT NULL,
  "user_agent" text NOT NULL,
  "expir_on" timestamptz NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "files" (
    "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
    "deleted_at" timestamptz,
    "name" text,
    "url" text,
    "mime" text,
    "size" bigint
);

ALTER TABLE "refresh_token" ADD FOREIGN KEY ("user_id") REFERENCES "students" ("id") ON DELETE CASCADE;
ALTER TABLE "blogs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "blogs" ADD FOREIGN KEY ("image") REFERENCES "files" ("url") ON DELETE CASCADE;
-- ALTER TABLE "projects" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "projects" ADD FOREIGN KEY ("image") REFERENCES "files" ("url") ON DELETE CASCADE;
