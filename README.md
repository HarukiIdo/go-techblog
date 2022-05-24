# go-techblog

## 概要
個人用の技術プログサイト

## 動機
学んだ技術の陳腐化を防ぐため、QiitaやZennなどのブログサイトに技術的なアウトプットをしたかったが、敷居が高く重い腰が上がらなかったため、気軽にアウトプットできる環境が欲しかった

## 技術的な挑戦
- 開発環境構築にDocker・Docker Composeを採用し、コマンドひとつでサーバとDBが起動する構成にした点
- Dockerコンテナで起動後ファイルの変更を検知し、自動でビルドをおこなってくれる「Air」というライブラリを導入し、開発速度を向上させた点
- Dockerを実際に導入してみて、アプリとして起動するまでの時間が明らかに短縮されただけでなく、新しく外部パッケージをインストールしたり、消したりが手軽に行えたり、コードで構成を管理できるためアプリ全体の構成を把握しやすいなどのメリットを感じている

## 課題
- HTML, CSS, JavaScriptをそのまま使っており、コード量が多くなっている
- 削除ボタンや新規作成ボタンなど、サイト閲覧者から見える必要のない要素が入っている

## TODO
- 認証機能を追加する
- ReactやVueなどのフレームワークを導入して、冗長になっている箇所をコンポーネント化するなど、フロントエンドのコード量を減らす
- 自己紹介用のサイトに組み込み、UXを向上させる
- テストコードを書く & CI環境を構築する

## 開発してみて
フロントエンド〜バックエンド・インフラの実装からデプロイまでを体験し、一つのWebサイトを公開する経験が出来た点では、開発工程の全体を掴む良い機会となった。 

## 使用技術
フロントエンド：HTML,CSS,JavaScript  
バックエンド：Go(Echo)
DB：MySQL  
インフラ：Docker, Docker Compose, Heroku  

## DBマイグーイション（golang-migrate/migrate）
[GitHubリポジトリ](https://github.com/golang-migrate/migrate)
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

## 実行
```
docker compose up -d
```
## 停止
```
docker compose down
```