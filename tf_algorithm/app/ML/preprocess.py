# -*- coding:utf-8 -*-
"""
予測前処理をまとめたモジュール  
詳細は次の通り  
- サジェストの取得：6
- 子サジェストの取得：5
- トレンドの取得：75日分→7日，25日，75日と分配する
- トレンドの補完：今回はなし
"""

from . import get_suggest
from . import get_trend


def preprocess(phrase:str, day_length:int=75):
    """
    動作
    引数
    """
    # サジェストの取得：6
    # 子サジェストの取得：5
    dict_child_suggest_words = get_suggest.get_google_suggest(phrase)


    # トレンドの取得：75日分
    ## initialize
    dict_child_trend = {}
    gt = get_trend.GetTrend()
    ### トレンド数値が獲得出来なかったときの0埋め
    trend_failed = [0]*day_length

    ## サジェストワードのループ
    for child_suggests in dict_child_suggest_words.values():
        ### 子サジェストワードのループ
        for child_suggest in child_suggests:
            temp_ = gt.get_trend_on_day([child_suggest], day_length+1)
            ### 不要な部分の削除
            try:
                ### トレンドが取得出来ている場合の処理
                del temp_["isPartial"]
                dict_child_trend[child_suggest] = [val[0] for val in temp_.values.tolist()]
            except:
                ### トレンドが取得出来なかった場合の処理
                dict_child_trend[child_suggest] = trend_failed

    return dict_child_trend, list(dict_child_suggest_words.keys())


if __name__=="__main__":
    """
    test code
    """
    # サジェストの取得：6
    # 子サジェストの取得：5
    print("get suggest...")
    dict_child_suggest_words = get_suggest.get_google_suggest("python")
    print("done")

    
    # トレンドの取得：75日分
    print("get trend...")
    ## initialize
    dict_child_trend = {}
    gt = get_trend.GetTrend()
    ### トレンド数値が獲得出来なかったときの0埋め
    trend_failed = [0]*75

    ## サジェストワードのループ
    for child_suggests in dict_child_suggest_words.values():
        ### 子サジェストワードのループ
        for child_suggest in child_suggests:
            temp_ = gt.get_trend_on_day([child_suggest], 75)
            ### 不要な部分の削除
            try:
                ### トレンドが取得出来ている場合の処理
                del temp_["isPartial"]
                dict_child_trend[child_suggest] = [val[0] for val in temp_.values.tolist()]
            except:
                ### トレンドが取得出来なかった場合の処理
                print("failed to get trend value")
                dict_child_trend[child_suggest] = trend_failed


    # 取得トレンド数が足らなかったときの処理
    ## サーバー書き込み時に対応することに

    # トレンドの補完：0の部分を移動平均
    ## 0のインデックスをカウントして概形を取得
    ## 前方向と後ろ方向に移動平均
    print("complement")
    print("done")
