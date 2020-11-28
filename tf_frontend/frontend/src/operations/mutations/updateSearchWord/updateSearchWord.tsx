import { ReactiveVar } from '@apollo/client';
import { SearchWord } from '../../../models/SearchWord';

export default (searchWordVar: ReactiveVar<SearchWord>) => {
  return (searchWord: SearchWord) => {
    searchWordVar(searchWord);
  };
};
