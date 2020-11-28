import React from 'react';
import { Container, Grid, Box } from '@material-ui/core';
import { parseCookies } from 'nookies';
import SearchForm from '../src/components/SearchForm/index';
import TrendWord from '../src/components/TrendWord/index';
import SuggestWord from '../src/components/SuggestWord/index';
import TrendGraph from '../src/components/TrendGraph/index';
import RegistList from '../src/components/RegisterList/index';
import TwitterTrend from '../src/components/TwitterTrend/index';
import SearchWait from '../src/components/SearchWait/index';
import RemoveAllCheck from '../src/components/RemoveAllCheck/index';

export default function Index() {
  return (
    <Container maxWidth="xl">
      <Box my={2}>
        <Grid container>
          <Grid item xs={4}>
            <SearchForm />
          </Grid>
          <Grid item xs={8}>
            <TrendWord />
          </Grid>
        </Grid>
      </Box>
      <Box>
        <Grid container>
          <Grid item xs={4}>
            <Box mb={2}>
              <SearchWait />
            </Box>
            <SuggestWord />
            <RemoveAllCheck />
          </Grid>
          <Grid item xs={6}>
            <Box m={1}>
              <TrendGraph />
            </Box>
          </Grid>
          <Grid item xs={2}>
            <Box m={1}>
              <RegistList />
            </Box>
            <Box m={1}>
              <TwitterTrend />
            </Box>
          </Grid>
        </Grid>
      </Box>
    </Container>
  );
}

Index.getInitialProps = async function(ctx: any) {
  const { res } = ctx;
  const cookies = parseCookies(ctx);

  // cookieが存在しなければSSRで /login へリダイレクト
  if (cookies.VTKT == null) {
    res.writeHead(302, { Location: '/login' });
    res.end();
  }

  return {};
};
