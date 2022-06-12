DROP DATABASE IF EXISTS go_db

CREATE DATABASE go_db;

use go_db;

CREATE TABLE IF NOT EXISTS articles(
  id int AUTO_INCREMENT,
  title varchar(100),
  body mediumtext NOT NULL,
  created datetime,
  updated datetime,
  PRIMARY KEY(id)
); 