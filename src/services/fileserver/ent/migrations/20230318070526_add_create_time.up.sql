-- modify "checkpoints" table
ALTER TABLE "checkpoints" ADD COLUMN "create_time" timestamptz NOT NULL;
