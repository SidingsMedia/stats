// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package repository

import (
	"github.com/jackc/pgx/v5"
)

type ViewsRepository interface {
}

type viewsRepository struct {
	db *pgx.Conn
}

func NewViewsRepository(db *pgx.Conn) ViewsRepository {
	repository := &viewsRepository{
		db: db,
	}
	return repository
}