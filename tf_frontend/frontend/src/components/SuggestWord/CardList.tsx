import React from 'react';
import { Grid } from '@material-ui/core';
import { TrendSuggest } from '../../models/TrendSuggest';
import Card from './Card';

type Props = {
  trendSuggest: TrendSuggest;
};

// カード一覧
const CardList: React.FC<Props> = props => {
  const { trendSuggest } = props;
  const Cards = trendSuggest.map((suggest, key) => {
    const { keyword, childSuggests } = suggest;
    return (
      /* eslint 'react/no-array-index-key': 'off' */
      <Grid item xs={6} key={`${keyword}-${key}`}>
        <Card keyword={keyword} childSuggests={childSuggests} />
      </Grid>
    );
  });

  return <Grid container>{Cards}</Grid>;
};

export default CardList;
