# -*- coding:utf-8 -*-
"""
トレンドを取得するためのモジュール  
参考URL：https://pypi.org/project/pytrends/  
クラスの基本構造と2つの関数get_trend_on_day,get_trend_on_hourは拝借した  
"""
from pytrends.request import TrendReq
import datetime
import pytrends


class GetTrend():
    """
    Googleトレンド指数の取得  
    """

    def __init__(self,hl='ja-JP', tz=360):
        """
        # 動作
        メンバ変数を初期化する
        # 引数
        hl：host language
        tz：time zone
        """
        ## リクエスト設定
        self.req = TrendReq(hl, tz, timeout=(2,10))

        ## 月日数配列 2 4 6 9 11のみ30(28)にちしかない
        self.days_in_month = [31, 28, 31, 30, 31, 30, 31, 31, 20, 31, 30, 31]


    # 月またぎ，年またぎのチェック
    def check_days_in_month(self, date_info, day_term=0):
        """
        # 動作  
        年またぎ，月またぎの扱いについて  
        ※ただし，googleトレンドが勝手にやってくれるから必要ない模様  
        # 引数  
        date_info：  
        day_term：  
        """
        ## パラメータの取得
        year_, month_, day_ = date_info

        ## 日付の差分をとる
        day_dif = day_ - day_term

        ## もし差分がマイナスになるのであれば一ヶ月前に戻す必要がある 
        if day_dif<=0:
        
            ### 1月→12月に差し戻す　＝　1年戻す
            if month_==1:
                year_  = year_ - 1
                month_ = 12+1 # ここはあとで調整が入るので先に足しておく
            ### 閏年の場合は2月は29日間ある
            elif month_==3 and year_%4==0:
                self.days_in_month[1] = 29
            ### 一ヶ月前に戻す
            month_ = month_ - 1
            ### 月末に差分を足す
            day_ = self.days_in_month[month_-1] + day_dif

        else:
            day_ = day_ - day_term


        return year_, month_, day_


    # 1日毎にトレンドを取得
    def get_trend_on_day(self, kw_list=["python"], day_term=1):
        """
        # 動作  
        トレンドをリストで返す  
        FIFOで○○日前から順に入っている  
        # 引数  

        """
        # トレンド回収終了日（登録日の前日→データが出揃っているため）
        end_datetime = datetime.datetime.now() - datetime.timedelta(days=1)

        # トレンド回収開始日
        start_datetime = end_datetime - datetime.timedelta(days=day_term)


        # リクエスト詳細設定
        self.req.build_payload(kw_list, cat=0, timeframe='{} {}'.format(start_datetime.date(), end_datetime.date()), geo='JP', gprop='')
        # cat : カテゴリ
        # timeframe : 開始時刻, today 5-y=5年前 ⇄ 2015-04-02と同義
        # geo : ex)US:united states, JP:japan
        # gprop : What Google property to filter


        # 1日おきのトレンド取得
        trend_results = self.req.interest_over_time()


        return trend_results


    # 1時間毎にトレンドを取得
    def get_trend_on_hour(self, kw_list=["python"], day_term=1):
        """
        # 動作  
        start day の0時からend dayの0時までのデータを収集する  
        # 引数  

        """

        ## トレンド回収終了日（登録日の前日→データが出揃っているため）
        end_datetime =datetime.datetime.now()
        year_end = end_datetime.year; month_end = end_datetime.month; day_end = end_datetime.day 
        year_end, month_end, day_end = self.check_days_in_month([year_end, month_end, day_end],20)
        end_day = "{}-{}-{}".format(year_end, month_end, day_end)
        print("end day : "+ end_day)

        ## トレンド回収開始日
        year_start, month_start, day_start = self.check_days_in_month([year_end, month_end, day_end], day_term)
        start_day = "{}-{}-{}".format(year_start, month_start, day_start)
        print("start day : "+ start_day)

        ## リクエスト詳細設定
        self.req.build_payload(kw_list, cat=0, timeframe='{} {}'.format(start_day, end_day), geo='JP', gprop='')
        

        trend_results = self.req.get_historical_interest(kw_list, 
                    year_start=year_start, month_start=month_start, 
                    day_start=day_start, hour_start=0,
                    year_end=year_end, month_end=month_end, 
                    day_end=day_end, hour_end=0, 
                    cat=0, geo='JP', gprop='', sleep=0)
        

        return trend_results


    # 欠損を探す
    def check_missing_value(self, target_list):
        """
        # 動作  
        欠損があったら補完する，欠損は一回しかない想定  
        # 引数：  
        """
        ## 初期化
        counter = 0; zero_index = 0
        save_counter = 0; save_index = 0

        ## 探索のループ
        for i in target_list:
            ### 0がきたとき
            if target_list[i]==0:
                ### 初めてきたとき
                if counter==0: zero_index = i; counter += 1
                ### 2回目のとき
                elif counter>0:
                    ### 連番なら更新
                    if i==zero_index+1: counter += 1; zero_index = i
                    ### 連番じゃなければリセット
                    else: counter = 0; zero_index = 0
                    ### 3回目以降は保存する
                    if counter>1: save_counter=counter; save_index=i
            ### 0以外は無視
            else :pass
        
        return save_counter, save_index


    # 欠損の補完
    def complete_missing_value(self, target_list):
        ## 欠損を探す
        ## 個数と位置
        missing_num, missing_index = self.check_missing_value(target_list)

        ## 欠損がなければ終わり
        if missing_num==0:
            return 0

        ## 前半か後半かの判断
        first_half = self.check_missing_position(target_list, missing_index)


        ## 欠損があればSMAで補完する


    # 欠損が配列のどの位置にあるのか
    def check_missing_position(self, target_list, missing_index):
        """
        # 動作：欠損が配列のの半分より前ならTrue，後ろならFalseを返す
        """
        ## 前半分のとき
        if len(target_list)/2 <= missing_index:
            return True
        ## 後ろ半分のとき
        elif len(target_list)/2 <= missing_index:
            return False



if __name__ == "__main__":
    """
    test code
    """
    gt = GetTrend()
    # results = gt.get_trend_on_hour(["zoom","bluejeans"],14)
    # results = gt.get_trend_on_day(["zoom","bluejeans"], 76)
    results = gt.get_trend_on_day(["zoom","bluejeans"], 151)
    print(len(results.zoom.values))
    print(results.zoom.values)
    print(int(len(results.zoom.values)/24))
    # import numpy  as np
    # np.savetxt("sample.csv",results.zoom.values)
    # print(results.zoom.values)
