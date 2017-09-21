import superagentPromise from 'superagent-promise';
import _superagent from 'superagent';

const superagent = superagentPromise(_superagent, global.Promise);

const API_ROOT = process.env.API_ROOT || 'http://localhost:7007';

const responseBody = res => res.body;

const requests = {
  del: url =>
    superagent.del(`${API_ROOT}${url}`).then(responseBody),
  get: url =>
    superagent.get(`${API_ROOT}${url}`).then(responseBody),
  put: (url, body) =>
    superagent.put(`${API_ROOT}${url}`, body).then(responseBody),
  post: (url, body) =>
    superagent.post(`${API_ROOT}${url}`, body).then(responseBody => { console.log(responseBody); return responseBody.body; })
};

const limit = (count, p) => `page=${p || 1}&size=${count || 10}`;
const Articles = {
  all: page =>
    requests.get(`/topic?${limit(10, page)}`),
  feed: () =>
    requests.get('/topic?page=1&size=10'),
  get: id =>
    requests.get(`/topic/${id}`),
  upvote: id =>
    requests.post(`/topic/${id}/upvote`),
  downvote: id =>
    requests.post(`/topic/${id}/downvote`),
  create: topic =>
    requests.post('/topic', topic)
};

export default {
  Articles,
};
