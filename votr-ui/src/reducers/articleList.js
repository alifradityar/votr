import {
  ARTICLE_UPVOTED,
  ARTICLE_DOWNVOTED,
  SET_PAGE,
  HOME_PAGE_LOADED,
  HOME_PAGE_UNLOADED,
  CHANGE_TAB,
} from '../constants/actionTypes';

export default (state = {}, action) => {
  switch (action.type) {
    case ARTICLE_UPVOTED:
    case ARTICLE_DOWNVOTED:
      return {
        ...state,
        articles: state.articles.map(article => {
          if (article.id === action.payload.data.id) {
            return {
              ...article,
              upvote: action.payload.data.upvote,
              downvote: action.payload.data.downvote
            };
          }
          return article;
        })
      };
    case SET_PAGE:
      return {
        ...state,
        articles: action.payload.topics,
        articlesCount: action.payload.topics.total,
        currentPage: action.page
      };
    case HOME_PAGE_LOADED:
      return {
        ...state,
        pager: action.pager,
        articles: action.payload[0].data.topics,
        articlesCount: action.payload[0].data.total,
        currentPage: 1,
        tab: action.tab
      };
    case HOME_PAGE_UNLOADED:
      return {};
    case CHANGE_TAB:
      return {
        ...state,
        pager: action.pager,
        articles: action.payload.data.topics,
        articlesCount: action.payload.data.total,
        tab: action.tab,
        currentPage: 1
      };
    default:
      return state;
  }
};
