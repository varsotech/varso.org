-- modify "users" table
ALTER TABLE "users" ADD COLUMN "delete_time" timestamptz NULL, ADD COLUMN "banned" boolean NOT NULL DEFAULT false;
