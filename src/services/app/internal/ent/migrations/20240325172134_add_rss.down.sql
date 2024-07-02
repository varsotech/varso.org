-- reverse: create "news_item_authors" table
DROP TABLE "news_item_authors";
-- reverse: create index "rssauthor_uuid" to table: "rss_authors"
DROP INDEX "rssauthor_uuid";
-- reverse: create index "rssauthor_name_rss_author_organization" to table: "rss_authors"
DROP INDEX "rssauthor_name_rss_author_organization";
-- reverse: create index "rssauthor_name" to table: "rss_authors"
DROP INDEX "rssauthor_name";
-- reverse: create "rss_authors" table
DROP TABLE "rss_authors";
-- reverse: create index "persons_email_key" to table: "persons"
DROP INDEX "persons_email_key";
-- reverse: create index "person_uuid" to table: "persons"
DROP INDEX "person_uuid";
-- reverse: create index "person_email" to table: "persons"
DROP INDEX "person_email";
-- reverse: create "persons" table
DROP TABLE "persons";
-- reverse: create index "newsitem_uuid" to table: "news_items"
DROP INDEX "newsitem_uuid";
-- reverse: create index "newsitem_rss_guid" to table: "news_items"
DROP INDEX "newsitem_rss_guid";
-- reverse: create index "news_items_rss_guid_key" to table: "news_items"
DROP INDEX "news_items_rss_guid_key";
-- reverse: create "news_items" table
DROP TABLE "news_items";
-- reverse: create index "rssfeed_uuid" to table: "rss_feeds"
DROP INDEX "rssfeed_uuid";
-- reverse: create index "rssfeed_rss_feed_url" to table: "rss_feeds"
DROP INDEX "rssfeed_rss_feed_url";
-- reverse: create index "rss_feeds_rss_feed_url_key" to table: "rss_feeds"
DROP INDEX "rss_feeds_rss_feed_url_key";
-- reverse: create "rss_feeds" table
DROP TABLE "rss_feeds";
-- reverse: create index "organization_unique_name" to table: "organizations"
DROP INDEX "organization_unique_name";
-- reverse: modify "organizations" table
ALTER TABLE "organizations" DROP COLUMN "website_url", DROP COLUMN "description", DROP COLUMN "unique_name", ADD COLUMN "rss_feed_url" character varying NOT NULL;
