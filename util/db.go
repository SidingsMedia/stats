// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Initialise the timescale timeseries database connection
func InitTimescaleDB(addr string, uname string, pwd string, dbname string) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", uname, pwd, addr, dbname)

	ctx := context.Background()
    conn, err := pgx.Connect(ctx, dsn)
    return conn, err
}