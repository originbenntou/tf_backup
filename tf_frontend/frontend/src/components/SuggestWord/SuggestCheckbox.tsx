import React from 'react';
import { Checkbox } from '@material-ui/core';
import { createStyles, makeStyles } from '@material-ui/core/styles';
import { SelectGraphs } from '../../models/SelectGraphs';
import { ChildSuggest } from '../../models/TrendSuggest';

type Props = {
  keyword: string;
  childSuggest: ChildSuggest;
  selectGraphs: SelectGraphs;
  onClickCheckbox: Function;
};

const useStyles = makeStyles(() =>
  createStyles({
    checkbox: {
      width: '18px',
      height: '18px',
    },
  }),
);

const SuggestCheckbox: React.FC<Props> = ({ keyword, childSuggest, selectGraphs, onClickCheckbox }) => {
  const classes = useStyles();
  const key = `${keyword} ${childSuggest.word}`;
  const isSameKey = selectGraphs.some(graph => graph.key === key);

  let disabled = false;
  let checked = false;
  if (selectGraphs.length >= 5 && !isSameKey) {
    disabled = true;
  }

  if (isSameKey) {
    checked = true;
  }

  // 検索結果が無い場合
  if (!keyword) {
    return null;
  }
  return (
    <Checkbox
      className={classes.checkbox}
      checked={checked}
      disabled={disabled}
      color="primary"
      onChange={onClickCheckbox(childSuggest)}
    />
  );
};

export default SuggestCheckbox;
