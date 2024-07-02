-- modify "access_logs" table
ALTER TABLE "access_logs" ADD COLUMN "forwarded_for" character varying NOT NULL, ADD COLUMN "forwarded_proto" character varying NOT NULL, ADD COLUMN "forwarded_host" character varying NOT NULL, ADD COLUMN "forwarded_port" character varying NOT NULL, ADD COLUMN "forwarded_server" character varying NOT NULL, ADD COLUMN "request_id" character varying NOT NULL, ADD COLUMN "user_agent" character varying NOT NULL;
-- create index "accesslog_uri" to table: "access_logs"
CREATE INDEX "accesslog_uri" ON "access_logs" ("uri");
