import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Paper, Grid } from '@material-ui/core';

const useStyles = makeStyles({
  paper: {
    backgroundColor: '#DDDDDD',
    height: '400px',
    lineHeight: '400px',
    textAlign: 'center',
    fontWeight: 'bold',
    color: '#BBBBBB',
  },
});

export default function TwitterTrend() {
  const classes = useStyles();

  return (
    <Paper className={classes.paper}>
      <Grid>This is Twitter Trend Component...</Grid>
    </Paper>
  );
}
