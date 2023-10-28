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
  "path" varchar(512) NOT NULL
  UNIQUE ("domain", "port", "scheme", "path")
);

CREATE TABLE "counter" (
    "page" integer UNIQUE PRIMARY KEY,
    "count" integer NOT NULL DEFAULT 0
);

CREATE TABLE "view" (
  "time" timestamptz NOT NULL,
  "page" integer NOT NULL,
  "user_agent" text DEFAULT null
);

ALTER TABLE "page" ADD FOREIGN KEY ("domain") REFERENCES "authorised_domain" ("id");
ALTER TABLE "counter" ADD FOREIGN KEY ("page") REFERENCES "page" ("id");
ALTER TABLE "view" ADD FOREIGN KEY ("page") REFERENCES "page" ("id");

SELECT create_hypertable('view', by_range('time'));
CREATE INDEX ix_page_time ON view (page, time DESC);

CREATE OR replace FUNCTION find_page(
    _domain INTEGER, 
    _path VARCHAR(512), 
    _schema SCHEMA, 
    _port INTEGER, 
    OUT _page_id INTEGER
) LANGUAGE plpgsql AS $ func $ 
    BEGIN LOOP 
        SELECT
            id 
        FROM
            page 
        WHERE
            page.domain = _domain 
            AND page.path = _path 
            AND page.scheme = _schema 
            AND page.port = _port INTO _page_id;

        EXIT WHEN FOUND;

        INSERT INTO
            page AS p ( domain, path, scheme, port ) 
        VALUES
            (
                _domain,
                _path,
                _schema,
                _port 
            )
            ON conflict ( domain, path, scheme, port ) DO nothing returning p.id INTO _page_id;

        EXIT WHEN FOUND;
    END LOOP;
END $func$;
