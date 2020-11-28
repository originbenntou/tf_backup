# 検証方法

- 正常系だけ

## gRPC

### ■account

#### RegisterUser

- ユーザー登録

```
grpcurl -import-path ./proto/account -proto user.proto -d '{"email":"verify1@gmail.com", "password":"verifyverify", "name":"山田太郎", "company_id":"1"}' -plaintext localhost:50051 account.UserService.RegisterUser
```

#### VerifyUser

- ログイン
- 存在してるユーザーのみ

```
grpcurl -import-path ./proto/account -proto user.proto -d '{"email":"verify1@gmail.com", "password":"verifyverify"}' -plaintext localhost:50051 account.UserService.VerifyUser
```

#### SendRecoverEmail

- パスワード再発行-メール送信
- 存在してるユーザーのみ

```
grpcurl -import-path ./proto/account -proto user.proto -d '{"email":"2929admin@trend-find.work", "name":"Admin"}' -plaintext localhost:50051 account.UserService.SendRecoverEmail
```

#### RecoverPassword

- パスワード再発行-発行
- トークンと認証キーは SendRecoverEmail で取得

```
grpcurl -import-path ./proto/account -proto user.proto -d '{"recoverToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTkzNDQwNjM4LCJpYXQiOiIyMDIwLTA2LTI4VDIzOjIzOjU4LjYzMjE5NjYrMDk6MDAiLCJ1c2VyIjoiYnJzYWZvZWUwYTZ0NmEyM2w1ajAifQ.5kaS7G2HFEgaxeHKbIrKA3OP6ppPVQF98QjTLnfCsPI", "authKey":"XQFOqvATwsctuKpS", "password":"gundamgundam"}' -plaintext localhost:50051 account.UserService.RecoverPassword
```

### trend

#### TrendSearch

- キーワード検索
- 存在してるユーザーのみ

```
grpcurl -import-path ./proto/trend -proto trend.proto -d '{"searchWord":"破滅の刃", "userUuid":"9bsv0s2v7s8002m4ap2g"}' -rpc-header 'x-user-uuid:9bsv0s2v7s8002m4ap2g' -plaintext localhost:50052 trend.TrendService.TrendSearch
```

#### TrendSuggest

- サジェスト表示
- 存在してるユーザーのみ
- 存在してるサジェストのみ

```
grpcurl -import-path ./proto/trend -proto trend.proto -d '{"searchId":1}' -rpc-header 'x-user-uuid:9bsv0s2v7s8002m4ap2g' -plaintext localhost:50052 trend.TrendService.TrendSuggest
```

#### TrendHistory

- 検索履歴
- 存在してるユーザーのみ

```
grpcurl -import-path ./proto/trend -proto trend.proto -d '{"userUuid":"9bsv0s2v7s8002m4ap2g"}' -rpc-header 'x-user-uuid:9bsv0s2v7s8002m4ap2g' -plaintext localhost:50052 trend.TrendService.TrendHistory
```

## GraphQL

- X-Request-Id は使われてないが、何らかの値を入れることで一連の動きをログトレースできる
- 

### ■account

#### RegisterUser

```
curl 'http://localhost:8080/account' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"mutation {registerUser(user: {email: \"2929user1@2929.co.jp\" password: \"2929password\" name: \"jhon\" companyId: 1})}"}' --compressed -H 'X-Request-Id: TEST01'
```

#### VerifyUser

- 存在してるユーザーのみ

```
curl 'http://localhost:8080/account' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"query {verifyUser(email: \"2929user1@2929.co.jp\" password: \"2929password\")}"}' --compressed -H 'X-Request-Id: TEST01'
```

##### テスト用アカウント

- 動作確認用でテストアカウントを用意した。これでログインすれば新しくユーザーを作る必要はない。

```
curl 'http://localhost:8080/account' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"query {verifyUser(email: \"2929admin@trend-find.work\" password: \"=%,u1ss;s<OtM!kL\")}"}' --compressed -H 'X-Request-Id: TEST01'
```

### SendRecoverEmail

- 存在してるユーザーのみ

```
curl 'http://localhost:8080/account' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"query {sendRecoverEmail(email: \"2929admin@trend-find.work\" name: \"Admin\")}"}' --compressed -H 'X-Request-Id: TEST01'
```

### RecoverPassword

- トークンと認証キーは SendRecoverEmail で取得

```
curl 'http://localhost:8080/account' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"mutation {recoverPassword(recoverToken: \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTkzNDQzNTI3LCJpYXQiOiIyMDIwLTA2LTI5VDAwOjEyOjA3LjY4NzkzMzMrMDk6MDAiLCJ1c2VyIjoiOWJzdjBzMnY3czgwMDJtNGFwMmcifQ.swtkRKyuv4by1wgUTQGXvUDCNGTbEL6ftC84ga59_V4\" authKey: \"CNxeoQMGyIElphdJ\" password: \"AdminAdmin\")}"}' --compressed
```

### ■trend

- VTKTにログインで払い出されたトークンを入れる

#### trendSearch

```
curl 'http://localhost:8080/trend' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"query {trendSearch(keyword: \"鬼滅の刃\")}"}' --compressed --cookie 'VTKT=373ab6df-7399-46e6-93bc-23e1357e7389' -H 'X-Request-Id: TEST01'
```

#### trendSearch

```
curl 'http://localhost:8080/trend' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"query {trendSuggest(suggestId: 1) {keyword childSuggests {word growth {short medium long} graphs {short {date value} medium {date value} long {date value}}}}}"}' --compressed --cookie 'VTKT=373ab6df-7399-46e6-93bc-23e1357e7389' -H 'X-Request-Id: TEST01'
```

#### trendHistory

```
curl 'http://localhost:8080/trend' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"query {trendHistory {suggestId status}}"}' --compressed --cookie 'VTKT=373ab6df-7399-46e6-93bc-23e1357e7389' -H 'X-Request-Id: TEST01'
```

## redis

やり方をよく忘れるので
https://github.com/TrendFindProject/tf_backend/pull/1
