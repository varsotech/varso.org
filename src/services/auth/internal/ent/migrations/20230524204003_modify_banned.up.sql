-- modify "users" table
ALTER TABLE "users" DROP COLUMN "banned", ADD COLUMN "ban_time" timestamptz NULL;
