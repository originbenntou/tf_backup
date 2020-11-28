# DDD構成

※前提知識までgraphqlと共通

[toc]

## 目的

- レイヤの役割を明確化し、何をどこに書くべきか開発メンバにわかるようにする
- 疎結合なモジュールに分割されることで、テスタビリティに優れた構成にする

## 前提知識

レイヤードに関する記事は多数あるが、細かい部分で違いがある

以下の構成が似ている

- https://qiita.com/hmarf/items/7f4d39c48775c205b99b

## レイヤー構成

`interface層,application層,domain層,infrastructure層` の4層で構成される

### interface

- `handler` のみ
- クライアントからのリクエストに応じて各サービス(user,company,plan)の呼び出しを行う
- routingはgRPCに任せている(`RegisterUserServiceServer`)

### application

- `sevice` のみ
- サービスの処理全体をコントロールする（MVCだとControllerの立ち位置）
- 各ドメイン(user,company,plan)の処理を適宜呼び出して、サービスとして振る舞いを記述する
    - 例えばユーザー登録をする場合、userテーブルとcompanyテーブルにそれぞれ干渉するので適宜ドメインを呼び出して処理をする

### domain

- `domain層` は `service,repository,model` で構成されるのが基本だが `service` を使うほど複雑な処理はやらないため省略している
    - 本来 `domain - service` が `controller` の役割を果たすが、それだと `application層` が薄くなりすぎるのでやめた。サービスの規模が大きくひとつのリクエストで多くの `domain` に干渉するようにあれば必要になるのだろう。
- `domain - repository` はテーブル単位で分割されている

### interface

- `datastore` のみ
    - DB操作が記述される

## ディレクトリ構成

```
./
├── Dockerfile
├── application
│   └── service
│       └── user_service.go
├── constant
│   └── mysql.go
├── docker-local // ローカル用のDockerfileを作成
│   ├── Dockerfile
│   └── fresh // ホットリロード用
│       └── runner.conf
├── domain
│   ├── model
│   │   ├── company.go
│   │   ├── session.go
│   │   └── user.go
│   └── repository
│       ├── company_repository.go
│       ├── plan_repository.go
│       ├── session_repository.go
│       └── user_repository.go
├── infrastructure
│   └── datastore
│       ├── company_repository.go
│       ├── plan_repository.go
│       ├── session_repository.go
│       └── user_repository.go
├── interfaces
│   └── handler
│       └── handler.go
├── main.go
├── registry
│   ├── container
│   │   ├── container.go
│   │   ├── handler_container.go
│   │   ├── repository_container.go
│   │   └── service_container.go
│   └── registry.go
└── tmp
    └── runner-build
```

## サービス構成

以下のサービス群がアカウントサービスとして提供される

### UserService

- RegisterUser
    - ユーザー登録
    - リクエストパラメータのEmail,会社情報を元に存在確認をクリアしたのちに,パスワードをハッシュ化してデータベースに保存する
- VerifyUser
    - ユーザーログイン
    - リクエストパラメータのEmail,パスワード情報がデータベースの情報と一致すればセッショントークンとユーザー情報（パスワード以外）を返却する
        - ユーザー情報返却はGatewayでRedis（ユーザーセッションキャッシュ）に登録するために必要
    - Redis登録とは別にMySQLでもセッション管理を行う
        - 同一会社のプラン上限ログインチェック
            - プラン上限ログインの場合は403
        - 同一ユーザーの多重ログインチェック
            - 多重ログインの場合は、古い有効セッションを削除

### CompanyServide

未実装（会社登録画面をフロントから提供するようになればつくるかも）

### PlanService

未実装（プラン登録画面をフロントから提供するようになればつくるかも）

## DI

考え方は以下を参考

- https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/

※つくり方は全然違っている
