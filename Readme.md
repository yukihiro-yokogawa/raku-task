# 起動
・ローカル起動について
1. プロジェクトrootにあるdockerディレクトリ内に.envファイルを作成
```
.env

MARIADB_ROOT_PASSWORD=
MARIADB_DATABASE=
MARIADB_USER=
MARIADB_PASSWORD=
```
2. .envにvalueを入力後、dockerディレクトリ内で以下を実行（既に実行済みの場合コンテナの起動を確認すること）
```
docker-compose up -d
```
3. プロジェクトルートにconfig.iniファイルを作成
4. iniファイルに以下を記載し適宜valueを追加
```
config.ini

[web]
port = 8080
logfile = logging.log

[db]
driver = 
name = 
userName = 
password = 
protocol = 
```
5. プロジェクトルートで以下を実行
```
go run main.go
```