import React from 'react';
import { Button } from '@material-ui/core';

type Props = {
  keyword: string;
  onClick: VoidFunction;
};

/**
 * 検索完了コンポーネント
 */
const Completed: React.FC<Props> = props => {
  const { keyword, onClick } = props;
  return (
    <Button variant="contained" color="secondary" onClick={onClick}>
      {keyword}
    </Button>
  );
};

export default Completed;
