CREATE TYPE http_methods AS ENUM ('POST', 'GET', 'PUT', 'DELETE', 'PATCH');
CREATE TABLE "public"."apis" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NULL,
    "url" TEXT NULL,
    "method" http_methods NULL,
    "response" TEXT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW()
);