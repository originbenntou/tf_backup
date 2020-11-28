/**
 * 文字数を制限する
 * @param str 省略対象の文字
 * @param len 文字数制限
 */
export const truncate = (str: string, len: number) => {
  return str.length > len ? `${str.slice(0, len)}…` : str;
};

/**
 * ページング
 * @param items ページング対象アイテム
 * @param currentPage 現在のページ
 * @param sizePerPage 1ページの件数
 */
export const pagination = (items: any[], currentPage = 1, sizePerPage = 1) => {
  // 全体のページ数を取得
  const totalPage = Math.ceil(items.length / sizePerPage);
  // 現在のページ
  let page = currentPage;

  // 範囲外のページが渡された場合
  if (page < 1) {
    page = 1;
  }
  if (page > totalPage) {
    page = totalPage;
  }

  const hasNextPage = page !== totalPage;

  const start = sizePerPage * (page - 1);
  const end = sizePerPage * page;
  const pages = items.slice(start, end);
  return { pages, hasNextPage };
};
