# DDD構成

※前提知識までgrpcと共通

[toc]

## 目的

- レイヤの役割を明確化し、何をどこに書くべきか開発メンバにわかるようにする
- 疎結合なモジュールに分割されることで、テスタビリティに優れた構成にする

## 前提知識

レイヤードに関する記事は多数あるが、細かい部分で違いがある

以下の構成が似ている

- https://qiita.com/hmarf/items/7f4d39c48775c205b99b

## レイヤー構成

レイヤードがgraphql（正確にはgqlgen）と相性が良くないように思える

以下の構成となった

- `interface,graphql,infrastructure`
    - `application,domain` が `graphql` 配下の各ディレクトリに置き換わっている

### interface

- ミドルウェアやロギング処理を集めた
- `handler` は数が少ないので `main.go` に書いている

### graphql

- gqlgenから生成され、以下の役割となっている（例: account）
    - account.graphql
        - スキーマ
    - accont.resolvers.go
        - API実装
        - 主にここだけを改修していく
    - generated/generated.go
        - 自動生成 触らない
    - gqlgen.yml
        - gqlgen生成の際に参照する定義ファイル
    - model/models_gen.go
        - 自動生成 触らない
    - resolver.go
        - gRPCクライアントを書く
        - サービスが増えたら改修が必要

### infrastructure

- gPRC,redisの接続処理が書いてある

## ディレクトリ構成

```
.
├── Dockerfile
├── docker-local
│   ├── Dockerfile
│   └── fresh
│       └── runner.conf
├── graphql
│   ├── account
│   │   ├── account.graphql
│   │   ├── account.resolvers.go
│   │   ├── generated
│   │   │   └── generated.go
│   │   ├── gqlgen.yml
│   │   ├── model
│   │   │   └── models_gen.go
│   │   └── resolver.go
│   └── trend
│       ├── generated
│       │   └── generated.go
│       ├── gqlgen.yml
│       ├── model
│       │   └── models_gen.go
│       ├── resolver.go
│       ├── trend.graphql
│       └── trend.resolvers.go
├── infrastructure
│   ├── grpc
│   │   └── client
│   │       └── adaptor.go
│   └── redis
│       └── client
│           └── adaptor.go
├── interfaces
│   ├── interceptor
│   │   └── interceptor.go
│   ├── middleware
│   │   └── middleware.go
│   └── support
│       ├── x_trace_id.go
│       └── x_user_id.go
├── main.go
└── tmp
    └── runner-build
```

## API構成

- recoveryUser
    - パスワード再登録を開始する際に叩かれる
    - リクエストパラメータのemailに対応するユーザーに、有効なセッションハッシュを付与したURLをメール送付する
    - ユーザーはそのURLからパスワード変更リクエストする必要がある
- verifyUser
    - ユーザーログインの際に叩かれる
    - 有効なセッショントークンを返却する
- registerUser
    - ユーザー登録の際に叩かれる
    - 一般ユーザーは使えない（TFプロジェクト管理者用）
- updateUser
    - パスワード再登録画面で新しいパスワードを入力した際に叩かれる
    - リクエストパラメータのセッションハッシュが有効であれば、対応するユーザーのパスワードを新しいものに書き換える
    - その後自動でログインすることはさせないので、ユーザーは新しいパスワードで再ログインをする必要がある
