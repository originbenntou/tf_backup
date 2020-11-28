import createSetTrendSuggest from './setTrendSuggest/setTrendSuggest';
import createUpdateSearchWord from './updateSearchWord/updateSearchWord';
import createAddSelectGraph from './addSelectGraph/addSelectGraph';
import createDeleteSelectGraph from './deleteSelectGraph/deleteSelectGraph';
import createHideSelectGraph from './hideSelectGraph/hideSelectGraph';
import createResetSelectGraphs from './resetSelectGraphs/resetSelectGraphs';
import { searchWordVar, trendSuggestVar, selectGraphsVar } from '../../cache';

export const trendSearchMutations = {
  updateSearchWord: createUpdateSearchWord(searchWordVar),
  setTrendSuggest: createSetTrendSuggest(trendSuggestVar),
  addSelectGraph: createAddSelectGraph(selectGraphsVar),
  deleteSelectGraph: createDeleteSelectGraph(selectGraphsVar),
  hideSelectGraph: createHideSelectGraph(selectGraphsVar),
  resetSelectGraphs: createResetSelectGraphs(selectGraphsVar),
};
