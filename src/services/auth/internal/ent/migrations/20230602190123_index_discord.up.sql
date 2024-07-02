-- create index "user_discord_user_id" to table: "users"
CREATE UNIQUE INDEX "user_discord_user_id" ON "users" ("discord_user_id");
