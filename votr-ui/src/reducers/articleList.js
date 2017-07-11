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
          if (article.slug === action.payload && action.payload.article.slug) {
            return {
              ...article,
              upvoted: action.payload.article.upvoted,
              downvoted: action.payload.article.downvoted
            };
          }
          return article;
        })
      };
    case SET_PAGE:
      return {
        ...state,
        articles: action.payload.articles,
        articlesCount: action.payload.articlesCount,
        currentPage: action.page
      };
    case HOME_PAGE_LOADED:
      console.log("HOME_PAGE_LOADED")
      console.log(action.payload)
      return {
        ...state,
        pager: action.pager,
        articles: action.payload[0].data,
        articlesCount: action.payload[0].articlesCount,
        currentPage: 0,
        tab: action.tab
      };
    case HOME_PAGE_UNLOADED:
      return {};
    case CHANGE_TAB:
      return {
        ...state,
        pager: action.pager,
        articles: action.payload.articles,
        articlesCount: action.payload.articlesCount,
        tab: action.tab,
        currentPage: 0
      };
    default:
      return state;
  }
};
