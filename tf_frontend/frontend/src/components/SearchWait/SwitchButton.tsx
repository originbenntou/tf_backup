import React from 'react';
import InProgress from './InProgress';
import Completed from './Completed';

type Props = {
  keyword: string;
  isProgress: boolean;
  onClick: VoidFunction;
};

/**
 * 検索状態によって、表示するボタンを切り替えるコンポーネント
 */
const SwitchButton: React.FC<Props> = props => {
  const { keyword, isProgress, onClick } = props;
  if (isProgress) {
    return <InProgress keyword={keyword} />;
  }
  return <Completed keyword={keyword} onClick={onClick} />;
};

export default SwitchButton;
