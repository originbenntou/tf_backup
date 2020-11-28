import React from 'react';
import { Grid, Paper, Divider, Typography } from '@material-ui/core';
import { ResponsiveContainer, LineChart, Line, CartesianGrid, XAxis, YAxis, Tooltip } from 'recharts';
import { makeStyles } from '@material-ui/core/styles';
import { TrendingUp } from '@material-ui/icons';
import moment from 'moment';
import { PAST_TICK_COUNT, FUTURE_TICK_COUNT, TICK_COUNT, EACH_GRAPH_DATA_TYPE, LINE_TYPE } from '../../constants/graph';

type Props = {
  data: EACH_GRAPH_DATA_TYPE;
  lines: LINE_TYPE;
  period: number;
  title: string;
};

const useStyles = makeStyles({
  tableTitle: {
    padding: '10px 20px',
    margin: '0px',
  },
});

/**
 * 横軸の区切りを返す関数
 * @param startDate 開始日
 * @param endDate 終了日
 */
const getTicks = (startDate: moment.Moment, endDate: moment.Moment) => {
  const diffDays = endDate.diff(startDate, 'days');

  const current = startDate;
  const velocity = Math.round(diffDays / TICK_COUNT);

  const ticks = [startDate.valueOf()];

  for (let i = 1; i < TICK_COUNT; i++) {
    ticks.push(current.add(velocity, 'days').valueOf());
  }

  return ticks;
};

/**
 * 日付のフォーマッタ
 * @param date 日付
 */
const dateFormatter = (date: string) => {
  return moment(date).format('YYYY/MM/DD');
};

const Chart: React.FC<Props> = ({ data, lines, period, title }) => {
  const classes = useStyles();
  const Lines = lines.map(line => {
    return <Line key={line.dataKey} type="monotone" dataKey={line.dataKey} stroke={line.color} />;
  });

  // 期間を指定
  const term = period / FUTURE_TICK_COUNT;
  const pastDay = term * PAST_TICK_COUNT;

  // 現在の日付の取得
  const currentDate = data[0]
    ? moment(data[0].date)
        .startOf('day')
        .add(pastDay, 'days')
    : moment().startOf('day');
  const startDate = moment(currentDate).subtract(pastDay, 'days');
  const endDate = moment(currentDate).add(period - 1, 'days');

  // 期間内を満たすデータの作成
  const fillTicksData = data.filter(d => d.date <= endDate.valueOf());
  const domain: [number, number] = [currentDate.valueOf(), endDate.valueOf()];
  const ticks = getTicks(startDate, endDate);

  return (
    <Paper style={{ height: '285px' }}>
      <h3 className={classes.tableTitle}>
        <Grid container alignContent="center">
          <Grid item>
            <TrendingUp />
          </Grid>
          <Grid item>
            <Typography>{`${title}（${period}日間）`}</Typography>
          </Grid>
        </Grid>
      </h3>
      <Divider />
      <ResponsiveContainer width="100%" minHeight={250} height={250}>
        <LineChart data={fillTicksData} margin={{ top: 20, right: 20, bottom: 20, left: 0 }}>
          {Lines}
          <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
          <XAxis
            dataKey="date"
            type="number"
            scale="time"
            tickFormatter={dateFormatter}
            domain={domain}
            ticks={ticks}
          />
          <YAxis type="number" domain={[0, 100]} />
          <Tooltip labelFormatter={t => moment(t).format('YYYY/MM/DD')} />
        </LineChart>
      </ResponsiveContainer>
    </Paper>
  );
};

export default Chart;
