-- Drop the table if it exists to avoid errors
DROP TABLE IF EXISTS "public"."apis";
-- Create the table with the unique constraint on the 'name' column
CREATE TABLE "public"."apis" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "url" TEXT NULL,
    "method" VARCHAR(20) NULL,
    "response" TEXT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);