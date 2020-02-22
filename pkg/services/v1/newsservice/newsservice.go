package newsservice

import (
	"context"

	"github.com/bungysheep/news-api/pkg/configs"
	newsmodel "github.com/bungysheep/news-api/pkg/models/v1/news"
	"github.com/bungysheep/news-api/pkg/protocols/redis"
	newsrepository "github.com/bungysheep/news-api/pkg/repositories/v1/newsrepository"
)

// INewsService type
type INewsService interface {
	DoPost(context.Context, *newsmodel.News) error
	DoRead(context.Context, int64) ([]*newsmodel.News, error)
}

type newsService struct {
	NewsRepository newsrepository.INewsRepository
}

// NewNewsService - Create news service
func NewNewsService(newsRepo newsrepository.INewsRepository) INewsService {
	return &newsService{
		NewsRepository: newsRepo,
	}
}

func (newsSvc *newsService) DoPost(ctx context.Context, data *newsmodel.News) error {
	err := redis.RedisClient.Publish(configs.REDISNEWSPOSTCHANNEL, data).Err()
	if err != nil {
		return err
	}

	return nil
}

func (newsSvc *newsService) DoRead(ctx context.Context, page int64) ([]*newsmodel.News, error) {
	result := make([]*newsmodel.News, 0)

	if page < 1 {
		page = 1
	}

	itemNews, err := newsSvc.NewsRepository.GetByID(ctx, 0)
	if err != nil {
		return result, err
	}
	result = append(result, itemNews)

	return result, nil
}
