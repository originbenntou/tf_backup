# tf_infra
## 検証環境構築
### 事前作業
- Docker For MacのインストールとKubernetes有効化
- Skaffold　インストール
```
brew install skaffold
```

### 環境起動
- tf_infraディレクトリにいらっしゃいませ
```
ディレクトリ構成は以下であることを想定しています。
- tf_backend
- tf_frontend
- tf_infra
```
- privateリポジトリにフルアクセスできるGITHUBアクセストークンをセット
```
export GITHUB_ACCESS_TOKEN=アクセストークン

# 永続化させる場合は .bash_profile に書き込んでください
export GITHUB_ACCESS_TOKEN=アクセストークン
```

- Skaffoldの実行 ※これ以降はファイル監視して自動でビルド走ります
```
skaffold dev
```

## 検証環境について
### 画面URL
    + フロント: `http://localhost:30080/`
    + PlayGround: `http://localhost:30880/` 

### Database接続先
    + MySQL: localhost:30306
    + Redis: localhost:30379

# 本番環境構築作業メモ
## 下準備　※GCP作業
現在は手作業で行います。ごめんなさい。

1. Google Cloud SDKが必要(´Д⊂ｸﾞｽﾝ
2. GCPプロジェクトの作成
    - プロジェクト名:　trend-find（予定） かぶらなければ、プロジェクトIDも同一になる 
3. APIの有効化
    - CloudSQLAdmin
    - CloudBuildAPI
    - KubernetesAPI
4. CloudSQLのインスタンス作成
    - CloudSQLに接続して、データベースの初期化などは必要です
5. MemoryStore For Redisのインスタンス作成
6. IAMでCloud Build サービス アカウントに,Kubernetes Engine 管理者を追加する
7. GCloudコマンドを確認する
    ```
    gcloud config set project 【プロジェクトID】
    gcloud config list
    ```
8. 静的IPを取得する IP名: trend-find-ip
    ```
    gcloud compute addresses create trend-find-ip --global
    ```
9. Domain取得 ※未確定
    ```
    trend-find.work
    ```
10. GoogleManagedSSL取得 SSL名: trend-find-ssl
    ```
    gcloud beta compute ssl-certificates create trend-find-ssl --domains trend-find.work
    ``` 

## CloudBuild
Account/Gateway/Front/AlgorithmのそれぞれでCloudBuildのマニフェストが必要です

## GKE作業 
1. Cluster作成 クラスタ名: trend-find-cluster　num-nodes, machine-typeは要調整
    ```
    gcloud container clusters create trend-find-cluster --num-nodes=1 --zone=asia-northeast1-a --machine-type=g1-small --disk-size 20 --scopes=gke-default,sql-admin,storage-rw
    ```
2. DefaultClusterのセット　※ご自身の環境用
    ```
    gcloud container clusters get-credentials trend-find-cluster --zone=asia-northeast1-a 
    gcloud config set container/cluster trend-find-cluster
    gcloud config list
    ```
3. CloudSQLProxy用にSecretを作成しておく（credentialsファイルはサービスアカウントから落とせる） 
    ```
    kubectl create secret generic cloudsql-instance-credentials --from-file /Users/mike.kodama/google-cloud/trend-find.json
    ```

4. さぁ！Apply！！
```
kubectl apply -f k8s/production
```

## マイグレーション
```
kubectl apply -f k8s/develop/job/db-migrate-up.yaml
```
