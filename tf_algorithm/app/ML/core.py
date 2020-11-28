# -*- coding:utf-8 -*-

import copy
import datetime
import time

from . import get_trace
from . import get_trend
from . import preprocess
from .my_functions import *


def trend_find(search_word, save_output=False):
    """
    1．予測前処理:preprocess()→preprocess.py
    サジェストの取得：6
    子サジェストの取得：5
    トレンドの取得：150日分；前75日はトレース，後75日は予測に使う
    トレンドの補完：0の部分を移動平均(本当に0のときもあるので保留)
    FIFOで75日前から順に入っている

    2．予測:prediction()→get_predictor.py
    移動平均で次の日の値を予測：7日，25日，75日の3通り行う

    3．予測後処理（予定）
    MySQLサーバーに書き込み
    """

    # 1．予測前処理:preprocess()→preprocess.py
    ## 子サジェストワードのトレンド値の取得
    result, suggest_words = preprocess.preprocess(search_word, day_length=70)
    
    ## 検索ワードのトレンド値の取得
    search_trend = get_trend.GetTrend().get_trend_on_day([search_word], day_term=141)
    ### 余計なものを削除する
    del search_trend["isPartial"]
    search_trend = [val[0] for val in search_trend.values.tolist()]
    ### grad算出
    search_trend = get_difference_in(search_trend)


    list_result  = []
    dict_suggest = {}
    list_child_suggests = []
    dict_child_suggest  = {}
    # 2.子サジェストのgrad算出
    current_trend = {} # 最新のトレンド値
    for index, key in enumerate(result.keys()):
        dict_term = {}
        current_trend[key] = result[key][-1]
        grad = get_difference_in(result[key])

        short_ = get_trace.trace_trend(grad[-(7-1):], search_trend[-((2*7)-1):])
        middle_ = get_trace.trace_trend(grad[-(28-1):], search_trend[-((2*28)-1):])
        long_ = get_trace.trace_trend(grad, search_trend)
        # 4.予測
        ans = difference_to_value(start=current_trend[key], difference=short_)
        ## 2ターム追加
        ans[0:0] = result[key][-3:]
        # ans[0:0] = result[key][-8:] # test code
        # get_predictor.show_graph(ans, name="short.png")
        # 結果をまとめる
        list_short = summerize_as_dict(ans,"short")

        ## 中期：同上
        ans = difference_to_value(start=current_trend[key], difference=middle_)
        ans[0:0] = result[key][-9:]
        # ans[0:0] = result[key][-29:] # test code
        # get_predictor.show_graph(ans, name="middle.png") # test code
        list_middle = summerize_as_dict(ans, "medium")
        ## 長期：同上
        ans = difference_to_value(start=current_trend[key], difference=long_)
        ans[0:0] = result[key][-21:]
        # ans[0:0] = result[key][-71:] # test code
        # get_predictor.show_graph(ans, name="long.png") # test code
        list_long = summerize_as_dict(ans, "long")


        # 方向の推定
        dict_growth={}
        dict_growth["short"] = estimate_direction(short_)
        dict_growth["medium"] = estimate_direction(middle_)
        dict_growth["long"] = estimate_direction(long_)

        ## 戻り値の作成
        ### 値の準備
        dict_term["short"] = copy.deepcopy(list_short); dict_term["medium"] = copy.deepcopy(list_middle); dict_term["long"] = copy.deepcopy(list_long)
        ### 検索ワード＋サジェストワードの準備
        suggest_word = suggest_words[int(index/6)]
        ### 子サジェストの辞書
        dict_child_suggest["word"]   = key.replace(suggest_word,"")
        dict_child_suggest["growth"] = dict_growth.copy()
        dict_child_suggest["graphs"]  = dict_term.copy()
        ### 一つの配列に子サジェストの辞書をまとめる
        list_child_suggests.append(dict_child_suggest.copy())



        ### サジェストワードごとに子サジェストワードの値をまとめる
        if (index+1)%6==0:
            dict_suggest["keyword"] = suggest_word.replace(search_word,"")
            dict_suggest["childSuggests"] = copy.deepcopy(list_child_suggests)
            list_result.append(dict_suggest.copy())
            ### 子サジェストの配列のリセット
            list_child_suggests = []
            ### サジェストの辞書のリセット
            dict_suggest = {}

        ## 子サジェストの辞書のリセット
        dict_child_suggest = {}

    import json
    if save_output:
        with open("test_with_term.json",mode="w")as F:
            json.dump(list_result, F, indent=4, ensure_ascii=False)
            
    # print(json.dumps(list_result, indent=4, ensure_ascii=False))
    return json.dumps(list_result, indent=4, ensure_ascii=False)





if __name__ == "__main__":
    """
    実数値：予測の計算，MA算出
    微分値：トレース，矢印

    実数値：grad算出，トレース，予測の算出
    微分値：
    矢印未実装

    1．予測前処理:preprocess()→preprocess.py
    1．1．ワードの取得  
    サジェストの取得：6  
    子サジェストの取得：6  
    1．2．トレンド値の取得  
    子サジェストは70日分，検索ワードは140日分；前70日はトレース；後70日は予測に使う  
    FIFOで70日前(140日前)から順に入っている  

    2．子サジェストのgrad算出
    子サジェストに関してはgrad値と基準の日のデータがあれば全てのパターンを再現できる  

    3. トレース  

    4．予測:prediction()→get_predictor.py
    移動平均で次の日の値を予測：7日，28日，70日の3通り行う

    5．予測後処理
    json形式の指定のフォーマットで返せるように加工
    """

    print(trend_find("zoom"));exit()
    start = time.time()
    
    now_ = str(datetime.datetime.today()).split(" ")
    now_ = now_[0].replace("-", "")
    
    seach_word = "python"
    

    # 1．予測前処理:preprocess()→preprocess.py
    ## 子サジェストワードのトレンド値の取得
    result, suggest_words = preprocess.preprocess(seach_word, day_length=70)
    
    ## 検索ワードのトレンド値の取得
    search_trend = get_trend.GetTrend().get_trend_on_day([seach_word], day_term=141)
    ### 余計なものを削除する
    del search_trend["isPartial"]
    search_trend = [val[0] for val in search_trend.values.tolist()]
    ### grad算出
    search_trend = get_difference_in(search_trend)


    list_result  = []
    dict_suggest = {}
    list_child_suggests = []
    dict_child_suggest  = {}
    # 2.子サジェストのgrad算出
    current_trend = {} # 最新のトレンド値
    for index, key in enumerate(result.keys()):
        dict_term = {}
        current_trend[key] = result[key][-1]
        grad = get_difference_in(result[key])

        short_ = get_trace.trace_trend(grad[-(7-1):], search_trend[-((2*7)-1):])
        middle_ = get_trace.trace_trend(grad[-(28-1):], search_trend[-((2*28)-1):])
        long_ = get_trace.trace_trend(grad, search_trend)
        # 4.予測
        ans = difference_to_value(start=current_trend[key], difference=short_)
        ## 2ターム追加
        ans[0:0] = result[key][-3:]
        # 結果をまとめる
        list_short = summerize_as_dict(ans,"short")

        ## 中期：同上
        ans = difference_to_value(start=current_trend[key], difference=middle_)
        ans[0:0] = result[key][-9:]
        list_middle = summerize_as_dict(ans,"medium")
        ## 長期：同上
        ans = difference_to_value(start=current_trend[key], difference=long_)
        ans[0:0] = result[key][-21:]
        list_long = summerize_as_dict(ans,"long")

        # 方向の推定
        dict_growth={}
        dict_growth["short"] = estimate_direction(short_)
        dict_growth["medium"] = estimate_direction(middle_)
        dict_growth["long"] = estimate_direction(long_)

        ## 戻り値の作成
        ### 値の準備
        dict_term["short"] = copy.deepcopy(list_short); dict_term["medium"] = copy.deepcopy(list_middle); dict_term["long"] = copy.deepcopy(list_long)
        ### 検索ワード＋サジェストワードの準備
        suggest_word = suggest_words[int(index/6)]
        ### 子サジェストの辞書
        dict_child_suggest["word"]   = key.replace(suggest_word,"")
        dict_child_suggest["growth"] = dict_growth.copy()
        dict_child_suggest["graphs"]  = dict_term.copy()
        ### 一つの配列に子サジェストの辞書をまとめる
        list_child_suggests.append(dict_child_suggest.copy())



        ### サジェストワードごとに子サジェストワードの値をまとめる
        if (index+1)%6==0:
            dict_suggest["keyword"] = suggest_word.replace("コロナ","")
            dict_suggest["childSuggests"] = copy.deepcopy(list_child_suggests)
            list_result.append(dict_suggest.copy())
            ### 子サジェストの配列のリセット
            list_child_suggests = []
            ### サジェストの辞書のリセット
            dict_suggest = {}

        ## 子サジェストの辞書のリセット
        dict_child_suggest = {}

        import json
        with open("test_with_term.json",mode="w")as F:
            json.dump(list_result, F, indent=4, ensure_ascii=False)
        # print(json.dumps(list_result, indent=4, ensure_ascii=False))

    exit()
    end = time.time()
    print("Process time :", end-start)


