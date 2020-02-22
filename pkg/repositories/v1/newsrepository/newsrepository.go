package newsrepository

import (
	"context"
	"database/sql"
	"fmt"

	newsmodel "github.com/bungysheep/news-api/pkg/models/v1/news"
)

// INewsRepository type
type INewsRepository interface {
	GetByID(context.Context, int64) (*newsmodel.News, error)
}

type newsRepository struct {
	DB *sql.DB
}

// NewNewsRepository - Create news repository
func NewNewsRepository(db *sql.DB) INewsRepository {
	return &newsRepository{DB: db}
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
