package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Закрытие соединения
//CheckDBConnection trying connect to db.
func (s *ServerRepo) Close() {
	s.cancel()
	s.db.Close()
}

// Тест соединения
//CheckDBConnection trying connect to db.
func (s *ServerRepo) CheckDBConnection(ctx context.Context) error {
	err := s.db.PingContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Создание таблиц
func (s *ServerRepo) createTables(ctx context.Context) error {
	db := s.db
	ctx, cancelFunc := context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunc()

	q := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_uuid VARCHAR(36),
		user_enc_id VARCHAR(36),
		date_add timestamp
	)`
	if _, err := db.ExecContext(ctx, q); err != nil {
		return err
	}

	return nil
}

// Создание соединения
func NewServerRepo(ctx context.Context, c string) (*ServerRepo, error) {
	db, err := sql.Open("postgres", c)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)

	sr := &ServerRepo{
		connStr: c,
		db:      db,
		cancel:  cancel,
	}
	if err := sr.CheckDBConnection(ctx); err != nil {
		return nil, err
	}

	if err := sr.createTables(ctx); err != nil {
		return nil, err
	}
	return sr, nil
}
