package newsrepository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	newsmodel "github.com/bungysheep/news-api/pkg/models/v1/news"
	"github.com/bungysheep/news-api/pkg/protocols/elasticsearch"
	"github.com/bungysheep/news-api/pkg/protocols/redis"
)

var (
	ctx  context.Context
	repo INewsRepository
	db   *sql.DB
	mock sqlmock.Sqlmock
	data []*newsmodel.News
)

func TestMain(m *testing.M) {
	ctx = context.TODO()

	db, mock, _ = sqlmock.New()
	defer db.Close()

	repo = NewNewsRepository(db, redis.RedisClient, elasticsearch.ESClient)

	data = append(data, &newsmodel.News{
		ID:      1,
		Author:  "Author A",
		Body:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent sapien mauris, viverra vel egestas sit amet, mattis sed libero. Etiam sed diam et felis venenatis porta. Aliquam semper sem eget lectus tristique vulputate. Aliquam euismod nisi at justo congue tempus. Sed faucibus non sapien sit amet condimentum. Sed rutrum ligula odio, sit amet bibendum diam sagittis a. Phasellus sit amet risus tellus. Ut elementum venenatis arcu vitae vulputate. Nulla venenatis, magna et luctus gravida, mi lorem molestie ipsum, sed malesuada erat justo id nunc. Integer sodales sem ac ipsum dapibus lobortis. Vivamus auctor felis non magna ultricies, laoreet posuere tellus posuere.",
		Created: time.Now(),
	})

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestNewsRepository(t *testing.T) {
	t.Run("Get News by ID", getByIDTest(ctx))
}

func getByIDTest(ctx context.Context) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("Get fail", getByIDFail(ctx, data[0]))

		t.Run("Get unexisting", getByIDUnexisting(ctx, data[0]))

		t.Run("Get row error", getByIDRowError(ctx, data[0]))

		t.Run("Get existing", getByIDExisting(ctx, data[0]))
	}
}

func getByIDFail(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		expQuery := mock.ExpectPrepare(
			`SELECT id, author, body, created
			FROM news`).ExpectQuery()
		expQuery.WithArgs(input.GetID()).WillReturnError(fmt.Errorf("Failed reading news"))

		res, err := repo.GetByID(ctx, input.GetID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}

		if res != nil {
			t.Errorf("Expect news is nil")
		}
	}
}

func getByIDUnexisting(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "body", "created"})

		expQuery := mock.ExpectPrepare(
			`SELECT id, author, body, created
			FROM news`).ExpectQuery()
		expQuery.WithArgs(input.GetID()).WillReturnRows(rows)

		res, err := repo.GetByID(ctx, input.GetID())
		if err != nil {
			t.Errorf("Expect error is nil")
		}

		if res != nil {
			t.Errorf("Expect news is nil")
		}
	}
}

func getByIDRowError(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		tmNow := time.Now().In(time.UTC)

		rows := sqlmock.NewRows([]string{"id", "author", "body", "created"}).
			AddRow(input.GetID(), input.GetAuthor(), input.GetBody(), tmNow).
			RowError(0, fmt.Errorf("Failed reading news"))

		expQuery := mock.ExpectPrepare(
			`SELECT id, author, body, created
			FROM news`).ExpectQuery()
		expQuery.WithArgs(input.GetID()).WillReturnRows(rows)

		res, err := repo.GetByID(ctx, input.GetID())
		if err == nil {
			t.Errorf("Expect error is not nil")
		}

		if res != nil {
			t.Errorf("Expect news is nil")
		}
	}
}

func getByIDExisting(ctx context.Context, input *newsmodel.News) func(t *testing.T) {
	return func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "author", "body", "created"}).
			AddRow(input.GetID(), input.GetAuthor(), input.GetBody(), input.GetCreated())

		expQuery := mock.ExpectPrepare(
			`SELECT id, author, body, created
			FROM news`).ExpectQuery()
		expQuery.WithArgs(input.GetID()).WillReturnRows(rows)

		res, err := repo.GetByID(ctx, input.GetID())
		if err != nil {
			t.Fatalf("Failed to read news: %v", err)
		}

		if res == nil {
			t.Fatalf("Expect news is not nil")
		}

		if res.GetID() != input.GetID() {
			t.Errorf("Expect news id %d, but got %d", input.GetID(), res.GetID())
		}

		if res.GetAuthor() != input.GetAuthor() {
			t.Errorf("Expect author %s, but got %s", input.GetAuthor(), res.GetAuthor())
		}

		if res.GetBody() != input.GetBody() {
			t.Errorf("Expect body %s, but got %s", input.GetBody(), res.GetBody())
		}

		if res.GetCreated() != input.GetCreated() {
			t.Errorf("Expect created %s, but got %s", input.GetCreated(), res.GetCreated())
		}
	}
}
