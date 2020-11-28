"""
汎用的に使うモジュール
"""
# -*- coding:utf-8 -*-
import copy
import datetime

import numpy as np


# 変化率算出
def get_difference_in(list_, scaling_factor=1, input_type="value"):
    """
    # 動作  
    （後の値との差分/現在値）で変化率を算出する．※1次元でなければ0を返す．  
    戻り値の配列の長さが入力時の長さより1短くなる．  

    # 引数  
    list_：ターゲット  
    scaling_factor：スケーリング係数，デフォルトは1  
    type：valueかrateを指定返す，デフォルトはvalue  
    """

    if np.ndim(list_)==1:
        if input_type=="rate":
            return [(list_[num+1] - list_[num])/list_[num]/scaling_factor for num in range(len(list_)-1)]
        elif input_type=="value":
            return [(list_[num+1] - list_[num])/scaling_factor for num in range(len(list_)-1)]
        else:
            return 0
    else:
        return 0


# [0,1]正規化
def normalization(list_):
    """
    # 動作
    [0,1]になるように正規化する．
    # 引数

    """

    # 最大値と最小値の取得
    max_val = max(list_); min_val = min(list_)


    # 最小値をバイアスとして取り除き，最大値で現在値を割る
    if np.ndim(list_)==1:
        return [(value -min_val)/(max_val-min_val) for value in list_]
    else:
        return 0


# データの詳細表示
def data_detail(list_, name="", data_=False):
    print("---",name,"---")
    print("length :",len(list_))
    print("type :", type(list_))
    # print("data shape :", np.shape(list_))
    if data_: print("data :", list_)


# 平均プーリング
def average_pooling(list_, stride=10, k_size=20, scaling_factor=1, max_itr=50):
    """
    動作  
    与えられたリストを平均プーリングする  
    ちなみに，stride,max_itrを1にするとk_sizeのカーネルを一度だけフィルタリングする動作をする  
    引数  
    - list_：対象のリスト，型：list，デフォルト；なし
    - stride：ストライド幅，，型：int，デフォルト：10
    - k_size：カーネルサイズ，型：int，デフォルト：20
    - scaling_factor：係数，型：int，デフォルト：1
    - max_itr：ループ回数，型：int，デフォルト：50
    """
    itr = 0
    arr_ = []


    while(itr<max_itr):
        tem = list_[itr*stride:(itr+1)*stride+k_size]
        arr_.append( sum(tem)/k_size/scaling_factor )
        itr+=1


    return arr_



# 数値のリストを日付をkeyに，値をvalueにした辞書に変換
def summerize_as_dict(ans:list, day_term:str):
    list_ = []
    if   day_term=="short" : day_back = 2
    elif day_term=="medium": day_back = 8
    elif day_term=="long"  : day_back = 20
    else : day_back = 0
    for index_, val in enumerate(ans):
        dict_ = {}
        predict_day = str(datetime.date.today()+datetime.timedelta( days=(index_ - day_back) )).replace("-","")
        dict_["date"] = predict_day
        dict_["value"]= val
        list_.append(dict_.copy())
    return list_


# トレンド方向の推定
def estimate_direction(target:list):
    if abs(sum(target))<20:
        direction = "FLAT"
    elif sum(target)<0:
        direction = "DOWN"
    elif sum(target)>0:
        direction = "UP"
    return direction


# 差分を実測値へ変換
def difference_to_value(start:int, difference:list):
    """
    動作
    """
    real_value = start
    list_result = []
    for val in difference:
        real_value += val
        list_result.append(real_value)
    base = copy.deepcopy(list_result)
    ## 後処理
    min_val = min(list_result); max_val = max(list_result)


    ### 負の値の分だけ桁上げ，100を超えたら正規化
    flag_carry = min_val<0; flag_overflow = max_val>100
    # print("carry :",flag_carry,"overflow :",flag_overflow)
    index_min = list_result.index(min_val)
    for index, val in enumerate(list_result):
        ### 負の値の分だけ桁上げ
        carry = abs(min_val) if flag_carry else 0
        ### 最小値の符号が正ならパスする
        if index==index_min and not flag_carry:
            list_result[index] = 0
        ### 最大値が100を超えたら桁上げ＋正規化
        if flag_overflow:
            list_result[index] = int(((val+carry)/(max_val+carry))*100)
        ### 最小値が0を下回ったら桁上げ＋正規化（オーバーフローの可能性もあるので上で対処）
        elif flag_carry:
            list_result[index] = int(((val+carry)/(100+carry))*100)
    if min(list_result)<0:
        print("start value:",start)
        print(min_val,max_val, base, list_result)
        # import numpy as np
        # np.savetxt("test_array.csv",base)
        # exit()
    # print(min_val,max_val, base, list_result)

    return list_result


if __name__ =="__main__":
    """
    test code
    """
    import numpy as np
    test_ = np.loadtxt("test_array.csv")
    print(test_)
    difference_to_value(0,test_)
