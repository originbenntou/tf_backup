# -*- coding:utf-8 -*-
"""
傾向の後追い処理をまとめたモジュール  
詳細は次の通り  
- ベースにするデータの取得：検索ワードの既定期間の2倍前(150-0日前，50-0日前，14-0日前)のトレンドを取得
- 後追いするデータの取得：子サジェストワードの規定期間の2倍前(取得期間+予測期間)トレンドを取得
- トレースの比率を取得：検索ワードの前半分と子サジェストワードの前半分を比較して差分を比率として返す
- トレースする：検索ワードの後半に距離に応じた比率をかけて，子サジェストワードの後半に加算する
"""


# 比率のトレースを取得
def get_trace_degree(base, target, hyper_parameter:float=1.0):
    """
    動作  
    2つのリストの差分をとって比率として返す  
    距離が近ければtrace_dgreeは大きい（最大1）  
    距離が遠ければtrace_dgreeは小さい（最小0）  
    引数  
    - base：元のリスト  
    - target：比較するリスト  
    - hyper_parameter：調整可能なパラメータ，トレースの信頼値のようなもの，デフォルト；0.3
    """
    
    ## 差分をとって足し合わせる
    degree_seed = sum([abs(target[num] - base[num]) for num in range(len(target))])
    # print("平均絶対距離 :", degree_seed/len(base))
    
    ## 長さで割って1点あたりの比率としておく
    ## 長さ：平均，100：差分の値域は[-100, 100]
    ### 距離が近ければtrace_dgreeは大きい
    ### 距離が遠ければtrace_dgreeは小さい
    trace_degree = 1-(degree_seed/len(base)/100)
    if trace_degree>=1: # degree_seed/len(base)がマイナスになったとき（ありえない）
        trace_degree = 0
    elif trace_degree<0:# degree_seed/len(base)が100を超えたとき
        trace_degree = 0


    ## 最後にハイパーパラメータでチューニング
    trace_degree *= hyper_parameter


    return trace_degree


# トレースする
def get_traced_value(base, target, trace_degree, hyper_parameter=1.0):
    """
    動作  
    トレース対象とトレースの比率を元にトレース量を算出．元のリストにトレース量を加算する  
    引数  
    - base：元のリスト
    - target：トレースするリスト
    - trace_degree：トレースの比率  
    - hyper_parameter：チューニング可能な適当な値
    """

    ## トレース量を算出
    target = list(map(lambda x: int(x * hyper_parameter), target))

    ## トレース実行
    ## (1-トーレス度合い)*元の値 + トレース度合い*トレース
    traced_base = [ int((1-trace_degree)*base[num] + trace_degree*target[num]) for num in range(len(base))]
    

    return traced_base


# テスト
def test_tracing():
    """
    動作  
    トレースできるかどうかを試すだけ  
    """
    import get_trend
    gt = get_trend.GetTrend()
    results = gt.get_trend_on_day(["zoom","python"], 51)


    search_word = results["zoom"].values
    child_suggest = results["python"].values


    traced_base = trace_trend(child_suggest, search_word)


    import matplotlib.pyplot as plt
    plt.plot(traced_base, label="traced")
    plt.plot(child_suggest, label="child_suggest")
    plt.plot(search_word, label="search_word")
    plt.legend()
    plt.show()


# 傾向の後追いをする
def trace_trend(base, target, hyper_parameter=1.0):
    """  
    動作  
    baseとtargetの前半の距離をもとにbaseの値をtargetの後半の値に近づける  
    引数  
    全てFIFO形式で後半に新しい値が入っている  
    - base：子サジェストデータのリスト，実測値で構成されている  
    - target：検索ワードのリスト，前半も後半も実測値  
    - hyperparameter：外側から調整可能なパラメータ，影響を与える度合いを調整できる，デフォルトは1.0  
    """
    length_ = int(len(base))

    ## トレース度合い取得
    trace_degree = get_trace_degree(base, target[:length_])
    

    ## トレース値の取得
    traced_value = get_traced_value(base, target[length_+1:], trace_degree, hyper_parameter)


    return traced_value


if __name__ == "__main__":
    test_tracing()
    print("hoge")



