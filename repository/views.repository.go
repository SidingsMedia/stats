// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/SidingsMedia/stats/domain"
)

type ViewsRepository interface {
	AddView(domain.View) error
	GetDomain(string) (*domain.Domain, error)
	GetPageID(domain.Page) (int, error)
}

type viewsRepository struct {
	db *pgx.Conn
}

// Log a view of a given page
func (r *viewsRepository) AddView(view domain.View) error {
	transaction, err := r.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer transaction.Rollback(context.Background())

	// First log view
	_, err = transaction.Exec(
		context.Background(),
		`INSERT INTO
			view (time, page, user_agent) 
		VALUES
			(
				NOW(), 
				(
					find_page($1, $2, $3, $4)
				),
				$5
			)`,
		view.Page.Domain.ID,
		view.Page.Path,
		view.Page.Schema.String(),
		view.Page.Port,
		view.UserAgent,
	)
	if err != nil {
		return err
	}

	// Now increment counter
	_, err = transaction.Exec(
		context.Background(),
		`WITH cte AS 
		(
			SELECT
				find_page($1, $2, $3, $4) AS page_id 
		)
		INSERT INTO
			counter (page, count) 
		VALUES
			(
				(
					SELECT
						page_id 
					FROM
						cte
				),
				1 
			)
			ON CONFLICT (page) DO 
			UPDATE
			SET
				count = counter.count + 1 
			WHERE
				counter.page = 
				(
					SELECT
						page_id 
					FROM
						cte 
				)`,
		view.Page.Domain.ID,
		view.Page.Path,
		view.Page.Schema.String(),
		view.Page.Port,
	)
	if err != nil {
		return err
	}
	return transaction.Commit(context.Background())
}

// Retrieve the domain for a given page
func (r *viewsRepository) GetDomain(q string) (*domain.Domain, error) {
	row := &domain.Domain{}
	err := r.db.QueryRow(context.Background(), "select id, domain from authorised_domain where domain=$1", q).Scan(&row.ID, &row.Domain)

	if err == pgx.ErrNoRows {
		// Domain doesn't exist
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return row, nil
}

func (r *viewsRepository) GetPageID(page domain.Page) (int, error) {
	var id int
	err := r.db.QueryRow(
		context.Background(), `
		SELECT
			id 
		FROM
			page 
			INNER JOIN
				authorised_domain 
				ON authorised_domain.id = page.domain 
		WHERE
			authorised_domain.domain = $1`, 
		page.Domain.Domain,
	).Scan(&id)
	return id, err
}

func NewViewsRepository(db *pgx.Conn) ViewsRepository {
	repository := &viewsRepository{
		db: db,
	}
	return repository
}