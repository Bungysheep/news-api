package newscontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bungysheep/news-api/pkg/controllers/v1/basecontroller"
	newsmodel "github.com/bungysheep/news-api/pkg/models/v1/news"
	"github.com/bungysheep/news-api/pkg/protocols/database"
	"github.com/bungysheep/news-api/pkg/protocols/elasticsearch"
	"github.com/bungysheep/news-api/pkg/protocols/redis"
	newsrepository "github.com/bungysheep/news-api/pkg/repositories/v1/newsrepository"
	"github.com/bungysheep/news-api/pkg/services/v1/newsservice"
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

	newsSvc := newsservice.NewNewsService(newsrepository.NewNewsRepository(database.DbConnection, redis.RedisClient, elasticsearch.ESClient))
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

	log.Printf("Retrieving News page '%v'.\n", page)

	newsSvc := newsservice.NewNewsService(newsrepository.NewNewsRepository(database.DbConnection, redis.RedisClient, elasticsearch.ESClient))
	result, err := newsSvc.DoRead(r.Context(), page)
	if err != nil {
		newsCtl.WriteResponse(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}

	newsCtl.WriteResponse(w, http.StatusOK, true, result, "")
}
