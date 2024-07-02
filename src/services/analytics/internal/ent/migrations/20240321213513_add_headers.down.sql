-- reverse: create index "accesslog_uri" to table: "access_logs"
DROP INDEX "accesslog_uri";
-- reverse: modify "access_logs" table
ALTER TABLE "access_logs" DROP COLUMN "user_agent", DROP COLUMN "request_id", DROP COLUMN "forwarded_server", DROP COLUMN "forwarded_port", DROP COLUMN "forwarded_host", DROP COLUMN "forwarded_proto", DROP COLUMN "forwarded_for";
