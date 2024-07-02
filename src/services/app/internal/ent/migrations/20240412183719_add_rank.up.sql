-- modify "news_items" table
ALTER TABLE "news_items" ADD COLUMN "rank" bigint NOT NULL DEFAULT 0;
ALTER TABLE "news_items" ALTER COLUMN "rank" DROP DEFAULT;
