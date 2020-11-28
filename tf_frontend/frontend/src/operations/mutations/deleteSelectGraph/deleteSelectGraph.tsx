import { ReactiveVar } from '@apollo/client';
import { SelectGraphs } from '../../../models/SelectGraphs';
import { LINE_COLORS } from '../../../constants/graph';

export default (selectGraphsVar: ReactiveVar<SelectGraphs>) => {
  return (key: string) => {
    const selectGraphs = selectGraphsVar();
    const newSelectGraphs: SelectGraphs = [];
    let colorKey = 0;

    for (const selectGraph of selectGraphs) {
      if (selectGraph.key !== key) {
        newSelectGraphs.push({
          ...selectGraph,
          color: LINE_COLORS[colorKey],
        });
        colorKey += 1;
      }
    }
    selectGraphsVar(newSelectGraphs);
  };
};
