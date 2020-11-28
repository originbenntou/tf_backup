import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { TrendingFlat } from '@material-ui/icons';
import { Typography } from '@material-ui/core';
import { pink, blue } from '@material-ui/core/colors';

type Props = {
  growth: string;
};

const useStyles = makeStyles({
  down: {
    color: pink[500],
    transform: 'rotate(45deg)',
  },
  up: {
    color: blue[500],
    transform: 'rotate(-45deg)',
  },
});

const GrowthArrow: React.FC<Props> = ({ growth }) => {
  const classes = useStyles();

  let icon;
  if (growth === 'FLAT') {
    icon = <TrendingFlat />;
  }

  if (growth === 'DOWN') {
    icon = <TrendingFlat className={classes.down} />;
  }

  if (growth === 'UP') {
    icon = <TrendingFlat className={classes.up} />;
  }
  return <Typography align="center">{icon}</Typography>;
};

export default GrowthArrow;
