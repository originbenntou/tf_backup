import React from 'react';
import { Box, Button, Grid } from '@material-ui/core';
import { trendSearchMutations } from '../../operations/mutations';

const RemoveAllCheck = () => {
  const { resetSelectGraphs } = trendSearchMutations;
  const handleClick = () => {
    resetSelectGraphs();
  };

  return (
    <Box m={1}>
      <Grid container alignContent="center">
        <Button variant="contained" color="secondary" onClick={handleClick}>
          チェックボックス一括解除
        </Button>
      </Grid>
    </Box>
  );
};

export default RemoveAllCheck;
