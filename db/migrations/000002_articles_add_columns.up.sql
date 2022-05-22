ALTER TABLE articles
  add column body mediumtext NOT NULL,
  add column created datetime,
  add column updated datetime;

UPDATE articles Set created = CURRENT_TIMESTAMP where created is null;
UPDATE articles Set updated = CURRENT_TIMESTAMP where updated is null;
