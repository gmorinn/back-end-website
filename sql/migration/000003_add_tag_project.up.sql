BEGIN;

CREATE TYPE "project_tag" AS ENUM (
    'webdevelopment',
    'socialmedia'
);

-- add "tag" project_tag NOT NULL DEFAULT 'webdevelopment' to projects
ALTER TABLE "projects" ADD COLUMN "tag" project_tag NOT NULL DEFAULT 'webdevelopment';

-- change column "language" to "language" text DEFAULT NULL
ALTER TABLE "projects" ALTER COLUMN "language" TYPE text;

COMMIT; 