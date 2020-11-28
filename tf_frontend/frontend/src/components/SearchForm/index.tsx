import React, { useState } from 'react';
import axios from 'axios';
import AutoSuggest from 'react-autosuggest';
import { makeStyles, createStyles } from '@material-ui/core/styles';
import { Paper, Box, Grid, Button } from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import { useLazyQuery, useQuery } from '@apollo/client';
import { GET_SEARCH_WORD } from '../../operations/queries/getSearchWord';
import { TREND_SEARCH } from '../../operations/queries/trendSearch';
import { trendSearchMutations } from '../../operations/mutations';

const useStyles = makeStyles(() =>
  createStyles({
    root: {
      width: '100%',
    },
    input: {
      display: 'flex',
      alignItems: 'center',
    },
    inputIcon: {
      verticalAlign: 'middle',
      marginLeft: 20,
    },
    suggest: {
      display: 'inline-flex',
      verticalAlign: 'middle',
    },
    suggestWord: {
      marginLeft: 8,
    },
  }),
);

export default function SearchInput() {
  const classes = useStyles();
  const [suggestions, setSuggestions] = useState<string[]>([]);
  const { updateSearchWord } = trendSearchMutations;

  // トレンドワードからの検索ワードを取得
  const searchWordQueryResult = useQuery(GET_SEARCH_WORD);
  const { searchWord } = searchWordQueryResult.data;

  // トレンド検索
  const [trendSearch] = useLazyQuery(TREND_SEARCH, {
    onCompleted: data => {
      console.log('data-----==========', data);
      // TODO: 検索ボタンが押された時に、待機中ワードに入れる
    },
  });

  // トレンド検索実行
  const doSearch = (keyword: string) => {
    if (keyword.trim().length > 0) {
      trendSearch({ variables: { keyword } });
    }
  };

  // 入力コンポーネント部分のカスタム
  const renderInputComponent = (inputProps: any) => {
    return (
      <div className={classes.input}>
        <SearchIcon className={classes.inputIcon} />
        <input {...inputProps} />
      </div>
    );
  };

  // suggestionを表示
  const renderSuggestion = (suggestion: string) => {
    return (
      <span className={classes.suggest}>
        <SearchIcon /> <span className={classes.suggestWord}>{suggestion}</span>
      </span>
    );
  };

  // suggestionをGoogle Autocomplete APIを叩いて取得
  async function fetchSuggestions(value: string) {
    // Google Autocomplete API
    const response = await axios.get(`/complete/search?q=${value}}&hl=ja&output=toolbar&client=firefox`);
    const suggestions: string[] = response.data[1];
    setSuggestions(suggestions);
  }

  // suggestionを要求
  const onSuggestionsFetchRequested = ({ value }: { value: string }) => {
    fetchSuggestions(value);
  };

  // suggestion選択
  const onSuggestionSelected = (_: any, { suggestionValue }: { suggestionValue: string }) => {
    updateSearchWord(suggestionValue);
  };

  // input props
  const inputProps = {
    placeholder: 'ワードを入力',
    value: searchWord,
    onChange: (_: any, { newValue }: { newValue: string }) => {
      updateSearchWord(newValue);
    },
  };

  return (
    <Grid container>
      <Grid item xs={10}>
        <Box mr={1}>
          <Paper className={classes.root}>
            <AutoSuggest
              suggestions={suggestions}
              onSuggestionsClearRequested={() => setSuggestions([])}
              onSuggestionsFetchRequested={onSuggestionsFetchRequested}
              onSuggestionSelected={onSuggestionSelected}
              getSuggestionValue={suggestion => suggestion}
              renderSuggestion={renderSuggestion}
              renderInputComponent={renderInputComponent}
              inputProps={inputProps}
              highlightFirstSuggestion
            />
          </Paper>
        </Box>
      </Grid>
      <Grid item xs={2}>
        <Button
          variant="contained"
          color="secondary"
          onClick={() => {
            doSearch(searchWord);
          }}
        >
          予測する
        </Button>
      </Grid>
    </Grid>
  );
}
