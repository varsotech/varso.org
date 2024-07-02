-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "ban_time", ADD COLUMN "banned" boolean NOT NULL DEFAULT false;
