import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Box, Grid, Paper } from '@material-ui/core';
import { useQuery } from '@apollo/client';
import { GET_TREND_HISTORY } from '../../operations/queries/getTrendHistory';
import { Histories } from '../../models/SearchHistories';
import Pagination from './Pagination';

const useStyles = makeStyles({
  basePaper: {
    minHeight: '55px',
  },
});

const SearchWait = () => {
  const classes = useStyles();
  const { data = { trendHistory: [] } } = useQuery<{ trendHistory: Histories }>(GET_TREND_HISTORY, {
    pollInterval: 10000,
  });
  const histories = data.trendHistory;

  // トレンド予測終了で、まだ結果を見ていない件数
  const unReadCompletedTaskNum = histories.filter(history => history.status === 'COMPLETED' && !history.isRead).length;
  // トレンド予測中件数
  const inProgressTaskNum = histories.filter(history => history.status === 'IN_PROGRESS').length;

  return (
    <Paper className={classes.basePaper}>
      <Box>
        <Grid container alignContent="center" alignItems="center">
          <Grid item xs={3}>
            <Box m={1}>
              <Grid>検索待機中ワード</Grid>
              <Grid>
                完了（{unReadCompletedTaskNum}/{unReadCompletedTaskNum + inProgressTaskNum}）
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={9}>
            <Pagination histories={histories} />
          </Grid>
        </Grid>
      </Box>
    </Paper>
  );
};

export default SearchWait;
