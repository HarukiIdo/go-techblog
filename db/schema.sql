CREATE TABLE IF NOT EXISTS articles (
  id int AUTO_INCREMENT,
  title varchar(100),
  PRIMARY KEY (id)
);

INSERT INTO articles(id, title) VALUES
(0, 'タイトル1'),
(0, 'タイトル2'),
(0, 'タイトル3');