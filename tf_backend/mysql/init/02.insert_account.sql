/*
# アカウント

2929ユーザー2929Admin, 2929A会社, スタンダードプラン（テスト用アカウント）
2929ユーザーB, 2929B会社, アドバンスプラン
2929ユーザーC, 2929C会社, エンタープライズプラン
*/

use account;
INSERT INTO plan VALUES (0, "Standard", 30000, 3);
INSERT INTO plan VALUES (0, "Advanced", 50000, 5);
INSERT INTO plan VALUES (0, "Enterprise", 100000, 10);

INSERT INTO company VALUES (0, "2929Acompany", 1, "2020-04-01 10:00:00", "2020-04-01 10:00:00");
INSERT INTO company VALUES (0, "2929Bcompany", 2, "2020-04-01 10:00:00", "2020-04-01 10:00:00");
INSERT INTO company VALUES (0, "2929Ccompany", 3, "2020-04-01 10:00:00", "2020-04-01 10:00:00");

/* テスト用アカウント */
INSERT INTO user VALUES (0, "9bsv0s2v7s8002m4ap2g", "2929admin@trend-find.work", "$2a$10$C8nTGobdJ1OZGcPA7emt1ORRUoBEagkYS4Xh/e.97wtYRuByMeKhO", "Admin", "1", "2020-04-01 10:00:00", "2020-04-01 10:00:00");
/* 以下のユーザーはダミー */
INSERT INTO user VALUES (0, "9bsv0s6jivi002g0e6cg", "2929Buser@test.com", "bbbbb", "2929B", "2", "2020-04-01 10:00:00", "2020-04-01 10:00:00");
INSERT INTO user VALUES (0, "9bsv0s6rj0h002jo1ri0", "2929Cuser@test.com", "ccccc", "2929C", "3", "2020-04-01 10:00:00", "2020-04-01 10:00:00");
