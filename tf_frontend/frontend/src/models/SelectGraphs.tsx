import { Graphs } from './Graphs';

export type SelectGraph = {
  key: string;
  keyword: string;
  color: string;
  isVisible: boolean;
  graphs: Graphs;
};

export type SelectGraphs = SelectGraph[];
