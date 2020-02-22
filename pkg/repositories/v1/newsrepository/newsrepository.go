package newsrepository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bungysheep/news-api/pkg/configs"
	newsmodel "github.com/bungysheep/news-api/pkg/models/v1/news"
	redisv7 "github.com/go-redis/redis/v7"
	elasticv7 "github.com/olivere/elastic/v7"
)

// INewsRepository type
type INewsRepository interface {
	GetByID(context.Context, int64) (*newsmodel.News, error)
	Publish(*newsmodel.News) error
	GetIDsByPage(context.Context, int) ([]int64, error)
}

type newsRepository struct {
	DB          *sql.DB
	RedisClient *redisv7.Client
	ESClient    *elasticv7.Client
}

// NewNewsRepository - Create news repository
func NewNewsRepository(db *sql.DB, redisClient *redisv7.Client, esClient *elasticv7.Client) INewsRepository {
	return &newsRepository{
		DB:          db,
		RedisClient: redisClient,
		ESClient:    esClient,
	}
}

func (newsRepo *newsRepository) GetByID(ctx context.Context, id int64) (*newsmodel.News, error) {
	result := newsmodel.NewNews()

	conn, err := newsRepo.DB.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed connecting to database, error: %v", err)
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx,
		`SELECT id, author, body, created
		FROM news 
		WHERE id=$1`)
	if err != nil {
		return nil, fmt.Errorf("Failed preparing read news, error: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Failed reading news, error: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("Failed retrieve news record, error: %v", err)
		}
		return nil, nil
	}

	if err := rows.Scan(
		&result.ID,
		&result.Author,
		&result.Body,
		&result.Created); err != nil {
		return nil, fmt.Errorf("Failed retrieve news record value, error: %v", err)
	}

	return result, nil
}

func (newsRepo *newsRepository) Publish(data *newsmodel.News) error {
	err := newsRepo.RedisClient.Publish(configs.REDISNEWSPOSTCHANNEL, data).Err()
	if err != nil {
		return err
	}

	return nil
}

func (newsRepo *newsRepository) GetIDsByPage(ctx context.Context, page int) ([]int64, error) {
	ids := make([]int64, 0)

	fromIdx := (page - 1) * configs.NUMBERRECORDS
	descCreatedSort := elasticv7.NewFieldSort("created").Desc()
	searchResult, err := newsRepo.ESClient.Search().Index("news").SortBy(descCreatedSort).From(fromIdx).Size(configs.NUMBERRECORDS).Do(ctx)
	if err != nil {
		return ids, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var newsData map[string]interface{}
		err := json.Unmarshal(hit.Source, &newsData)
		if err != nil {
			return ids, err
		}

		newsID, err := strconv.ParseInt(fmt.Sprint(newsData["id"]), 10, 64)
		if err != nil {
			return ids, err
		}

		if newsID > 0 {
			ids = append(ids, newsID)
		}
	}

	return ids, nil
}
