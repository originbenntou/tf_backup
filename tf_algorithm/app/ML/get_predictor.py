"""
予測モデルを学習するためのモジュール
"""

from .my_functions import *

import matplotlib.pyplot as plt
import copy


# 移動平均モデル
class MA():
    """
    Moving Average Model
    移動平均モデル：将来の予測値は過去の予測値と実績値との誤差により決まるという考え方
    過去から現在に向かって誤差が伝播してくるだけで，学習がないので速い
    また，変化が遅れて伝播するという性質がある
    動作：差分の移動平均値を求めて，次の時刻の値を予測する
    """
    def __init__(self, length=25):
        """
        """

    # 次の時刻のデータを移動平均から予測する
    def pred_one_by_one(self,x:list, return_type:str="trend"):
        """
        # 動作
        差分値の移動平均をとって返す
        # 引数
        - x：日付に紐づいた値，型：list，デフォルト；なし
        想定している要素格納方式はFIFO形式で，要素番号が若いほど時刻的にも若い
        - return_type：戻り値の形式，型：str，デフォルト；trend，種類；trend，ma
        """
        # 初期化
        ## 総時刻の取得
        date_length = len(x)

        ## 最新のデータ
        current_data = x[-1]

        ## xを差分値に変換　
        x = get_difference_in(x)

        ## date_length分の平均値の取得
        avg  = average_pooling(x, stride=1, k_size=date_length, max_itr=1)


        # 予測
        ## date_length-1個のデータの取り出し
        x = x[1:date_length]

        ## date_length-1個のデータを取り出して，avgになるように差分を予測
        ## avgはlist型なので[0]を取り出す
        ## 平均値を元の合計値に戻す→xの合計値を引く→足りない値を取得
        t_n = avg[0]*date_length - sum(x)

        if return_type=="trend":
            ## 最新データに予測差分を合わせる
            current_data += t_n
            return current_data
        elif return_type=="ma":
            return t_n


# 検証パターン1：すべてのデータがあるとして，1時刻先を予測し続ける
def validation(pred_model, input_data, pred_num=50, pred_len=310):
    """
    動作  
    引数  
    pred_model  
    target_data  
    pred_num：移動平均に使う長さ（パターンの長さ）  
    pred_len：予測する回数
    """
    ## →予測に使ったデータ数だけ遅れて伝播する（はずである）
    ## 移動平均の値をため込む配列を用意，そこにどんどん値をいれていく
    ## 初期化
    ma_ptn1 = pred_model; pred_ptn1 = []
    
    ## ループ
    for i in range(pred_len):
        ### 10個毎に予測
        pred_ptn1.append(ma_ptn1.pred_one_by_one(input_data[i:i+pred_num]))
    
    ## 遅れて動くので後ろが足らなくなるので0で補完
    pred_ptn1 = pred_ptn1 + [0 for i in range(pred_num)]
    
    return pred_ptn1


# 検証パターン2：全パターンを表す一定量のデータがあるとして，1時刻先を予測し続けて全パターンを予測再現する
def prediction(pred_model:MA, input_data:list, pred_num:int=90, pred_len:int=270):
    """
    動作  
    ma値を返す  
    引数  
    - pred_model  
    - target_data  
    - pred_num：移動平均に使う長さ（パターンの長さ）  
    - pred_len：予測する回数
    """
    ## →予測に使ったデータと同じ動きをする（はずである）
    ## 予測する配列を用意，そこの値を更新していく
    ## 初期化
    ma_ptn2 = pred_model; pred_ptn2 = []
        
    ## 初期状態
    # print("初期状態",input_data) // test code
    target_data = copy.deepcopy(input_data)
    
    ## 予測のループ
    for i in range(pred_len):
        ### 予測値
        pred = ma_ptn2.pred_one_by_one(input_data, return_type="trend")
        # if pred<0:
        #     print(i, "トレンド予測値がマイナスになっています")
        #     print("pred",pred)
        #     print("input_data", input_data)
        #     exit()
        ### 予測値の保存
        pred_ptn2.append(pred)
        ### 予測値の代入
        input_data = input_data[1:]
        input_data.append(pred)
    pred_ptn2 = target_data[:pred_num] + pred_ptn2


    return pred_ptn2


# 描画
def show_graph(target_data, pred_data=[], name="test.png"):
    """
    グラフを描く
    """
    plt.plot(target_data)
    if not len(pred_data)==0:
        plt.plot(pred_data)
    plt.savefig(name)
    plt.show()


if __name__ =="__main__":
    """
    test code
    """

    # moving average
    ## データの読み込み
    target_data = np.loadtxt("sin.csv")


    ma_ptn1 = MA()
    ma_ptn2 = MA()
    input_data = target_data[:180].tolist()


    show_graph(target_data, validation(ma_ptn1, target_data))
    show_graph(target_data, prediction(ma_ptn2, input_data, pred_num=180, pred_len=180))
