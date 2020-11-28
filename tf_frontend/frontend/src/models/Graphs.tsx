export type Graph = {
  date: string;
  value: number;
};

export type Graphs = {
  short: Array<Graph>;
  medium: Array<Graph>;
  long: Array<Graph>;
};
