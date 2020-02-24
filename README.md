# News API
## Description
There are 2 main functions of these API as below:
1. Publish a News by sending the news through message queue and then the worker consumer will store it into database and indexing
2. Read all News with max 10 records per page sorted DESC by created from database and indexing

## Endpoint
#### POST /v1/news
###### Request
```
curl --location --request POST 'http://localhost:50051/v1/news' \
--header 'Content-Type: application/json' \
--data-raw '{
    "author": "Author A",
    "body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent sapien mauris, viverra vel egestas sit amet, mattis sed libero. Etiam sed diam et felis venenatis porta. Aliquam semper sem eget lectus tristique vulputate.",
    "created": "2020-02-24T15:04:05Z"
}'
```
###### Response
```
{"data":null,"message":"News has been posted.","success":true}
```

#### GET /v1/news
###### Request
```
curl --location --request GET 'http://localhost:50051/v1/news?page=1'
```
###### Response
```
{"data":[{"id":2,"author":"Author B","body":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent sapien mauris, viverra vel egestas sit amet, mattis sed libero. Etiam sed diam et felis venenatis porta. Aliquam semper sem eget lectus tristique vulputate.","created":"2020-02-25T15:04:05Z"},{"id":1,"author":"Author A","body":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent sapien mauris, viverra vel egestas sit amet, mattis sed libero. Etiam sed diam et felis venenatis porta. Aliquam semper sem eget lectus tristique vulputate.","created":"2020-02-24T15:04:05Z"}],"message":"","success":true}
```

## Docker compose
```
docker-compose up --build
```

## Heroku
https://bungysheep-news-api.herokuapp.com/
