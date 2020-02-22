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

	// REDISNEWSPOSTCHANNEL - Redis news post channel
	REDISNEWSPOSTCHANNEL = "NEWS_POST_CHANNEL"

	// NUMBERRECORDS - Redis news post channel
	NUMBERRECORDS = 3
)
