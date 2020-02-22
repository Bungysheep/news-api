package newsservice

import (
	"context"

	newsmodel "github.com/bungysheep/news-api/pkg/models/v1/news"
	newsrepository "github.com/bungysheep/news-api/pkg/repositories/v1/newsrepository"
)

// INewsService type
type INewsService interface {
	DoPost(context.Context, *newsmodel.News) error
	DoRead(context.Context, int) ([]*newsmodel.News, error)
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
	return newsSvc.NewsRepository.Publish(data)
}

func (newsSvc *newsService) DoRead(ctx context.Context, page int) ([]*newsmodel.News, error) {
	result := make([]*newsmodel.News, 0)

	ids, err := newsSvc.NewsRepository.GetIDsByPage(ctx, page)
	if err != nil {
		return result, nil
	}

	for _, newID := range ids {
		itemNews, err := newsSvc.NewsRepository.GetByID(ctx, newID)
		if err != nil {
			return result, err
		}

		if itemNews != nil {
			result = append(result, itemNews)
		}
	}

	return result, nil
}
