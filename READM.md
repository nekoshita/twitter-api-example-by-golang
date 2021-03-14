# twitter api example

# これはなに？
TwitterのAPIをたたくサンプル実装

# TwitterのAPI叩くためには何が必要？
- [TwitterDeveloperPortal](https://developer.twitter.com/)への登録
- [TwitterDeveloperPortal](https://developer.twitter.com/)でAppを作成しておく

# Twitter APIの公式ドキュメントは？
https://developer.twitter.com/en/docs/twitter-api/v1

# Twitter APIの認証について
- 全てのAPIをたたくには、[TwitterDeveloperPortal](https://developer.twitter.com/)で作成したAppの `Consmer Keys` による認証が必要
- 特定のユーザーに対するアクション（フォローするなど）などのAPIを叩く場合は、`Consmer Keys` に加え、Twitter Userの`Access Token`が必要
- Twitter Userの`Access Token`の取得方法
  - [TwitterDeveloperPortal](https://developer.twitter.com/)のAppのオーナーであるTwitter Userの`Access Token`は[TwitterDeveloperPortal](https://developer.twitter.com/)上で`Access Token`の取得可能
  - それ以外のTwitter Userの`Access Token`を取得するには、[Twitter Sign-in](https://developer.twitter.com/en/docs/authentication/guides/log-in-with-twitter)による認証をする必要がある

# サンプルの種類
- [no-user-context-api](./no-user-context-api/README.md)
  - ユーザー認証が不要なAPIをたたくサンプル
  - ツイートを検索したり、ユーザーのタイムラインを取得するだけならこれ
- [user-context-api-by-app-owner](./user-context-api-by-app-owner/README.md)
  - ユーザー認証が必要なAPIをたたくサンプル
  - Twitter Appのオーナーのユーザーを認証する前提なので、Access Tokensは[TwitterDeveloperPortal](https://developer.twitter.com/)で発行する前提
  - 不特定多数のTwitter Userに`Twitter Sign-in`を要求しないが、ユーザー認証が必要なAPIを叩く場合はこれ
- [user-context-api-by-any-user](./user-context-api-by-any-user/README.md)
  - ユーザー認証が必要なAPIをたたくサンプル
  - 不特定多数のTwitter Userに`Twitter Sign-in`を要求し、それぞれのTwitter Userの`Access Tokens`を使ってユーザー認証が必要なAPIを叩く場合はこれ
  - たとえば、`Twitter Sign-in`に対して、フォローするAPIをたたいてあげる、みたいなことやる場合はこれ

# 使用したライブラリ
Twitter APIを使えるライブラリ一覧はこちら
https://developer.twitter.com/en/docs/twitter-api/tools-and-libraries

今回はGo言語のライブラリを使いました
https://github.com/dghubble/go-twitter

# 機密情報の取り扱いについて
- 機密情報とはTwitter Userの``Access Tokens`やTwitter Appの`Consumer Key`
- これはサンプルなので、わりとてきとうに扱ってるので、本番で使う場合は注意してください
- Twitterも公式ドキュメントで機密情報の取り扱いに関して、注意しています
  - https://developer.twitter.com/ja/docs/basics/authentication/guides/securing-keys-and-tokens
