import { combineReducers } from 'redux';

// 必要なreducerを追加
import userReducer, { UserState } from './user';

export interface RootState {
  user: UserState;
}

const reducer = combineReducers({
  user: userReducer,
});

export default reducer;
