import { ReactiveVar } from '@apollo/client';
import { TrendSuggest } from '../../../models/TrendSuggest';

export default (trendSuggestVar: ReactiveVar<TrendSuggest>) => {
  return (trendSuggest: TrendSuggest) => {
    trendSuggestVar(trendSuggest);
  };
};
