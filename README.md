# go-techblog
## 概要
個人用の技術系プログサイト

## 動機
学んだ技術の陳腐化を防ぐため、QiitaやZennなどのブログサイトに技術的なアウトプットをしたかったが、敷居が高く重い腰が上がらなかったため、気軽にアウトプットできる環境が欲しかった

## 技術的な挑戦や工夫した点
- 開発環境構築にDocker・Docker Composeを導入して、ローカル環境を汚さないようにした点
- GoのフレームワークやORMパッケージを適宜導入して、冗長なコードを減らし、作業工数を削減した点
- DIを導入して、テスタが書きやすい設計を意識して実装した点
- ホットリロードを行うパッケージを導入し、コードを更新するたびにdockerを再起動する手間を削減した点
- DBのパスワードなどセキュリティレベルの高い情報を環境変数に設定することで、秘匿性を向上させた点
- Basic認証により、パスワードを知っている人だけが記事に関する操作ができるようにした点

## 課題
- HTML, CSS, JavaScriptをそのまま使っており、コード量が多くなっている
- 編集ボタンのような、サイト閲覧者から見える必要のない要素がページに入っているので、認証によりアクセス制御する必要がある

## TODO
- OAuth認証のようなJWTなどアクセストークンを用いるよりセキュリティが考慮された認証機能を導入して、アクセス制御をする
- ReactやVueなどのフレームワークを導入して、冗長になっている箇所をコンポーネント化するなど、フロントエンドのコード量を減らしたり、Webページのパフォーマンスを意識してみる
- 作ったアプリをホスティングして、公開する
  - AWS ECR, ECS, RDSを使った構成に検討中
- 自己紹介用のサイトに組み込み、UXを向上させる
- テストコードを書く & Github Actionsなどを使ってCI環境を構築する

## 開発してみて
- フロントエンドの実装〜APIサーバやRDBの処理を一通りして、一つのWebアプリとして作りあげる経験をしたことで、開発工程の全体を掴む良い機会となった。

- 今回Webアプリ開発に必須のDockerを使ってみたが、比較的簡単に環境構築を行うことができ、便利なツールだと感じた。今回はいろんなサイトや公式ドキュメントを通じて手探りで作ったため、もう少しcomposeやDockerfileの書き方を調査して、使いこなせるようになりたい。

## 使用技術
フロントエンド：HTML,CSS,JavaScript  
バックエンド：Go(Echo)
DB：MySQL  
インフラ：Docker, Docker Compose 

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