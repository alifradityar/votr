import superagentPromise from 'superagent-promise';
import _superagent from 'superagent';

const superagent = superagentPromise(_superagent, global.Promise);

const API_ROOT = 'http://localhost:7007';

const responseBody = res => res.body;

let token = null;
const tokenPlugin = req => {
  if (token) {
    req.set('authorization', `Token ${token}`);
  }
}

const requests = {
  del: url =>
    superagent.del(`${API_ROOT}${url}`).use(tokenPlugin).then(responseBody),
  get: url =>
    superagent.get(`${API_ROOT}${url}`).use(tokenPlugin).then(responseBody),
  put: (url, body) =>
    superagent.put(`${API_ROOT}${url}`, body).use(tokenPlugin).then(responseBody),
  post: (url, body) =>
    superagent.post(`${API_ROOT}${url}`, body).use(tokenPlugin).then(responseBody)
};

const limit = (count, p) => `page=${p || 1}&size=${count || 10}`;
const omitSlug = article => Object.assign({}, article, { slug: undefined })
const Articles = {
  all: page =>
    requests.get(`/topic?${limit(10, page)}`),
  del: slug =>
    requests.del(`/topic/${slug}`),
  feed: () =>
    requests.get('/topic?page=1&size=10'),
  get: slug =>
    requests.get(`/topic/${slug}`),
  upvote: slug =>
    requests.post(`/topic/${slug}/upvote`),
  downvote: slug =>
    requests.post(`/topic/${slug}/downvote`),
  create: topic =>
    requests.post('/topic', { topic })
};

export default {
  Articles,
  setToken: _token => { token = _token; }
};
