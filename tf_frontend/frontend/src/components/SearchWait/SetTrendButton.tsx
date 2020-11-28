import React from 'react';
import { Box, Tooltip } from '@material-ui/core';
import { useLazyQuery } from '@apollo/client';
import { withStyles } from '@material-ui/core/styles';
import { GET_TREND_SUGGEST } from '../../operations/queries/getTrendSuggest';
import { trendSearchMutations } from '../../operations/mutations/index';
import { History } from '../../models/SearchHistories';
import SwitchButton from './SwitchButton';
import { truncate } from '../../utils/index';

type Props = {
  history: History;
};

// 検索待機中ワードの文字数
const KEYWORD_LENGTH = 4;

const HtmlTooltip = withStyles(theme => ({
  tooltip: {
    backgroundColor: theme.palette.common.white,
    boxShadow: theme.shadows[1],
    color: 'rgba(0, 0, 0, 0.87)',
    fontSize: theme.typography.pxToRem(12),
    border: '1px solid #dadde9',
  },
}))(Tooltip);

const SetTrendButton: React.FC<Props> = ({ history }) => {
  const { suggestId, keyword, date, status } = history;
  const { setTrendSuggest, resetSelectGraphs } = trendSearchMutations;
  const [getSuggest] = useLazyQuery(GET_TREND_SUGGEST, {
    onCompleted: data => {
      resetSelectGraphs();
      setTrendSuggest(data.trendSuggest);
    },
  });

  return (
    <HtmlTooltip
      title={
        <>
          <Box m={1}>{keyword}</Box>
          <Box m={1}>検索日：{date}</Box>
        </>
      }
    >
      <span>
        <SwitchButton
          keyword={truncate(keyword, KEYWORD_LENGTH)}
          isProgress={status === 'IN_PROGRESS'}
          onClick={() => {
            getSuggest({ variables: { suggestId } });
          }}
        />
      </span>
    </HtmlTooltip>
  );
};

export default SetTrendButton;
