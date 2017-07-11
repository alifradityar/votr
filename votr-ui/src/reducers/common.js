import {
  APP_LOAD,
  REDIRECT,
  ARTICLE_SUBMITTED,
  DELETE_ARTICLE,
  ARTICLE_PAGE_UNLOADED,
  EDITOR_PAGE_UNLOADED,
  HOME_PAGE_UNLOADED,
} from '../constants/actionTypes';

const defaultState = {
  appName: 'Votr',
  token: null,
  viewChangeCounter: 0
};

export default (state = defaultState, action) => {
  switch (action.type) {
    case APP_LOAD:
      return {
        ...state,
        token: action.token || null,
        appLoaded: true
      };
    case REDIRECT:
      return { ...state, redirectTo: null };
    case ARTICLE_SUBMITTED:
      const redirectUrl = `article/${action.payload.article.slug}`;
      return { ...state, redirectTo: redirectUrl };
    case DELETE_ARTICLE:
      return { ...state, redirectTo: '/' };
    case ARTICLE_PAGE_UNLOADED:
    case EDITOR_PAGE_UNLOADED:
    case HOME_PAGE_UNLOADED:
      return { ...state, viewChangeCounter: state.viewChangeCounter + 1 };
    default:
      return state;
  }
};
