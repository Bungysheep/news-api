package newsservice

import (
	"context"
	"sync"

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

// DoPost - Post news
func (newsSvc *newsService) DoPost(ctx context.Context, data *newsmodel.News) error {
	return newsSvc.NewsRepository.Publish(data)
}

// DoRead - Read news
func (newsSvc *newsService) DoRead(ctx context.Context, page int) ([]*newsmodel.News, error) {
	result := make([]*newsmodel.News, 0)

	ids, err := newsSvc.NewsRepository.GetIDsByPage(ctx, page)
	if err != nil {
		return result, err
	}

	tempResult := make([]*newsmodel.News, len(ids))
	resultChan := make(chan *newsmodel.News)
	errChan := make(chan error)
	done := make(chan bool)

	go func() {
		for {
			select {
			case errTemp := <-errChan:
				err = errTemp
				done <- true
				return
			case item, more := <-resultChan:
				if more {
					result = append(result, item)
				} else {
					done <- true
					return
				}
			default:

			}
		}
	}()

	var wg sync.WaitGroup
	for i, newsID := range ids {
		wg.Add(1)
		go func(i int, newsID int64, wg *sync.WaitGroup) {
			defer wg.Done()

			itemNews, err := newsSvc.NewsRepository.GetByID(ctx, newsID)
			if err != nil {
				errChan <- err
				return
			}

			if itemNews != nil {
				tempResult[i] = itemNews
			}
		}(i, newsID, &wg)
	}
	wg.Wait()

	close(resultChan)
	close(errChan)

	<-done

	for _, item := range tempResult {
		if item != nil {
			result = append(result, item)
		}
	}

	return result, err
}
