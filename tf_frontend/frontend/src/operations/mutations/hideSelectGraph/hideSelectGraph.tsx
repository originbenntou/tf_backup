import { ReactiveVar } from '@apollo/client';
import { SelectGraphs } from '../../../models/SelectGraphs';

export default (selectGraphsVar: ReactiveVar<SelectGraphs>) => {
  return (key: string) => {
    const selectGraphs = selectGraphsVar();
    const clicked = selectGraphs.find(selectGraph => selectGraph.key === key);

    // クリックされたデータが見つからない場合
    if (!clicked) {
      return;
    }

    // 表示状態を変えた、新しいselectGraphsを作成
    const newSelectGraphs = selectGraphs.map(selectGraph => {
      if (selectGraph.key === key) {
        return { ...clicked, isVisible: !clicked.isVisible };
      }
      return selectGraph;
    });
    selectGraphsVar(newSelectGraphs);
  };
};
