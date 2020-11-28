import { ReactiveVar } from '@apollo/client';
import { SelectGraphs } from '../../../models/SelectGraphs';

export default (selectGraphsVar: ReactiveVar<SelectGraphs>) => {
  return () => {
    selectGraphsVar([]);
  };
};
