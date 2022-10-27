package repository

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	client *sql.DB
}

func NewStorage(c *sql.DB) *Storage {
	return &Storage{
		client: c,
	}
}

func (s *Storage) Ping() error {
	err := s.client.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}
	return nil
}
