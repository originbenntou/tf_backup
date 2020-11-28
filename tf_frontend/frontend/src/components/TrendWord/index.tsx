import React from 'react';
import { TrendingUp } from '@material-ui/icons';
import { Button, Grid, Box } from '@material-ui/core';
import axios from 'axios';
import { useState, useEffect } from 'react';
import { trendSearchMutations } from '../../operations/mutations/index';

const TrendWord = () => {
  const [trendWordList, setTrendWordList] = useState<string[]>([]);
  const { updateSearchWord } = trendSearchMutations;
  useEffect(() => {
    const getTrendKeyword = async () => {
      const requestBody = {
        geo: 'JP',
      };
      const res = await axios.post('/google/daily/trend', requestBody);
      const keyword: string[] = [];
      res.data.data.default.trendingSearchesDays[0].trendingSearches.splice(0, 5).forEach((trend: any) => {
        keyword.push(trend.title.query);
      });
      setTrendWordList(keyword);
    };
    getTrendKeyword();
  }, []);

  const TrendWordBoxList = trendWordList.map(trendWord => (
    <Box ml={3} key={trendWord}>
      <Button
        variant="contained"
        color="secondary"
        onClick={() => {
          updateSearchWord(trendWord);
        }}
      >
        {trendWord}
      </Button>
    </Box>
  ));
  return (
    <Grid container direction="row" alignItems="center">
      <Grid item>
        <TrendingUp />
      </Grid>
      <Grid item>トレンドワード</Grid>
      {TrendWordBoxList}
    </Grid>
  );
};

export default TrendWord;
