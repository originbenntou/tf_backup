import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Box, Grid, Paper, Divider, Typography } from '@material-ui/core';
import MessageIcon from '@material-ui/icons/Message';
import { useQuery } from '@apollo/client';
import GrowthArrow from '../Common/GrowthArrow';
import SuggestCheckbox from './SuggestCheckbox';
import { LINE_COLORS } from '../../constants/graph';
import { SelectGraph } from '../../models/SelectGraphs';
import { ChildSuggest } from '../../models/TrendSuggest';
import { GET_SELECT_GRAPHS } from '../../operations/queries/getSelectGraphs';
import { trendSearchMutations } from '../../operations/mutations/index';

const useStyles = makeStyles({
  tablePaper: {
    minHeight: '260px',
  },
  tableTitle: {
    padding: '5px 20px',
    margin: '0px',
  },
  messageIcon: {
    width: '16px',
  },
  tableHeader: {
    padding: '0px 20px',
  },
  growthCell: {
    width: '10px',
  },
  contentCell: {
    padding: '0px 20px',
  },
});

type Props = {
  keyword: string;
  childSuggests: ChildSuggest[];
};

const SuggestCard: React.FC<Props> = props => {
  const { addSelectGraph, deleteSelectGraph } = trendSearchMutations;
  const classes = useStyles();
  const { data } = useQuery(GET_SELECT_GRAPHS);
  const { keyword, childSuggests } = props;

  /**
   * チェックボックス押下時
   * @param row チェックされたサジェストデータ
   */
  const onClickCheckbox = (row: ChildSuggest) => () => {
    const key = `${keyword} ${row.word}`;
    const isIncludeSameKey = data.selectGraphs.some((graph: SelectGraph) => graph.key === key);

    // チェック済みのデータか判定
    if (isIncludeSameKey) {
      deleteSelectGraph(key);
    } else {
      const addGraph: SelectGraph = {
        key,
        keyword,
        color: LINE_COLORS[data.selectGraphs.length],
        isVisible: true,
        graphs: row.graphs,
      };
      addSelectGraph(addGraph);
    }
  };

  return (
    <Box m={1}>
      <Paper className={classes.tablePaper}>
        <h3 className={classes.tableTitle}>
          <Grid container alignContent="center">
            <Grid item>
              <MessageIcon className={classes.messageIcon} />
            </Grid>
            <Grid item>
              <Typography>{keyword}</Typography>
            </Grid>
          </Grid>
        </h3>
        <Grid container className={classes.tableHeader}>
          <Grid item xs={8}>
            <Typography>キーワード</Typography>
          </Grid>
          <Grid item xs={4}>
            <Grid container>
              <Grid item xs={4}>
                <Typography align="center">短</Typography>
              </Grid>
              <Grid item xs={4}>
                <Typography align="center">中</Typography>
              </Grid>
              <Grid item xs={4}>
                <Typography align="center">長</Typography>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
        <Divider />
        {childSuggests.map((row, key) => (
          /* eslint 'react/no-array-index-key': 'off' */
          <Grid container alignItems="center" className={classes.contentCell} key={`${row.word}-${key}`}>
            <Grid item xs={8}>
              <Grid container alignItems="center">
                <Grid item xs={2}>
                  <SuggestCheckbox
                    onClickCheckbox={onClickCheckbox}
                    keyword={keyword}
                    childSuggest={row}
                    selectGraphs={data.selectGraphs}
                  />
                </Grid>
                <Grid item xs={10}>
                  <Typography>{row.word}</Typography>
                </Grid>
              </Grid>
            </Grid>
            <Grid item xs={4}>
              <Grid container>
                <Grid item xs={4}>
                  <GrowthArrow growth={row.growth.short} />
                </Grid>
                <Grid item xs={4}>
                  <GrowthArrow growth={row.growth.medium} />
                </Grid>
                <Grid item xs={4}>
                  <GrowthArrow growth={row.growth.long} />
                </Grid>
              </Grid>
            </Grid>
          </Grid>
        ))}
      </Paper>
    </Box>
  );
};

export default SuggestCard;
