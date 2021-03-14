# TwitterAppのオーナー以外のユーザーから、ユーザー認証が必要なAPIを実行するサンプル

## これは何？
Twitterアカウントでサインインしたユーザーが、@nekoshita_yukiのツイッターアカウントをフォローするサンプル

## ユーザー認証が必要なAPIって？
たとえば、あるユーザーをフォローするAPIなど、ユーザーの認証が必要なAPI
https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/follow-search-get-users/api-reference/post-friendships-create


## 事前準備
- go 1.16.0
- [TwitterDeveloperPortal](https://developer.twitter.com/)でAppを作成しておく
- [TwitterDeveloperPortal](https://developer.twitter.com/)で対象のAppで以下を設定する
  - `App permissions`を`Read and Write`にしておく
  - `Authentication settings`を`3-legged OAuth`を`enabled`にしておく
  - `Authentication settings`で`Callback URLs`に`http://localhost:8080/twitter/callback`を設定しておく

## インストール
```
go mod vendor
```

## ローカル実行
ConsumerKeyを環境変数にセットする
```
# TwitterDeveloperPortalで、対象のAppのConsumerKeyを取得する
# https://developer.twitter.com/

$ export TWITTER_CONSUMER_KEY=xxxx
$ export TWITTER_CONSUMER_SECRET=xxxx
```

サーバーを起動する
```
$ go run main.go
```
