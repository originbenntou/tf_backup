import React from 'react';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import { parseCookies, destroyCookie } from 'nookies';
import ReissuePasswordCard from '../src/components/Auth/ReissuePasswordCard';

export default function Index() {
  return (
    <Container maxWidth="sm">
      <Box my={4}>
        <ReissuePasswordCard />
      </Box>
    </Container>
  );
}

Index.getInitialProps = async function(ctx: any) {
  const cookies = parseCookies(ctx);

  // cookieが存在すれば削除する
  if (cookies.VTKT != null) {
    destroyCookie(null, 'VTKT', { path: '/' });
  }

  return {};
};
