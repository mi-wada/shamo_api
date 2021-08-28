# shamo

## ローカルでの起動コマンド
1. docker compose build
2. docker compose run --rm api go run scripts/migrate.go
3. docker compose up -d

## prodでの起動コマンド

`docker compose -f docker-compose.prod.yml up -d`

## TODO
- t2.nanoでの起動時、メモリ不足によりmysqlのコンテナが立ち上がらない...

## TMP

```
sudo yum update -y
sudo amazon-linux-extras install docker
sudo curl -L https://github.com/docker/compose/releases/download/1.16.1/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo yum -y install git
git clone https://github.com/mi-wada/shamo_api.git
sudo usermod -a -G docker ec2-user
sudo service docker start
exit

```
