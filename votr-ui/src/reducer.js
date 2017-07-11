import article from './reducers/article';
import articleList from './reducers/articleList';
import { combineReducers } from 'redux';
import common from './reducers/common';
import editor from './reducers/editor';
import home from './reducers/home';

export default combineReducers({
  article,
  articleList,
  common,
  editor,
  home,
});
