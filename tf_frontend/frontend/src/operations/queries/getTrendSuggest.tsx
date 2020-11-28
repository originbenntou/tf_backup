import { gql } from '@apollo/client';

export const GET_TREND_SUGGEST = gql`
  query Trend($suggestId: Int!) {
    trendSuggest(suggestId: $suggestId) {
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
