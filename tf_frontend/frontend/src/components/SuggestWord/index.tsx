import React from 'react';
import { Box } from '@material-ui/core';
import { useQuery } from '@apollo/client';
import CardList from './CardList';
import { GET_LOCAL_TREND_SUGGEST } from '../../operations/queries/getLocalTrendSuggest';

export default function SuggestWord() {
  const { data } = useQuery(GET_LOCAL_TREND_SUGGEST);
  const trendSuggest = data.localTrendSuggest;

  return (
    <>
      <Box fontSize={18} fontFamily="Roboto" fontWeight="fontWeightBold">
        サジェストワード
      </Box>
      <CardList trendSuggest={trendSuggest} />
    </>
  );
}
