# shamo

## ローカルでの起動コマンド
1. docker compose build
2. docker compose run --rm api go run scripts/migrate.go
3. docker compose up -d

## prodでの起動コマンド

`docker compose -f docker-compose.prod.yml up -d`

## TODO
- t2.nanoでの起動時、メモリ不足によりmysqlのコンテナが立ち上がらない...
