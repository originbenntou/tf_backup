# TREND FIND PROJECT schema

## フォルダ構成

```
.
├── README.md
├── doc
└── graphql
    ├── account.graphql // ユーザー認証系API
    └── trend.graphql   // トレンドAPI
```

## GraphQL

### ドキュメント

※index.htmlをブラウザ閲覧

- [account](https://github.com/TrendFindProject/tf_schema/blob/master/doc/graphql/account/index.html)
- [trend](https://github.com/TrendFindProject/tf_schema/blob/master/doc/graphql/trend/index.html)

#### 応答例

- `data` にスキーマ定義上の返り値が入る
- `extensions - code` にAPIの応答値が入る
- httpステータスコードはGraphQLサーバーと疎通していれば200が返る

```
# verifyUser

## 正常系
{
  "data": {
    "verifyUser": "359b1e33-c626-4e95-880a-c2fbad98d9e2"
  },
  "extensions": {
    "code": 0
  }
}

## 異常系
{
  "errors": [
    {
      "message": "rpc error: code = NotFound desc = user is not found: 2929user2@2929.co.jp",
      "path": [
        "verifyUser"
      ]
    }
  ],
  "data": null,
  "extensions": {
    "code": 5
  }
}
```

### スキーマ更新

- `./graphql` 配下のスキーマ定義を変更
- フロントチーム、バックエンドチーム双方のチェックを受ける
- ドキュメント更新を実施
    - [graphdoc](https://github.com/2fd/graphdoc#readm) を利用

```
# インストール済みなら不要
$ npm install -g @2fd/graphdoc

# account
$ graphdoc -s ./graphql/account.graphql -o ./doc/graphql/account --force
# trend
$ graphdoc -s ./graphql/trend.graphql -o ./doc/graphql/trend --force
```
