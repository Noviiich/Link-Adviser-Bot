package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Noviiich/Link-Adviser-Bot/storage"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't ping database: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Save(ctx context.Context, p *storage.Page) (err error) {
	q := `INSERT INTO pages (url, username) VALUES (?, ?)`

	_, err = s.db.ExecContext(ctx, q, p.URL, p.UserName)
	if err != nil {
		return fmt.Errorf("can't save page: %w", err)
	}

	return nil
}

func (s *Storage) PickRandom(ctx context.Context, username string) (p *storage.Page, err error) {
	q := `SELECT url FROM pages WHERE username = ? ORDER BY RANDOM() LIMIT 1`

	var url string

	err = s.db.QueryRowContext(ctx, q, username).Scan(&url)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can't get random page: %w", err)
	}

	return &storage.Page{URL: url, UserName: username}, nil

}
