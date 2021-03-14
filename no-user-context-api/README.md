# ユーザー認証が不要なAPIを実行するサンプル

## これは何？
@nekoshita_yukiのタイムラインを取得するサンプル

## ユーザー認証が不要なAPIって？
たとえば、あるユーザーのタイムラインを取得するAPIなど、もともとTwitter上で公開されてるようなGET系
https://developer.twitter.com/en/docs/twitter-api/v1/tweets/timelines/api-reference/get-statuses-user_timeline


## 事前準備
- go 1.16.0
- [TwitterDeveloperPortal](https://developer.twitter.com/)でAppを作成しておく

## インストール
```
go mod vendor
```

## ローカル実行
ConsumerKeysとAccessKeys環境変数にセットする
```
# TwitterDeveloperPortalで、対象のAppのConsumerKeysとAccessTokensを取得する
# https://developer.twitter.com/

$ export TWITTER_CONSUMER_KEY=xxxx
$ export TWITTER_CONSUMER_SECRET=xxxx
```

タイムラインを取得するAPIをたたく
```
$ go run main.go
```
