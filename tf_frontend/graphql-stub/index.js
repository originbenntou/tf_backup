const express = require('express');
const { ApolloServer, gql } = require('apollo-server-express');

const { trendHistory } = require('./data/trendHistory');
const { trendSuggest } = require('./data/trendSuggest');
const { recoveryUser } = require('./data/recoveryUser');
const { verifyUser } = require('./data/verifyUser');

const { trendSearch } = require('./data/trendSearch');
const { registerUser } = require('./data/registerUser');
const { updateUser } = require('./data/updateUser');

// Construct a schema, using GraphQL schema language
const typeDefs = gql`
  type Query {
    trendSearch(keyword: String!): Int!
    trendHistory: [History]!
    trendSuggest(suggestId: Int!): [Suggest!]!
    recoveryUser(email: String!): Boolean!
    verifyUser(email: String!, password: String!): String!
  }

  type Mutation {
    registerUser(user: User!): Boolean!
    updateUser(user: User!): Boolean!
  }

  input User {
    email: String!
    password: String!
    name: String!
    companyId: Int!
  }

  type History {
    suggestId: Int!
    keyword: String!
    date: DateTime!
    status: Progress!
    isRead: Boolean!
  }

  enum Progress {
    IN_PROGRESS
    COMPLETED
  }

  type Suggest {
    keyword: String!
    childSuggests: [ChildSuggest!]!
  }

  type ChildSuggest {
    word: String!
    growth: Growth!
    graphs: Graphs!
  }

  type Growth {
    short: Arrow!
    medium: Arrow!
    long: Arrow!
  }

  type Graphs {
    short: [Graph!]!
    medium: [Graph!]!
    long: [Graph!]!
  }

  type Graph {
    date: DateTime!
    value: Float!
  }

  enum Arrow {
    UP
    FLAT
    DOWN
  }

  scalar DateTime
`;

// Provide resolver functions for your schema fields
const resolvers = {
  Query: {
    trendSearch: () => trendSearch,
    trendSuggest: () => trendSuggest,
    trendHistory: () => trendHistory,
    verifyUser: () => verifyUser,
    recoveryUser: () => recoveryUser,
  },
  Mutation: {
    registerUser: () => registerUser,
    updateUser: () => updateUser,
  },
};

const server = new ApolloServer({ typeDefs, resolvers });

const app = express();

// CORS設定
const corsOptions = {
  origin: 'http://localhost:3000',
  credentials: true,
};

server.applyMiddleware({ app, cors: corsOptions });

app.listen({ port: 4000 }, () =>
  console.log(`🚀 Server ready at http://localhost:4000${server.graphqlPath}`)
);
