import { parseCookies } from 'nookies';

export default function Index() {
  return {};
}

Index.getInitialProps = async function(ctx: any) {
  const { res } = ctx;
  const cookies = parseCookies(ctx);

  // cookieが存在すればSSRで /home へリダイレクト
  if (cookies.VTKT != null) {
    res.writeHead(302, { Location: '/home' });
  } else {
    res.writeHead(302, { Location: '/login' });
  }
  res.end();

  return {};
};
