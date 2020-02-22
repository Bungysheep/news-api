package newsservice

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	newsmodel "github.com/bungysheep/news-api/pkg/models/v1/news"
	"github.com/bungysheep/news-api/pkg/repositories/v1/newsrepository/mock_newsrepository"
	"github.com/golang/mock/gomock"
)

var (
	ctx  context.Context
	data []*newsmodel.News
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	data = append(data, &newsmodel.News{
		ID:      1,
		Author:  "Author A",
		Body:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent sapien mauris, viverra vel egestas sit amet, mattis sed libero. Etiam sed diam et felis venenatis porta. Aliquam semper sem eget lectus tristique vulputate. Aliquam euismod nisi at justo congue tempus. Sed faucibus non sapien sit amet condimentum. Sed rutrum ligula odio, sit amet bibendum diam sagittis a. Phasellus sit amet risus tellus. Ut elementum venenatis arcu vitae vulputate. Nulla venenatis, magna et luctus gravida, mi lorem molestie ipsum, sed malesuada erat justo id nunc. Integer sodales sem ac ipsum dapibus lobortis. Vivamus auctor felis non magna ultricies, laoreet posuere tellus posuere.",
		Created: time.Now(),
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestContactSystemService(t *testing.T) {
	t.Run("DoPost News failed", doPostFailed(ctx, data[0]))

	t.Run("DoPost News success", doPostSuccess(ctx, data[0]))

	t.Run("DoRead News failed get news ids", doReadFailedGetIds(ctx, data[0]))

	t.Run("DoRead News failed get news by id", doReadFailedGetNewsByID(ctx, data[0]))

	t.Run("DoRead News success", doReadSuccess(ctx, data[0]))
}

func doPostFailed(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_newsrepository.NewMockINewsRepository(ctl)

		repo.EXPECT().Publish(input).Return(fmt.Errorf("Failed publish to redis"))

		newsSvc := NewNewsService(repo)

		err := newsSvc.DoPost(ctx, input)
		if err == nil {
			t.Errorf("Expect error is not nil")
		}
	}
}

func doPostSuccess(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_newsrepository.NewMockINewsRepository(ctl)

		repo.EXPECT().Publish(input).Return(nil)

		newsSvc := NewNewsService(repo)

		err := newsSvc.DoPost(ctx, input)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}
	}
}

func doReadFailedGetIds(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		repo := mock_newsrepository.NewMockINewsRepository(ctl)

		repo.EXPECT().GetIDsByPage(ctx, 1).Return(nil, fmt.Errorf("Failed retrieving to news ids from elasticsearch"))

		newsSvc := NewNewsService(repo)

		news, err := newsSvc.DoRead(ctx, 1)
		if err == nil {
			t.Errorf("Expect error is not nil")
		}

		if news == nil {
			t.Fatalf("Expect news is not nill")
		}

		if len(news) > 0 {
			t.Errorf("Expect no news")
		}
	}
}

func doReadFailedGetNewsByID(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		newsIDs := []int64{1, 2}

		repo := mock_newsrepository.NewMockINewsRepository(ctl)

		repo.EXPECT().GetIDsByPage(ctx, 1).Return(newsIDs, nil)

		repo.EXPECT().GetByID(ctx, newsIDs[0]).Return(input, nil)
		repo.EXPECT().GetByID(ctx, newsIDs[1]).Return(nil, fmt.Errorf("Failed retrieving to news by id"))

		newsSvc := NewNewsService(repo)

		news, err := newsSvc.DoRead(ctx, 1)
		if err == nil {
			t.Errorf("Expect error is not nil")
		}

		if news == nil {
			t.Fatalf("Expect news is not nill")
		}

		if len(news) > 1 {
			t.Errorf("Expect 1 news")
		}
	}
}

func doReadSuccess(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		ctl := gomock.NewController(t)
		defer ctl.Finish()

		newsIDs := []int64{1, 2}

		repo := mock_newsrepository.NewMockINewsRepository(ctl)

		repo.EXPECT().GetIDsByPage(ctx, 1).Return(newsIDs, nil)

		repo.EXPECT().GetByID(ctx, newsIDs[0]).Return(input, nil)
		repo.EXPECT().GetByID(ctx, newsIDs[1]).Return(input, nil)

		newsSvc := NewNewsService(repo)

		res, err := newsSvc.DoRead(ctx, 1)
		if err != nil {
			t.Fatalf("Expect error is nil, but got %v", err)
		}

		if res == nil {
			t.Fatalf("Expect news is not nill")
		}

		if len(res) != 2 {
			t.Errorf("Expect 2 news, but got %d", len(res))
		}

		if res[0].GetID() != input.GetID() {
			t.Errorf("Expect news id %d, but got %d", input.GetID(), res[0].GetID())
		}

		if res[0].GetAuthor() != input.GetAuthor() {
			t.Errorf("Expect author %s, but got %s", input.GetAuthor(), res[0].GetAuthor())
		}

		if res[0].GetBody() != input.GetBody() {
			t.Errorf("Expect body %s, but got %s", input.GetBody(), res[0].GetBody())
		}

		if res[0].GetCreated() != input.GetCreated() {
			t.Errorf("Expect created %s, but got %s", input.GetCreated(), res[0].GetCreated())
		}
	}
}
