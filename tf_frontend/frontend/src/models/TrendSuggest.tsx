import { Graphs } from './Graphs';

export type ChildSuggest = {
  word: string;
  growth: {
    short: string;
    medium: string;
    long: string;
  };
  graphs: Graphs;
};

export type Suggest = {
  keyword: string;
  childSuggests: ChildSuggest[];
};

export type TrendSuggest = Suggest[];
