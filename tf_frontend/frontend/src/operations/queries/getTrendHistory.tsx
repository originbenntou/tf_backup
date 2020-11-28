import { gql } from '@apollo/client';

export const GET_TREND_HISTORY = gql`
  query Trend {
    trendHistory {
      suggestId
      keyword
      date
      status
      isRead
    }
  }
`;
