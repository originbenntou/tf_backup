// グラフ関係の定数を定義するファイル

// グラフの線の色を定義
export const LINE_COLORS = ['#4DA8E0', '#DF5F6A', '#4CB53F', '#3DB1A4', '#FF67C5'];

// ターム数
export const PAST_TICK_COUNT = 2;
export const FUTURE_TICK_COUNT = 7;
export const TICK_COUNT = PAST_TICK_COUNT + FUTURE_TICK_COUNT;

// Rechart用のフォーマットType
export type EACH_GRAPH_DATA_TYPE = {
  date: number;
  [key: string]: number;
}[];

// Rechart: 線のフォーマット
export type LINE_TYPE = {
  dataKey: string;
  color: string;
}[];
