package newscontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bungysheep/news-api/pkg/configs"
	"github.com/bungysheep/news-api/pkg/controllers/v1/basecontroller"
	newsmodel "github.com/bungysheep/news-api/pkg/models/v1/news"
	"github.com/bungysheep/news-api/pkg/protocols/database"
	"github.com/bungysheep/news-api/pkg/protocols/elasticsearch"
	"github.com/bungysheep/news-api/pkg/protocols/mq"
	"github.com/bungysheep/news-api/pkg/protocols/redis"
	newsrepository "github.com/bungysheep/news-api/pkg/repositories/v1/newsrepository"
	"github.com/bungysheep/news-api/pkg/services/v1/newsservice"
	redisv7 "github.com/go-redis/redis/v7"
)

// NewsController type
type NewsController struct {
	basecontroller.BaseResource
}

// NewNewsController - Creates news controller
func NewNewsController() *NewsController {
	return &NewsController{}
}

// PostNews - Posting a news
func (newsCtl *NewsController) PostNews(w http.ResponseWriter, r *http.Request) {
	log.Printf("Posting a news.\n")

	news := newsmodel.NewNews()
	err := json.NewDecoder(r.Body).Decode(news)
	if err != nil {
		newsCtl.WriteResponse(w, http.StatusBadRequest, false, nil, "Invalid create news request.")
		return
	}

	valid, message := news.DoValidate()
	if !valid {
		newsCtl.WriteResponse(w, http.StatusBadRequest, false, nil, message)
		return
	}

	newsSvc := newsservice.NewNewsService(newsrepository.NewNewsRepository(database.DbConnection, redis.RedisClient, mq.MqConnection, elasticsearch.ESClient))
	if err := newsSvc.DoPost(r.Context(), news); err != nil {
		newsCtl.WriteResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	newsCtl.WriteResponse(w, http.StatusAccepted, true, nil, "News has been posted.")
}

// GetNews - Retrieve news
func (newsCtl *NewsController) GetNews(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	page, err := strconv.Atoi(queries.Get("page"))

	if page < 1 {
		page = 1
	}

	cachedValue, err := redis.RedisClient.Get(fmt.Sprintf("news_page_%d", page)).Result()
	if err == redisv7.Nil {
		log.Printf("Cache news_page_%d does not exist.\n", page)
	} else if err != nil {
		newsCtl.WriteResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	} else {
		log.Printf("Retrieving News page '%v' from cached.\n", page)

		var cachedResult []interface{}
		json.Unmarshal([]byte(cachedValue), &cachedResult)
		newsCtl.WriteResponse(w, http.StatusOK, true, cachedResult, "")
		return
	}

	log.Printf("Retrieving News page '%v'.\n", page)

	newsSvc := newsservice.NewNewsService(newsrepository.NewNewsRepository(database.DbConnection, redis.RedisClient, mq.MqConnection, elasticsearch.ESClient))
	result, err := newsSvc.DoRead(r.Context(), page)
	if err != nil {
		newsCtl.WriteResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	byteResult, _ := json.Marshal(result)
	if err := redis.RedisClient.Set(fmt.Sprintf("news_page_%d", page), string(byteResult), configs.CACHEEXPIRYIN*time.Second).Err(); err != nil {
		log.Printf("Failed to cache the read response.\n")
	}

	newsCtl.WriteResponse(w, http.StatusOK, true, result, "")
}
