-- modify "news_items" table
ALTER TABLE "news_items" ADD COLUMN "paywalled" boolean NOT NULL DEFAULT false;
ALTER TABLE "news_items" ALTER COLUMN "paywalled" DROP DEFAULT;
