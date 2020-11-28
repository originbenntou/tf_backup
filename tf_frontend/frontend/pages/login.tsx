import React from 'react';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import { parseCookies } from 'nookies';
import LoginCard from '../src/components/Auth/LoginCard';

export default function Index() {
  return (
    <Container maxWidth="sm">
      <Box my={4}>
        <LoginCard />
      </Box>
    </Container>
  );
}

Index.getInitialProps = async function(ctx: any) {
  const { res } = ctx;
  const cookies = parseCookies(ctx);

  // cookieが存在すればSSRで /home へリダイレクト
  if (cookies.VTKT != null) {
    res.writeHead(302, { Location: '/home' });
    res.end();
  }

  return {};
};
