import React from 'react';
import App from 'next/app';
import Head from 'next/head';
import { Provider } from 'react-redux';
import { ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import { ApolloClient, createHttpLink, ApolloProvider } from '@apollo/client';
import fetch from 'isomorphic-unfetch';
import theme from '../src/theme';
import Header from '../src/components/Header';
import setupStore from '../src/stores/setup';
import { cache } from '../src/cache';
import '../styles/autosuggest.css';

const store = setupStore();

// GraphQLのoperationNameにあわせて、パスを変更する
const customFetch = (uri: string, options: RequestInit) => {
  const { operationName } = JSON.parse(options.body as string);
  const urlTable: { [key: string]: string } = {
    Account: '/account',
    Trend: '/trend',
  };

  let apiUrl = uri;
  if (urlTable[operationName]) {
    apiUrl += urlTable[operationName];
  }

  return fetch(apiUrl, options);
};

const link = createHttpLink({
  uri: process.env.apiUri,
  credentials: 'include',
  fetch: customFetch,
});

const client = new ApolloClient({
  link,
  cache,
});

export default class MyApp extends App {
  componentDidMount() {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector('#jss-server-side');
    if (jssStyles) {
      jssStyles.parentElement!.removeChild(jssStyles);
    }
  }

  render() {
    const { Component, pageProps } = this.props;

    return (
      <>
        <Head>
          <title>Trend Finder</title>
          <meta name="viewport" content="minimum-scale=1, initial-scale=1, width=device-width" />
        </Head>
        <ApolloProvider client={client}>
          <Provider store={store}>
            <ThemeProvider theme={theme}>
              {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
              <CssBaseline />
              <Header />
              <Component {...pageProps} />
            </ThemeProvider>
          </Provider>
        </ApolloProvider>
      </>
    );
  }
}
