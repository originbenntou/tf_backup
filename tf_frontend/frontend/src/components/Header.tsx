import React from 'react';
import { createStyles, makeStyles } from '@material-ui/core/styles';
import { AppBar, Toolbar, Typography, Button } from '@material-ui/core';
import { parseCookies, destroyCookie } from 'nookies';
import { useRouter } from 'next/router';

const useStyles = makeStyles(() =>
  createStyles({
    toolbar: {
      minHeight: '32px',
    },
    title: {
      flexGrow: 1,
    },
  }),
);

export default function Header() {
  const classes = useStyles();
  const router = useRouter();

  const logout = () => {
    destroyCookie(null, 'VTKT', { path: '/' });
    router.push('/login');
    return null;
  };

  const reissuePassword = () => {
    router.push('/recovery_password');
    return null;
  };

  // 設定ボタン（ログアウト、パスワード変更とか）
  const Setting = () => {
    const cookies = parseCookies();

    if (cookies && Object.hasOwnProperty.call(cookies, 'VTKT')) {
      return (
        <>
          <Button color="secondary">パスワード変更</Button>
          <Button color="secondary" onClick={logout}>
            ログアウト
          </Button>
        </>
      );
    }
    return (
      <>
        <Button color="secondary" onClick={reissuePassword}>
          パスワード再発行
        </Button>
      </>
    );
  };

  return (
    <AppBar position="static">
      <Toolbar className={classes.toolbar}>
        <Typography className={classes.title} color="secondary">
          Trend Finder
        </Typography>
        <Setting />
      </Toolbar>
    </AppBar>
  );
}
