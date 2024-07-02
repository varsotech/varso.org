-- modify "rss_feeds" table
ALTER TABLE "rss_feeds" ADD COLUMN "content_whitelist_regex" character varying NULL;
