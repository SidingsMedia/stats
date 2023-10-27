package repository

import (
	"github.com/jackc/pgx/v5"
)

type CounterRepository interface {
}

type counterRepository struct {
	db *pgx.Conn
}

func NewCounterRepository(db *pgx.Conn) CounterRepository {
	repository := &counterRepository{
		db: db,
	}
	return repository
}