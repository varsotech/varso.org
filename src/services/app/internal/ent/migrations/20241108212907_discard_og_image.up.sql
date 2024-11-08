ALTER TABLE "rss_feeds" ADD COLUMN "discard_og_image" boolean NOT NULL DEFAULT false;
ALTER TABLE "rss_feeds" ALTER COLUMN "discard_og_image" DROP DEFAULT;