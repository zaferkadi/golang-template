CREATE TABLE "genres" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "description" varchar(255) NOT NULL
);

CREATE TABLE "authors" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "bio" varchar(500) NOT NULL,
  "created_at" datetime DEFAULT (now())
);

CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "description" varchar(255) NOT NULL,
  "ISBN" char(13) NOT NULL,
  "genre_id" INT NOT NULL,
  "created_at" datetime DEFAULT (now())
);

CREATE TABLE "books_authors" (
  "book_id" int,
  "author_id" int,
  "is_main_author" BOOLEAN NOT NULL DEFAULT false,
  PRIMARY KEY ("book_id", "author_id")
);

ALTER TABLE "books" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id");

ALTER TABLE "books_authors" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books_authors" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

