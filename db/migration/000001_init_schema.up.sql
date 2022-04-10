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
     "id" bigserial PRIMARY KEY,
     "user" varchar NOT NULL,
     "title" varchar NOT NULL,
     "content" varchar NOT NULL,
     "created_at" timestamptz NOT NULL DEFAULT (now()),
     "upvotes" bigint NOT NULL DEFAULT 0
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user") REFERENCES "users" ("username");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "posts" ("user");

CREATE UNIQUE INDEX ON "posts" ("user");