import { createSlice } from '@reduxjs/toolkit';
import { Dispatch } from 'redux';
import axios from 'axios';
import Router from 'next/router';

export type UserState = {
  name: string;
};

// Stateの初期状態
const initialState: UserState = {
  name: '',
};

// Sliceを生成する
const slice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setName: (state, action): UserState => {
      return { ...state, name: action.payload };
    },
    clearName: (state): UserState => {
      return { ...state, name: '' };
    },
  },
});

// Reducerをエクスポートする
export default slice.reducer;

// Action Creatorsをエクスポートする
export const { setName, clearName } = slice.actions;

// ログイン処理
export function login() {
  return async (dispatch: Dispatch) => {
    try {
      // ランダムユーザAPIを利用して、適当な名前を取得
      const response = await axios('https://randomuser.me/api/');
      const { name } = response.data.results[0];

      // ホーム画面へ遷移
      Router.push('/home');
      // setNameアクションへ値を渡す
      dispatch(setName(`${name.first} ${name.last}`));
    } catch (err) {
      // エラー処理
    }
  };
}

// ログアウト処理
export function logout() {
  return async (dispatch: Dispatch) => {
    try {
      Router.push('/');
      dispatch(clearName());
    } catch (err) {
      // エラー処理
    }
  };
}
