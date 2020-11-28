import { configureStore, getDefaultMiddleware, EnhancedStore } from '@reduxjs/toolkit';
import logger from 'redux-logger';
import reducer, { RootState } from './root';
import 'react-redux';

// ReducerのStateの型をセットしている
declare module 'react-redux' {
  // eslint-disable-next-line @typescript-eslint/no-empty-interface
  interface DefaultRootState extends RootState {}
}

// storeの設定、ここでmiddlewareの設定も行なっている
const setupStore = (): EnhancedStore => {
  const middleware = [...getDefaultMiddleware()];

  // 開発環境の場合、Reduxのログを出力する
  if (process.env.NODE_ENV === 'development') {
    middleware.push(logger);
  }

  return configureStore({
    reducer,
    middleware,
    devTools: true,
  });
};

export default setupStore;
