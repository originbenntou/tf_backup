import React, { ChangeEvent, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Button, Card, CardContent, TextField, Box } from '@material-ui/core';
import { gql, useMutation } from '@apollo/client';

type Props = {
  rt: string;
};

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

const RECOVER_PASSWORD = gql`
  mutation Account($recoverToken: String!, $authKey: String!, $password: String!) {
    recoverPassword(recoverToken: $recoverToken, authKey: $authKey, password: $password)
  }
`;

const TokenPasswordCard: React.FC<Props> = ({ rt }) => {
  const classes = useStyles();
  const [successChangePasswordFlag, SetSuccessChangePasswordFlag] = useState(false);
  const [failureMessageFlag, SetFailureMessageFlag] = useState(false);
  const recoverToken = rt;

  // onClickで値を拾うためonChangeで入力値を格納
  let authKey: string;
  const authKeyAction = (event: ChangeEvent<HTMLInputElement>) => {
    authKey = event.target.value;
  };

  let password: string;
  const passwordAction = (event: ChangeEvent<HTMLInputElement>) => {
    password = event.target.value;
  };

  const [recoverPassword] = useMutation(RECOVER_PASSWORD, {
    onCompleted: data => {
      SetSuccessChangePasswordFlag(data.recoverPassword);
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
          <p style={{ display: successChangePasswordFlag ? 'none' : '' }}>
            認証コードと新たに設定するパスワードを入力して下さい。
          </p>
          <p style={{ display: successChangePasswordFlag ? '' : 'none' }}>
            パスワードの変更が完了しました。新しいパスワードでログインをご実施ください。
          </p>
        </Box>
        <Box m={5}>
          <TextField className={classes.input} label="認証コード" onChange={authKeyAction} />
          <TextField className={classes.input} label="新規パスワード" onChange={passwordAction} />
          <div style={{ display: failureMessageFlag ? '' : 'none' }} className={classes.danger}>
            照合情報に誤りがあります。
          </div>
          <Box display="flex" justifyContent="center" m={5}>
            <Button
              variant="contained"
              color="secondary"
              style={{ display: successChangePasswordFlag ? 'none' : '' }}
              onClick={() => {
                recoverPassword({ variables: { recoverToken, authKey, password } });
              }}
            >
              パスワードを更新する
            </Button>
          </Box>
        </Box>
      </CardContent>
    </Card>
  );
};

export default TokenPasswordCard;
