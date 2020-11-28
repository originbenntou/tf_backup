# TREND FIND PROJECT backend

## 環境構築

### 前提

- Go 1.14
- gRPC, Protocol Buffers
- GraphQL

#### 技術周り

- https://github.com/TrendFindProject/tf_backend/tree/master/doc

### 起動

```
# clone
cd $GOPATH/src/github.com/TrendFindProject
git clone git@github.com:TrendFindProject/tf_backend.git

# ビルド用GITHUBアクセストークンをセット
export GITHUB_ACCESS_TOKEN=repoにフルアクセス権限を持つトークン

# コンテナ起動
docker-compose build --no-cache
docker-compose up
```

### SendGrid

docker-compose起動時に環境変数へセット
キーは[ここ](https://tsuku-tsuku.slack.com/archives/CTXKRDUSY/p1593090154167300?thread_ts=1593083483.158900&cid=CTXKRDUSY)

```
export SEND_GRID_API_KEY={APIキー}
```

### JWT

キーは[ここ]()
※商用の格納場所はまだ決めてない

```
# ローカル
export JWT_SECRET_KEY=XQznco6b+dmIIkHR0+8ibWlh3iOW1kRH2WtYNx2HC7Ot

## 新規作成
openssl rand -base64 30
```

## 動作確認

- [GraphQL playground](http://localhost:8080)

## 検証

- https://github.com/TrendFindProject/tf_backend/blob/master/doc/testcase.md

## 改修方法

### protoファイル更新

#### Go

```
make gen_go_proto SVC=account
make gen_go_proto SVC=trend
```

#### Python

```
make gen_py_proto SVC=trend
```

### graphqlスキーマ更新

※graphqlスキーマは[tf_schema](https://github.com/TrendFindProject/tf_schema)から持ってきてコピペ

- [account](https://github.com/TrendFindProject/tf_backend/blob/master/gateway/graphql/account/account.graphql)
- [trend](https://github.com/TrendFindProject/tf_backend/blob/master/gateway/graphql/trend/trend.graphql)

```
SERVICE=account
SERVICE=trend

cd ./gateway/graphql/$SERVICE
go run github.com/99designs/gqlgen generate
```

## スキーマ

### GraphQL

- [schema](https://github.com/TrendFindProject/tf_schema)

### gRPC

- [proto](https://github.com/TrendFindProject/tf_backend/tree/master/proto)

### DB

- [account](https://github.com/TrendFindProject/trendfindproject/blob/master/backend/mysql/init/00.account.sql)
- [trend](https://github.com/TrendFindProject/trendfindproject/blob/master/backend/mysql/init/00.trend.sql)

## テスト用アカウント

```
2929admin@trend-find.work
=%,u1ss;s<OtM!kL
```
