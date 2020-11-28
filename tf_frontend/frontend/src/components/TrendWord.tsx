import React from 'react';
import { TrendingUp } from '@material-ui/icons';
import { Button, Grid, Box } from '@material-ui/core';

export default function TrendWord() {
  return (
    <>
      <Grid container direction="row" alignItems="center">
        <Grid item>
          <TrendingUp />
        </Grid>
        <Grid item>トレンドワード</Grid>
        <Box ml={3}>
          <Button variant="contained" color="secondary" onClick={event => event.preventDefault()}>
            #コロナ
          </Button>
        </Box>
        <Box ml={3}>
          <Button variant="contained" color="secondary" onClick={event => event.preventDefault()}>
            #インフルエンザ
          </Button>
        </Box>
        <Box ml={3}>
          <Button variant="contained" color="secondary" onClick={event => event.preventDefault()}>
            #タピオカ
          </Button>
        </Box>
      </Grid>
    </>
  );
}
