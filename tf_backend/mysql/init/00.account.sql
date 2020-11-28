DROP DATABASE IF EXISTS account;
CREATE DATABASE account DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

use account;

DROP TABLE IF EXISTS plan;
CREATE TABLE plan (
  id INT unsigned NOT NULL auto_increment COMMENT 'プランID',
  name VARCHAR(255) NOT NULL COMMENT 'プラン名',
  price INT unsigned NOT NULL DEFAULT 0 COMMENT 'プラン金額',
  capacity INT unsigned  NOT NULL DEFAULT 3 COMMENT 'ログイン上限',
  PRIMARY KEY (id)
) COMMENT 'プラン情報';

DROP TABLE IF EXISTS company;
CREATE TABLE company (
  id INT unsigned NOT NULL auto_increment COMMENT '会社ID',
  name VARCHAR(255) NOT NULL COMMENT '会社名',
  plan_id INT unsigned NOT NULL COMMENT 'プランID',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (id),
  FOREIGN KEY (plan_id) REFERENCES plan(id)
) COMMENT '会社情報';

DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id INT unsigned NOT NULL auto_increment COMMENT 'ユーザーID（使用しない）',
  uuid VARCHAR(255) NOT NULL COMMENT 'ユニークID',
  email VARCHAR(255) NOT NULL COMMENT 'ユーザーEmailアドレス',
  password VARCHAR(1023) NOT NULL COMMENT 'ユーザーパスワードハッシュ',
  name VARCHAR(255) NOT NULL COMMENT 'ユーザー名',
  company_id INT unsigned NOT NULL COMMENT '会社ID',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (id),
  UNIQUE KEY (uuid),
  UNIQUE KEY (email),
  FOREIGN KEY (company_id) REFERENCES company(id)
) COMMENT 'ユーザー情報';

DROP TABLE IF EXISTS session;
CREATE TABLE session (
  id INT unsigned NOT NULL auto_increment COMMENT 'セッションID',
  token VARCHAR(255) NOT NULL COMMENT 'セッショントークン',
  user_uuid VARCHAR(255) NOT NULL COMMENT 'ユーザーユニークID',
  company_id INT unsigned NOT NULL COMMENT '会社ID',
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時（マイクロ秒）',
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時（マイクロ秒）',
  PRIMARY KEY (id),
  UNIQUE KEY (token),
  FOREIGN KEY (user_uuid) REFERENCES user(uuid),
  FOREIGN KEY (company_id) REFERENCES company(id)
) COMMENT 'セッション情報';

DROP TABLE IF EXISTS recover_session;
CREATE TABLE recover_session (
  id INT unsigned NOT NULL auto_increment COMMENT 'リカバーセッションID',
  user_uuid VARCHAR(255) NOT NULL COMMENT 'ユーザーユニークID',
  auth_key VARCHAR(1023) NOT NULL COMMENT 'リカバー認証キー',
  recover_token VARCHAR(1023) NOT NULL COMMENT 'リカバートークン',
  created_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '作成日時（マイクロ秒）',
  updated_at datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新日時（マイクロ秒）',
  PRIMARY KEY (id),
  UNIQUE KEY (user_uuid),
  FOREIGN KEY (user_uuid) REFERENCES user(uuid)
) COMMENT 'リカバーセッション情報';
