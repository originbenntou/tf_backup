import React from 'react';
import { Grid } from '@material-ui/core';
import SetTrendButton from './SetTrendButton';
import { Histories } from '../../models/SearchHistories';

type Props = {
  histories: Histories;
};

// 待機中ボタン一覧
const ButtonList: React.FC<Props> = props => {
  const { histories } = props;
  const Buttons = histories.map(history => {
    const { suggestId } = history;
    return (
      <Grid item xs={4} key={suggestId}>
        <SetTrendButton history={history} />
      </Grid>
    );
  });

  return (
    <Grid container alignContent="center">
      {Buttons}
    </Grid>
  );
};

export default ButtonList;
