DROP DATABASE IF EXISTS go_db

CREATE DATABASE go_db;

use go_db;

CREATE TABLE IF NOT EXISTS articles(
  id int AUTO_INCREMENT,
  title varchar(100),
  PRIMARY KEY(id)
); 

ALTER TABLE articles
  add column body mediumtext NOT NULL,
  add column created datetime,
  add column updated datetime;

UPDATE articles 
Set created = CURRENT_TIMESTAMP 
where created is null;

UPDATE articles 
Set updated = CURRENT_TIMESTAMP 
where updated is null;