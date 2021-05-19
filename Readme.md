# 起動
ローカル
```
go run main.go
'''

# 設定ファイル
1. プロジェクトルートにconfig.iniファイルを作成
2. iniファイルに以下を記載し適宜valueを追加
```
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
port ポート番号
logfile ログファイルの名前
driver 使用するdbのドライバ名
name データベース名
userName データベースに接続するアカウント名
password データベースに接続するパスワード

3. 必要な記載があれば適宜config.iniに記載