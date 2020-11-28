import React, { ChangeEvent, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Button, Card, CardContent, TextField, Box } from '@material-ui/core';
import { setCookie } from 'nookies';
import { useRouter } from 'next/router';
import { gql, useLazyQuery } from '@apollo/client';

const useStyles = makeStyles({
  root: {
    minWidth: 275,
  },
  input: {
    width: '100%',
    marginBottom: 20,
  },
  danger: {
    backgroundColor: '#ebccd1',
    borderColor: '#f2dede',
    color: '#a94442',
    padding: '10px 0px',
    textAlign: 'center',
  },
});

const VERIFY_USER = gql`
  query Account($email: String!, $password: String!) {
    verifyUser(email: $email, password: $password)
  }
`;

export default function LoginCard() {
  const classes = useStyles();
  const router = useRouter();
  // 5.ユーザー未存在エラー    16. パスワード照合エラー
  const [verificationErrorFlag, SetVerificationErrorFlag] = useState(false);
  // 7.  ログイン上限超過エラー
  const [overLoginUserErrorFlag, SetOverLoginUserErrorFlag] = useState(false);
  // サーバエラー
  const [serverErrorFlag, SetServerErrorFlag] = useState(false);

  // onClickで値を拾うためonChangeで入力値を格納
  let email: string;
  const emailAction = (event: ChangeEvent<HTMLInputElement>) => {
    email = event.target.value;
  };

  let password: string;
  const passwordAction = (event: ChangeEvent<HTMLInputElement>) => {
    password = event.target.value;
  };

  const [verifyUser] = useLazyQuery(VERIFY_USER, {
    onCompleted: data => {
      // tokenが取得できなかったら戻す
      if (data.verifyUser == null) {
        return;
      }

      setCookie(null, 'VTKT', data.verifyUser, {
        maxAge: 6 * 24 * 60 * 60,
        path: '/',
        domain: process.env.cookieDomain,
      });

      router.push('/home');
    },

    onError: error => {
      // 全部非表示にする
      SetServerErrorFlag(false);
      SetVerificationErrorFlag(false);
      SetOverLoginUserErrorFlag(false);

      if (error.networkError) {
        // サーバエラー。ネットワークエラーはエラー返す
        SetServerErrorFlag(true);
      }

      error.graphQLErrors.forEach(error => {
        if (error!.extensions!.code === 16 || error!.extensions!.code === 5) {
          // ユーザ未存在はエンドユーザに隠す
          SetVerificationErrorFlag(true);
        } else if (error!.extensions!.code === 7) {
          // ログイン数上限エラー
          SetOverLoginUserErrorFlag(true);
        }
      });
    },
  });

  return (
    <Card className={classes.root}>
      <CardContent>
        <Box m={5}>
          <TextField className={classes.input} label="メールアドレス" name="login_id" onChange={emailAction} />
          <TextField className={classes.input} label="パスワード" name="password" onChange={passwordAction} />
          <div className={classes.danger} style={{ display: verificationErrorFlag ? '' : 'none' }}>
            メールアドレスかパスワードが間違っています
          </div>
          <div className={classes.danger} style={{ display: overLoginUserErrorFlag ? '' : 'none' }}>
            ログイン数上限エラーです、契約プランをお見直しください
          </div>
          <div className={classes.danger} style={{ display: serverErrorFlag ? '' : 'none' }}>
            サーバエラーです、お手数ですがリロード後再度お試しください
          </div>
          <Box display="flex" justifyContent="center" m={5}>
            <Button
              variant="contained"
              color="secondary"
              onClick={() => {
                verifyUser({ variables: { email, password } });
              }}
            >
              ログイン
            </Button>
          </Box>
        </Box>
      </CardContent>
    </Card>
  );
}
