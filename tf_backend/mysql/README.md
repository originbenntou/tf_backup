# マイグレーション

https://github.com/golang-migrate/migrate

## local

makeで完結

### マイグレーションファイル生成

```
make migrate_init TITLE=alter_user_table
# tf_backend/mysql/migrations配下にバージョニングされたファイルができる
```

### バージョンアップ

```
make migrate_up SVC=account VERSION=000001
```

### 全バージョン

```
make migrate_up SVC=account
```

### バージョンダウン

```
make migrate_down SVC=account VERSION=000001
```

## develop（skaffold）

init後にjobを実行

```

```

## production

```

```
