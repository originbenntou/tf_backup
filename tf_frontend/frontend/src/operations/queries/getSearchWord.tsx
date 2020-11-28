import { gql } from '@apollo/client';

export const GET_SEARCH_WORD = gql`
  query GetSearchWord {
    searchWord @client
  }
`;
