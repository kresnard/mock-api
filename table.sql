CREATE TABLE "public"."apis" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NULL,
    "url" TEXT NULL,
    "method" VARCHAR(20) NULL,
    "response" TEXT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);