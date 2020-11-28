import React from 'react';
import { Box } from '@material-ui/core';
import { useQuery } from '@apollo/client';
import moment from 'moment';
import Chart from './Chart';
import Label from './Label';
import { EACH_GRAPH_DATA_TYPE, LINE_TYPE } from '../../constants/graph';
import { trendSearchMutations } from '../../operations/mutations';
import { SelectGraphs } from '../../models/SelectGraphs';
import { Graph } from '../../models/Graphs';
import { GET_SELECT_GRAPHS } from '../../operations/queries/getSelectGraphs';

/**
 * Rechart用のグラフデータに整形
 * @param keyword キーワード
 * @param graphs グラフデータ
 * @param eachGraphData Rechart用のデータ
 */
const formatGraphData = (keyword: string, graphs: Graph[], eachGraphData: EACH_GRAPH_DATA_TYPE) => {
  for (let gKey = 0; gKey < graphs.length; gKey++) {
    const graph = graphs[gKey];

    // 日付をGraph表示用にtimestampにする
    const formattedDate = moment(graph.date).valueOf();

    // 一つの日付に、Suggest + ChildSuggestをキーとして、複数のデータを挿入
    // ex) {date: xxxxxx, コロナ 災い: 10, コロナ 感染者: 20}
    const graphObj = eachGraphData.find(gData => gData.date === formattedDate) || { date: formattedDate };
    graphObj[keyword] = graph.value;

    // それぞれのグラフデータに挿入
    eachGraphData[gKey] = graphObj;
  }
};

export default function TrendGraph() {
  const { deleteSelectGraph, hideSelectGraph } = trendSearchMutations;
  const { data = { selectGraphs: [] } } = useQuery<{ selectGraphs: SelectGraphs }>(GET_SELECT_GRAPHS);

  /**
   * グラフラベルの削除ボタン押下時
   * @param label クリックされたラベル
   */
  const handleDelete = (label: string) => () => {
    deleteSelectGraph(label);
  };

  // グラフラベルがクリックされた時(非表示処理)
  const handleClick = (label: string) => () => {
    hideSelectGraph(label);
  };

  const lines: LINE_TYPE = [];
  const shortGraphData: EACH_GRAPH_DATA_TYPE = [];
  const mediumGraphData: EACH_GRAPH_DATA_TYPE = [];
  const longGraphData: EACH_GRAPH_DATA_TYPE = [];
  for (const selectGraph of data.selectGraphs) {
    if (selectGraph.isVisible) {
      lines.push({
        dataKey: selectGraph.key,
        color: selectGraph.color,
      });
    }

    formatGraphData(selectGraph.key, selectGraph.graphs.short, shortGraphData);
    formatGraphData(selectGraph.key, selectGraph.graphs.medium, mediumGraphData);
    formatGraphData(selectGraph.key, selectGraph.graphs.long, longGraphData);
  }

  // グラフのラベル
  const Labels = data.selectGraphs.map(selectGraph => {
    return (
      <Label
        key={selectGraph.key}
        color={selectGraph.color}
        label={selectGraph.key}
        isVisible={selectGraph.isVisible}
        handleDelete={handleDelete}
        handleClick={handleClick}
      />
    );
  });

  return (
    <>
      <Box fontSize={18} fontFamily="Roboto" fontWeight="fontWeightBold">
        トレンド予測
      </Box>
      <Box mb={2} style={{ minHeight: '30px' }}>
        {Labels}
      </Box>
      <Box mb={2}>
        <Chart data={shortGraphData} lines={lines} period={7} title="短期的トレンド" />
      </Box>
      <Box mb={2}>
        <Chart data={mediumGraphData} lines={lines} period={28} title="中期的トレンド" />
      </Box>
      <Box mb={2}>
        <Chart data={longGraphData} lines={lines} period={70} title="長期的トレンド" />
      </Box>
    </>
  );
}
