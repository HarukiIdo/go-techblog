# go-techblog


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