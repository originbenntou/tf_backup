import { gql } from '@apollo/client';

export const TREND_SEARCH = gql`
  query Trend($keyword: String!) {
    trendSearch(keyword: $keyword)
  }
`;
