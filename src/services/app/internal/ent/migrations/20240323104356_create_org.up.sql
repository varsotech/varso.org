-- create "organizations" table
CREATE TABLE "organizations" ("uuid" uuid NOT NULL, "create_time" timestamptz NOT NULL, "name" character varying NOT NULL, "rss_feed_url" character varying NOT NULL, PRIMARY KEY ("uuid"));
-- create index "organization_uuid" to table: "organizations"
CREATE UNIQUE INDEX "organization_uuid" ON "organizations" ("uuid");
