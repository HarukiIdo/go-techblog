# go-techblog
個人用の技術プログサイトです。

## 動機
学んだ技術の陳腐化を防ぐため、QiitaやZennなどのブログサイトに技術的なアウトプットをしたいかったが、敷居が高く重い腰が上がらなかったため、気軽にアウトプットできる環境が欲しかった

## 開発してみて
フロントエンド〜バックエンド・インフラの実装からデプロイし、一つのWebサイトを公開する経験が出来た点では、全体像を理解する良い機会となった。 

## 課題
- フロントエンドにReactやVueなどのフレームワークを導入してUIをよくする
- 自己紹介用のサイトに組み込む
- Ginなどのフレームワークを用いてコード量を減らす
- テストコードを書く & CI環境を構築する

## 使用技術
フロントエンド：HTML,CSS(BootStrap),JavaScript  
バックエンド：Go(Echo)+sqlx  
DB：MySQL  
インフラ：Docker, Docker Compose, Heroku  

## DBマイグーイション
### ファイルの作成
```
migrate create -ext sql -dir YOUR_DATABASE_DIRECTORY -seq YOUR_FILE_NAME
```
ex
```
migrate create -ext sql -dir db/migrations -seq create_aritcle
```

### マイグレーションの実行
```
migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up MIGRATION_VERSION
```

- ex
    - 000002_articles_add_columns.upが実行される
```
migrate -database "mysql://go_user:password@tcp(127.0.0.1:3306)/go_db?multiStatements=true" -path=db/migrations/ up 1
```