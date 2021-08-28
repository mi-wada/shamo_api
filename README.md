# shamo

## prodでの起動コマンド

`docker compose -f docker-compose.prod.yml up -d`


## ローカルでビルドしてデプロイ
`scp -i ~/.ssh/udemy_sample.pem -r ~/playground/shamo/front/build  ec2-user@13.230.90.123:/home/ec2-user/shamo/front`

## TODO
- t2.nanoでの起動時、メモリ不足によりmysqlのコンテナが立ち上がらない...
