ALTER TABLE articles
  add column body mediumtext NOT NULL,
  add column createdat datetime,
  add column updatedat datetime;

UPDATE articles 
Set createdat = CURRENT_TIMESTAMP 
where createdat is null;

UPDATE articles 
Set updatedat = CURRENT_TIMESTAMP 
where updatedat is null;
