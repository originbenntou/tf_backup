import { ReactiveVar } from '@apollo/client';
import { SelectGraph, SelectGraphs } from '../../../models/SelectGraphs';

export default (selectGraphsVar: ReactiveVar<SelectGraphs>) => {
  return (addGraph: SelectGraph) => {
    const selectGraphs = selectGraphsVar();
    const newSelectGraphs: SelectGraphs = [...selectGraphs, addGraph];
    selectGraphsVar(newSelectGraphs);
  };
};
