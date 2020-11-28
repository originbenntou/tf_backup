DROP DATABASE IF EXISTS trend;
CREATE DATABASE trend DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

use trend;

DROP TABLE IF EXISTS search;
CREATE TABLE search (
  id INT unsigned NOT NULL auto_increment COMMENT '検索ID',
  search_word VARCHAR(255) NOT NULL COMMENT '検索ワード',
  date DATE NOT NULL COMMENT '検索開始日',
  status tinyint(1) NOT NULL DEFAULT 0 COMMENT '検索進捗 0:検索進行中 1:検索完了',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (id),
  UNIQUE KEY (search_word, date)
) COMMENT '検索情報';

DROP TABLE IF EXISTS suggest;
CREATE TABLE suggest (
  id INT unsigned NOT NULL auto_increment COMMENT 'サジェストID',
  search_id INT unsigned NOT NULL COMMENT '検索ID',
  suggest_word VARCHAR(255) NOT NULL COMMENT 'サジェストワード(検索1件あたりサジェストワード6個)',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (id),
  FOREIGN KEY (search_id) REFERENCES search(id)
) COMMENT 'サジェスト';

DROP TABLE IF EXISTS child_suggest;
CREATE TABLE child_suggest (
  id INT unsigned NOT NULL auto_increment COMMENT '子サジェストID',
  suggest_id INT unsigned NOT NULL COMMENT 'サジェストID',
  child_suggest_word VARCHAR(255) NOT NULL COMMENT '小サジェストワード(サジェストワード1個あたり小サジェストワード6個)',
  short tinyint(1) NOT NULL DEFAULT 0 COMMENT '短期伸び率 0:UP,1:FLAT,2:DOWN',
  medium tinyint(1) NOT NULL DEFAULT 0 COMMENT '中期伸び率 0:UP,1:FLAT,2:DOWN',
  `long` tinyint(1) NOT NULL DEFAULT 0 COMMENT '長期伸び率 0:UP,1:FLAT,2:DOWN',
  short_graphs VARCHAR(2047) DEFAULT NULL COMMENT '短期グラフ情報JSON配列',
  medium_graphs VARCHAR(2047) DEFAULT NULL COMMENT '中期グラフ情報JSON配列',
  long_graphs VARCHAR(4095) DEFAULT NULL COMMENT '長期グラフ情報JSON配列',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (id),
  FOREIGN KEY (suggest_id) REFERENCES suggest(id)
) COMMENT '子サジェスト';

DROP TABLE IF EXISTS history;
CREATE TABLE history (
  id INT unsigned NOT NULL auto_increment COMMENT 'トレンド検索履歴ID',
  user_uuid VARCHAR(255) NOT NULL COMMENT 'ユーザーユニークID',
  search_id INT unsigned NOT NULL COMMENT '検索ID',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (id),
  UNIQUE KEY (user_uuid, search_id),
  FOREIGN KEY (user_uuid) REFERENCES account.user(uuid),
  FOREIGN KEY (search_id) REFERENCES search(id)
) COMMENT 'トレンド検索履歴';

/* child_suggestのgraph_longの文字数が大きすぎるため、制約を解除 */
SET GLOBAL sql_mode = 'NO_ENGINE_SUBSTITUTION';
