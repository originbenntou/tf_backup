import { gql } from '@apollo/client';

export const GET_LOCAL_TREND_SUGGEST = gql`
  query Trend {
    localTrendSuggest @client {
      keyword
      childSuggests {
        word
        growth {
          short
          medium
          long
        }
        graphs {
          short {
            date
            value
          }
          medium {
            date
            value
          }
          long {
            date
            value
          }
        }
      }
    }
  }
`;
