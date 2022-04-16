-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-04-16T16:22:03.797Z

CREATE TABLE "users" (
  "username" varchar UNIQUE PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "karma" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "posts" (
  "id" bigint PRIMARY KEY,
  "user" varchar NOT NULL,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "subreddit" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "upvotes" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "subreddits" (
  "name" varchar UNIQUE PRIMARY KEY,
  "moderator" varchar NOT NULL,
  "followers" bigint NOT NULL DEFAULT 0,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "posts" ("user");

CREATE UNIQUE INDEX ON "posts" ("user");

ALTER TABLE "posts" ADD FOREIGN KEY ("user") REFERENCES "users" ("username");

ALTER TABLE "posts" ADD FOREIGN KEY ("subreddit") REFERENCES "subreddits" ("name");

ALTER TABLE "subreddits" ADD FOREIGN KEY ("moderator") REFERENCES "users" ("username");
