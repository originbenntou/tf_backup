# -*- coding:utf-8 -*-
"""
サジェストを取得するためのモジュール
参考URL：http://angelpinpoint.seesaa.net/article/453568317.html
"""

import argparse
from time import sleep
from string import ascii_lowercase
from string import digits
import requests
import urllib.parse


class GoogleAutoComplete:
    """
    GoogleAutoCompleteのワードを持ってくる  
    """

    def __init__(self):
        self.base_url = 'https://www.google.co.jp/complete/search?'\
                        'hl=ja&output=toolbar&ie=utf-8&oe=utf-8&'\
                        'client=firefox&q='

    def get_suggest(self, query):
        buf = requests.get(self.base_url + urllib.parse.quote_plus(query)).json()
        suggests = [ph for ph in buf[1]]
        return suggests

    def get_suggest_with_one_char(self, query):
        # キーワードそのものの場合のサジェストワード
        ret = self.get_suggest(query+' ')

        # # -rオプションがあればもう1段階
        # if self.recurse_mode:
        #     ret = self.get_uniq(ret)  # 事前に重複を除いておく
        #     addonelevel = []
        #     for ph in ret:
        #         addonelevel.extend(self.get_suggest(ph + ' '))
        #     ret.extend(addonelevel)

        return self.get_uniq(ret)

    # 重複を除く
    def get_uniq(self, arr):
        uniq_ret = []
        for x in arr:
            if x not in uniq_ret:
                uniq_ret.append(x)
        return uniq_ret


class YahooSugget:
    """
    YahooSuggestのワードを持ってくる  
    """

    def __init__(self):
        """
        # 動作

        # 引数

        """
        self.base_url = 'http://ff.search.yahoo.com/gossip?output=json&command='

    # サジェストワードを返す
    def get_suggest(self, query):
        """
        # 動作

        # 引数

        """

        buf = requests.get(self.base_url + query + "+").json()
        list_suggests = [words["key"] for words in buf["gossip"]["results"]]
        
        return list_suggests


    def get_suggest_with_one_char(self, query):
        """
        # 動作

        # 引数

        """
        
        ret = self.get_suggest(query+' ')

        return self.get_uniq(ret)


    # 重複を除く
    def get_uniq(self, arr):
        """
        # 動作

        # 引数

        """
        uniq_ret = []
        for x in arr:
            if x not in uniq_ret:
                uniq_ret.append(x)
        return uniq_ret


# 子サジェストの取得
def get_child_suggest(suggest_source, phrase):
    """
    # 動作
    サジェストワードは6個，子サジェストワードは5個取得する
    # 引数
    """

    # サジェストワードの取得
    list_suggest_words = suggest_source.get_suggest_with_one_char(phrase)

    
    # サジェストワードを6つに絞る
    list_suggest_words = list_suggest_words[:6] if len(list_suggest_words)>5 else list_suggest_words


    # 子サジェストワードの取得
    ## 初期化
    dict_child_suggest_words = {}
    # dict_child_suggest_words["source"] = phrase

    ## サジェストワードのループ
    for suggest_word in list_suggest_words:
        ### 子サジェストワードの取得
        child_suggest_words = suggest_source.get_suggest_with_one_char(suggest_word)
        ### 子サジェストワードを5つに絞る
        dict_child_suggest_words[suggest_word] = child_suggest_words[:6] if len(child_suggest_words)>5 else child_suggest_words


    return dict_child_suggest_words


# 子サジェスト取得の実行
def get_google_suggest(phrase):
    """
    # 動作
    GoogleAutocompleteAPIを使って子サジェスト30個を取得する  
    戻り値はdict型：キー（サジェストワード），値（子サジェストワードのリスト）  
    # 引数
    phrase：元になるワード
    """
    dict_child_suggest_words = get_child_suggest(suggest_source=GoogleAutoComplete(), phrase=phrase)
    return dict_child_suggest_words


# 子サジェスト取得の実行
def get_yahoo_suggest(phrase):
    """
    # 動作
    YahooSuggestAPIを使って子サジェスト30個を取得する  
    戻り値はdict型：キー（サジェストワード），値（子サジェストワードのリスト）  
    # 引数
    phrase：元になるワード
    """
    dict_child_suggest_words = get_child_suggest(suggest_source=YahooSugget(), phrase=phrase)
    return dict_child_suggest_words


if __name__ == "__main__":
    """
    test code
    """
    
    parser = argparse.ArgumentParser()
    parser.add_argument("phrase", help="調べたい単語")
    args = parser.parse_args()

    dict_child_suggest_words = get_child_suggest(suggest_source=GoogleAutoComplete(), phrase=args.phrase)
    print(dict_child_suggest_words)
