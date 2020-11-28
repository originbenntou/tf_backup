import React, { useState } from 'react';
import { Grid, IconButton } from '@material-ui/core';
import { KeyboardArrowLeft, KeyboardArrowRight } from '@material-ui/icons';
import ButtonList from './ButtonList';
import { pagination } from '../../utils/index';
import { Histories } from '../../models/SearchHistories';

type Props = {
  histories: Histories;
};

// 1ページの件数
const SIZE_PER_PAGE = 3;

const Pagination: React.FC<Props> = props => {
  const { histories } = props;
  const [page, setPage] = useState(1);
  const { pages, hasNextPage } = pagination(histories, page, SIZE_PER_PAGE);
  const hasPrevPage = page > 1;

  return (
    <Grid container alignItems="center">
      <Grid item xs={1}>
        <IconButton size="small" area-label="prev-page" disabled={!hasPrevPage} onClick={() => setPage(page - 1)}>
          <KeyboardArrowLeft />
        </IconButton>
      </Grid>
      <Grid item xs={10}>
        <ButtonList histories={pages} />
      </Grid>
      <Grid item xs={1}>
        <IconButton size="small" area-label="next-page" disabled={!hasNextPage} onClick={() => setPage(page + 1)}>
          <KeyboardArrowRight />
        </IconButton>
      </Grid>
    </Grid>
  );
};

export default Pagination;
