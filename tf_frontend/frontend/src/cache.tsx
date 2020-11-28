import { InMemoryCache, ReactiveVar, makeVar } from '@apollo/client';
import { SelectGraphs } from './models/SelectGraphs';
import { TrendSuggest } from './models/TrendSuggest';
import { SearchWord } from './models/SearchWord';

// 子サジェストの初期値
const initChildSuggest = {
  word: '',
  growth: {
    short: '',
    medium: '',
    long: '',
    __typename: 'Growth',
  },
  graphs: {
    short: [],
    medium: [],
    long: [],
    __typename: 'Graphs',
  },
  __typename: 'ChildSuggest',
};

// サジェストの初期値
const initSuggest = {
  keyword: '',
  childSuggests: [
    initChildSuggest,
    initChildSuggest,
    initChildSuggest,
    initChildSuggest,
    initChildSuggest,
    initChildSuggest,
  ],
  __typename: 'Suggest',
};

// キャッシュを作成するときに、初期値をセットする
const searchWordInitialValue: SearchWord = '';
const suggestIdInitialValue = 0;
const trendSuggestInitialValue: TrendSuggest = [
  initSuggest,
  initSuggest,
  initSuggest,
  initSuggest,
  initSuggest,
  initSuggest,
];
const selectGraphsInitialValue: SelectGraphs = [];

// Reactive Variableを作成
export const searchWordVar: ReactiveVar<SearchWord> = makeVar<SearchWord>(searchWordInitialValue);
export const suggestIdVar: ReactiveVar<number> = makeVar<number>(suggestIdInitialValue);
export const trendSuggestVar: ReactiveVar<TrendSuggest> = makeVar<TrendSuggest>(trendSuggestInitialValue);
export const selectGraphsVar: ReactiveVar<SelectGraphs> = makeVar<SelectGraphs>(selectGraphsInitialValue);

// typePoliciesの作成
export const cache: InMemoryCache = new InMemoryCache({
  typePolicies: {
    Query: {
      fields: {
        selectGraphs: {
          read() {
            return selectGraphsVar();
          },
        },
        searchWord: {
          read() {
            return searchWordVar();
          },
        },
        localTrendSuggest: {
          read() {
            return trendSuggestVar();
          },
        },
      },
    },
  },
});
