-- modify "news_items" table
ALTER TABLE "news_items" ADD COLUMN "deleted" boolean NOT NULL DEFAULT false, ADD COLUMN "delete_time" timestamptz NULL;
