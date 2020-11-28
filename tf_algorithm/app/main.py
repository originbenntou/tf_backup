import base64
import os
import json

from flask import Flask, request
from ML import core
from sql import tx
from sql.model.search import SearchModel
from sql.model.suggest import SuggestModel
from sql.model.child_suggest import ChildSuggestModel

app = Flask(__name__)


@app.route('/', methods=['POST'])
def index():
    envelope = request.get_json()
    if not envelope:
        msg = 'no Pub/Sub message received'
        print(f'error: {msg}')
        return f'Bad Request: {msg}', 400

    if not isinstance(envelope, dict) or 'message' not in envelope:
        msg = 'invalid Pub/Sub message format'
        print(f'error: {msg}')
        return f'Bad Request: {msg}', 400

    pubsub_message = envelope['message']

    print(pubsub_message)

    message_id = ''
    keyword = ''
    search_id = ''

    if isinstance(pubsub_message, dict) and 'data' in pubsub_message:
        message_id = pubsub_message['messageId']
        # カンマ区切りで検索ワードと検索IDが渡される
        data = base64.b64decode(pubsub_message['data']).decode('utf-8').strip().split(',')
        keyword = data[0]
        search_id = data[1]

    print(f'Start: MessageId={message_id}, Keyword={keyword}, SearchId={search_id}')

    return '', 204

    # ML起動
    # results = json.loads(core.trend_find(keyword))

    jtxt = [
        {
            'keyword': ' おもちゃ',
            'childSuggests': [
                {
                    'word': '元気です',
                    'growth': {
                        'short': 'FLAT',
                        'medium': 'FLAT',
                        'long': 'FLAT'
                    },
                    'graphs': {
                        'short': [
                            {
                                'date': '20200810',
                                'value': 0
                            },
                            {
                                'date': '20200810',
                                'value': 0
                            },
                            {
                                'date': '20200810',
                                'value': 0
                            }
                        ],
                        'medium': [
                            {
                                'date': '20200810',
                                'value': 0
                            },
                            {
                                'date': '20200810',
                                'value': 0
                            },
                            {
                                'date': '20200810',
                                'value': 0
                            }
                        ],
                        'long': [
                            {
                                'date': '20200810',
                                'value': 0
                            },
                            {
                                'date': '20200810',
                                'value': 0
                            },
                            {
                                'date': '20200810',
                                'value': 0
                            }
                        ]
                    }
                }
            ]
        }
    ]

    results = jtxt

    # テスト用
    # DBリファレンスのためバックエンド処理を事前にやっておく
    # tx.session.add(SearchModel(
    #     id=search_id,
    #     search_word=keyword.encode('utf-8'),
    #     date='2020-08-13',
    #     status=0
    # ))
    # tx.session.commit()

    # 伸び率変換
    growth_dict = {
        "UP": 0,
        "FLAT": 1,
        "DOWN": 2,
    }

    # 結果をDBに格納
    for result in results:
        sm = SuggestModel(
            search_id=search_id,
            suggest_word=result["keyword"].strip().encode('utf-8')
        )
        tx.session.add(sm)
        tx.session.commit()

        for childSuggest in result["childSuggests"]:
            tx.session.add(ChildSuggestModel(
                suggest_id=sm.id,
                child_suggest_word=childSuggest["word"].strip().encode('utf-8'),
                short=growth_dict[childSuggest["growth"]["short"]],
                medium=growth_dict[childSuggest["growth"]["medium"]],
                long=growth_dict[childSuggest["growth"]["long"]],
                short_graphs=str(childSuggest["graphs"]["short"]),
                medium_graphs=str(childSuggest["graphs"]["medium"]),
                long_graphs=str(childSuggest["graphs"]["long"]),
            ))
            tx.session.commit()

    print(f'Complete: MessageId={message_id}, Keyword={keyword}, SearchId={search_id}')

    return '', 204


if __name__ == '__main__':
    PORT = int(os.getenv('PORT')) if os.getenv('PORT') else 8080

    app.run(host='127.0.0.1', port=PORT, debug=True)
