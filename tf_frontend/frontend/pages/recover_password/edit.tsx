import React from 'react';
import { NextPage } from 'next';
import Container from '@material-ui/core/Container';
import Box from '@material-ui/core/Box';
import { parseCookies, destroyCookie } from 'nookies';
import TokenPasswordCard from '../../src/components/Auth/TokenPasswordCard';

type Props = {
  rt: string;
};

const Page: NextPage<Props> = ({ rt }) => {
  return (
    <Container maxWidth="sm">
      <Box my={4}>
        <TokenPasswordCard rt={rt} />
      </Box>
    </Container>
  );
};

Page.getInitialProps = async function(ctx: any) {
  const cookies = parseCookies(ctx);

  // cookieが存在すれば削除する
  if (cookies.VTKT != null) {
    destroyCookie(null, 'VTKT', { path: '/' });
  }

  return { rt: ctx.query.rt };
};

export default Page;
