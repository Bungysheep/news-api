package configs

const (
	// PORT - Port number
	PORT = "50051"

	// READTIMEOUT - Read timeout
	READTIMEOUT = 10

	// WRITETIMEOUT - Write timeout
	WRITETIMEOUT = 10

	// DBCONNSTRING - Database connection string
	DBCONNSTRING = "postgres://news-local-pg:news-local-pg@localhost:5432/news-local-pg?sslmode=disable"

	// ELASTICSEARCHURL - Elasticsearch url
	ELASTICSEARCHURL = "http://localhost:9200/"

	// REDISURL - Redis url
	REDISURL = "localhost:6379"

	// REDISAUTH - Redis url
	REDISAUTH = ""

	// RABBITMQURL - RabbitMQ url
	RABBITMQURL = "amqp://news-local-mq:news-local-mq@localhost:5672/"

	// MQNEWSPOSTQUEUE - RabbitMQ news post queue
	MQNEWSPOSTQUEUE = "NEWS_POST_QUEUE"

	// NUMBERRECORDS - Number records per page
	NUMBERRECORDS = 10

	// CACHEEXPIRYIN - Cache expiry in 60 (sec)
	CACHEEXPIRYIN = 60
)
