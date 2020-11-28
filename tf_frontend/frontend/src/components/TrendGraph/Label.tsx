import React from 'react';
import { Chip } from '@material-ui/core';
import { Visibility, VisibilityOff } from '@material-ui/icons';

type Props = {
  label: string;
  color: string;
  isVisible: boolean;
  handleDelete: Function;
  handleClick: Function;
};

const Label: React.FC<Props> = ({ label, color, isVisible, handleDelete, handleClick }) => {
  // グラフ表示の場合
  if (isVisible) {
    return (
      <Chip
        color="primary"
        size="small"
        label={label}
        style={{ backgroundColor: color }}
        onDelete={handleDelete(label)}
        onClick={handleClick(label)}
        icon={<Visibility />}
        clickable
      />
    );
  }
  return (
    <Chip
      color="primary"
      size="small"
      label={label}
      style={{ backgroundColor: '#CCCCCC' }}
      onDelete={handleDelete(label)}
      onClick={handleClick(label)}
      icon={<VisibilityOff />}
      clickable
    />
  );
};

export default Label;
