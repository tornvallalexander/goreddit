CREATE TABLE "subreddits" (
    "name" varchar UNIQUE PRIMARY KEY,
    "moderator" varchar NOT NULL,
    "followers" bigint NOT NULL DEFAULT 0,
    "description" varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "subreddits" ADD FOREIGN KEY ("moderator") REFERENCES "users" ("username");