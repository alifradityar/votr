# Votr API
Upvote and Downvote simple application.

## Technology
- Mysql for main storage
- Redis for cache storage
- Facebook Go Injection Library
- Gorilla mux for routing

## API Definition

```
POST /topic
payload:
{
    "title": "Test"
}
Create new topic


GET /topic?keyword=?&page=?&size=?
List topic with pagination feature based on keyword filter

GET /topic/all
Retrieve all topic sorted by score

GET /topic/<id>
Get single topic based on id

DELETE /topic/<id>
Delete single topic based on id

PUT /topic/<id>
payload:
{
    "title": "TestNew"
}
Update single topic title

POST /topic/upvote
empty payload
Upvote a topic (increase the score)

POST /topic/downvote
empty payload
Downvote a topic (reduce the score)
```

## How to Run ?
1. Setup MySQL and Redis
2. Run migration
3. Put this code under your $GOPATH
4. Setup config (MYSQL_HOST, MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_DATABASE, REDIS_HOST, REDIS_PORT)
5. go run main.go

## Improvement?
- Handle auth (every user should not be allowed to upvote multiple times)
- Use time for scoring
- Use Docker for easy setup :)