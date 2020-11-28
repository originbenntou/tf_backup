import { gql } from '@apollo/client';

export const GET_SELECT_GRAPHS = gql`
  query GetSelectGraphs {
    selectGraphs @client {
      key
      keyword
      color
      isVisible
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
`;
