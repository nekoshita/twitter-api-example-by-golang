# TwitterAppのオーナーのユーザーから、ユーザー認証が必要なAPIを実行するサンプル

## これは何？
TwitterAppのオーナーのユーザーが、@nekoshita_yukiのツイッターアカウントをフォローするサンプル

## ユーザー認証が必要なAPIって？
たとえば、あるユーザーをフォローするAPIなど、ユーザーの認証が必要なAPI
https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/follow-search-get-users/api-reference/post-friendships-create


## 事前準備
- go 1.16.0
- [TwitterDeveloperPortal](https://developer.twitter.com/)でAppを作成しておく
- [TwitterDeveloperPortal](https://developer.twitter.com/)で対象のAppで以下を設定する
  - `App permissions`を`Read and Write`にしておく

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
$ export TWITTER_USER_ACCESS_TOKEN=xxxx
$ export TWITTER_USER_ACCESS_SECRET=xxxx
```

フォローするAPIをたたく
```
$ go run main.go
```
