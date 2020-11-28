# find_trend_project

<p align="center">
    <img src="https://drive.google.com/uc?export=view&id=1s1Eo6FlpH9lQRKwgE5qQj_Y7mDtyfEJQ">
</p>

## 実行環境

macOS Catlina  

``` command
$conda --version  
conda 4.8.3  
```

``` command
$python --version  
python 3.7.3
```

### :green_book: ライブラリ

```command
# Name      Version  
外部ライブラリ  
numpy       1.18.1  
matplotlib  3.1.3  
requests    2.23.0  
urllib3      1.25.8  
pytrends    4.7.2  


標準ライブラリ  
argparse  
time  
string  
datetime  

```

## :eyes: 使い方

これで動く（はず）

```command
$python core.py
```

これでも動く（はず）

```python
import core

core.trend_find("キーワード")
```

### core.py処理詳細

 1．予測前処理:preprocess()→preprocess.py  
  サジェストの取得：6  
  子サジェストの取得：5  
  トレンドの取得：70日分  
  トレンドの補完：0の部分を移動平均(本当に0のときもあるので保留)  
  FIFOで75日前から順に入っている  
2．予測:prediction()→get_predictor.py  
  移動平均で次の日の値を予測：7日，28日，70日の3通り行う  
3．予測後処理（予定）  
  MySQLサーバーに書き込み

### 移動平均のテスト

sin波を読み込んでコマンドを実行すると以下の動作を確認できる  

```command
$python get_predict.py
```

1. 差分の移動平均で1時刻先の予測が正しいことを確認
2. パターンを全て含んだ配列に対して，同じパターンで動く想定の予測が正しいことを確認

## ローカル

```
docker-compose up
curl -H "Content-Type: application/json" -X POST -d @test.json localhost:8080

# ワンライナー
message=$(echo "{\"message\": {\"data\": \"$(echo -n 'Zガンダム,11' | base64)\", \"messageId\": \"10\"}}") && curl -H "Content-Type: application/json" -X POST -d "${message}" localhost:8080
```
