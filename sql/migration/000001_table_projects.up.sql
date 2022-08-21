BEGIN;

  CREATE TABLE "projects" (
    "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
    "deleted_at" timestamptz,
    "user_id" uuid NOT NULL,
    "title" text NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 2 AND char_length(title) <= 50),
    "content" text NOT NULL CONSTRAINT contentchk CHECK (char_length(content) >= 2 AND char_length(content) <= 500),
    "img_cover" text NOT NULL,
    "img_description" text NOT NULL,
    "language" text NOT NULL CONSTRAINT languagechk CHECK (language ~* '^[A-Za-z]{2}$'),
    "url" text NOT NULL
  );

  ALTER TABLE "projects" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;


COMMIT; 