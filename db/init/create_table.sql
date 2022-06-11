CREATE DATABASE IF NOT EXISTS go_db;

use go_db;

CREATE TABLE IF NOT EXISTS articles(
  id int AUTO_INCREMENT,
  title varchar(100),
  PRIMARY KEY(id)
); 