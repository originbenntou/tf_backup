import React, { ChangeEvent, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Button, Card, CardContent, TextField, Box } from '@material-ui/core';
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
  explainCard: {
    backgroundColor: '#FFFBD8',
    padding: '30px',
    color: '#1A1805',
  },
});

const SEND_RECOVER_EMAIL = gql`
  query Account($email: String!, $name: String!) {
    sendRecoverEmail(email: $email, name: $name)
  }
`;

export default function LoginCard() {
  const classes = useStyles();
  const [failureMessageFlag, SetFailureMessageFlag] = useState(false);
  const [authCode, SetAuthCode] = useState('');

  // onClickで値を拾うためonChangeで入力値を格納
  let email: string;
  const emailAction = (event: ChangeEvent<HTMLInputElement>) => {
    email = event.target.value;
  };

  let name: string;
  const nameAction = (event: ChangeEvent<HTMLInputElement>) => {
    name = event.target.value;
  };

  const [sendRecoverEmail] = useLazyQuery(SEND_RECOVER_EMAIL, {
    onCompleted: data => {
      SetAuthCode(data.sendRecoverEmail);
    },

    onError: error => {
      if (error != null) {
        SetFailureMessageFlag(true);
      }
    },
  });

  return (
    <Card className={classes.root}>
      <CardContent>
        <Box m={5} className={classes.explainCard}>
          <p style={{ display: authCode ? 'none' : '' }}>
            メールアドレスとお名前を入力してください。認証コードの発行と新しいパスワードを作成するためのリンクをメールでお送りします。
          </p>
          <p style={{ display: authCode ? '' : 'none' }}>メールを送信しました、メールボックスをご確認ください</p>
        </Box>
        <Box m={5}>
          <TextField className={classes.input} label="メールアドレス" onChange={emailAction} />
          <TextField className={classes.input} label="名前" onChange={nameAction} />
          <div style={{ display: failureMessageFlag ? '' : 'none' }} className={classes.danger}>
            メールアドレスか名前が間違っています
          </div>
          <Box display="flex" justifyContent="center" m={5}>
            <Button
              style={{ display: authCode ? 'none' : '' }}
              variant="contained"
              color="secondary"
              onClick={() => {
                sendRecoverEmail({ variables: { email, name } });
              }}
            >
              更新メールを送信
            </Button>
            <div style={{ display: authCode ? '' : 'none' }}>
              <p>認証コードは{authCode}です</p>
              <p>認証コードはパスワード変更で利用するのでお控えください</p>
            </div>
          </Box>
        </Box>
      </CardContent>
    </Card>
  );
}
