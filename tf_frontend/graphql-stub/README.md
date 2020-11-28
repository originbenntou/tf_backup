# GraphQL スタブ

## モックサーバー環境構築

1. `npm ci`
2. `npm run start`

GraphQL クエリの実行  
http://localhost:4000/graphql にアクセス

サンプルリクエストクエリ

```
query {
  trendSuggest(suggestId: 99) {
    keyword,
    childSuggests {
      word,
      growth {
        short
        medium
        long
      }
      graphs {
        short {
          date
          value
        }
        medium {
          date
          value
        }
        long {
          date
          value
        }
      }
    }
  }
}
```
