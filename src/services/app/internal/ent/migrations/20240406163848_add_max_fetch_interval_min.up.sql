-- modify "rss_feeds" table
ALTER TABLE "rss_feeds" ADD COLUMN "max_fetch_interval_min" bigint NULL;
