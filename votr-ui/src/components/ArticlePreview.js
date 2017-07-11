import React from 'react';
import { Link } from 'react-router';
import agent from '../agent';
import { connect } from 'react-redux';
import { ARTICLE_UPVOTED, ARTICLE_DOWNVOTED } from '../constants/actionTypes';

const styles = {
  voteContainer: {
    marginTop: '1.2rem'
  },
  vote: {
    margin: '0.2rem'
  }
};

const mapDispatchToProps = dispatch => ({
  upvote: id => dispatch({
    type: ARTICLE_UPVOTED,
    payload: agent.Articles.upvote(id)
  }),
  downvote: id => dispatch({
    type: ARTICLE_DOWNVOTED,
    payload: agent.Articles.downvote(id)
  })
});

const ArticlePreview = props => {
  const article = props.article;

  const handleUpvote = ev => {
    ev.preventDefault();
    props.upvote(article.id)
  };

  const handleDownvote = ev => {
    ev.preventDefault();
    props.downvote(article.id)
  };

  return (
    <div className="article-preview">
      <div className="article-meta">
        <div className="pull-xs-right" style={styles.voteContainer}>
          {article.upvote - article.downvote}
          <button style={styles.vote} className="btn btn-sm btn-outline-primary" onClick={handleUpvote}>
            <i className="ion-arrow-up-b"></i> 
          </button>
          <button style={styles.vote} className="btn btn-sm btn-outline-primary" onClick={handleDownvote}>
            <i className="ion-arrow-down-b"></i>
          </button>
        </div>
      </div>

      <Link className="preview-link">
        <h1>{article.title}</h1>
        <p>Posted at {new Date(article.created).toDateString()}</p>
      </Link>
    </div>
  );
}

export default connect(() => ({}), mapDispatchToProps)(ArticlePreview);
