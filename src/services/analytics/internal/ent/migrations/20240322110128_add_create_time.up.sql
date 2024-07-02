-- modify "access_logs" table
ALTER TABLE "access_logs" ADD COLUMN "create_time" timestamptz NOT NULL default NOW();
ALTER TABLE "access_logs" ALTER COLUMN "create_time" DROP DEFAULT;