import React from 'react';
import { Box, Button, CircularProgress } from '@material-ui/core';

type Props = {
  keyword: string;
};

/**
 * 検索進行中コンポーネント
 */
const InProgress: React.FC<Props> = props => {
  const { keyword } = props;
  return (
    <Button variant="contained" color="secondary" disabled>
      <Box mr={1}>
        <CircularProgress size={14} />
      </Box>
      {keyword}
    </Button>
  );
};

export default InProgress;
