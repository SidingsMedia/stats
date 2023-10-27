-- SPDX-FileCopyrightText: 2023 Sidings Media
-- SPDX-License-Identifier: MIT

CREATE TYPE "schema" AS ENUM (
  'http',
  'https'
);

CREATE TABLE "authorised_domain" (
  "id" SERIAL UNIQUE PRIMARY KEY,
  "domain" varchar(253) UNIQUE NOT NULL
);

CREATE TABLE "page" (
  "id" SERIAL UNIQUE PRIMARY KEY,
  "domain" integer NOT NULL,
  "port" integer DEFAULT NULL,
  "scheme" schema NOT NULL,
  "path" varchar(512) UNIQUE NOT NULL
);

CREATE TABLE "view" (
  "time" timestamptz NOT NULL,
  "page" integer NOT NULL,
  "user_agent" text DEFAULT null
);

ALTER TABLE "page" ADD FOREIGN KEY ("domain") REFERENCES "authorised_domain" ("id");
ALTER TABLE "view" ADD FOREIGN KEY ("page") REFERENCES "page" ("id");

SELECT create_hypertable('view', by_range('time'));
CREATE INDEX ix_page_time ON view (page, time DESC);
