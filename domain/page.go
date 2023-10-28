// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package domain

import (
	"fmt"

	"github.com/SidingsMedia/stats/types"
)

type Page struct {
	ID        int
	Domain    Domain
	Schema    types.Schema
	Port      uint16
	Path      string
}

// Return the URL of the page
func (p Page) Url() string {
	return fmt.Sprintf("%s://%s:%d/%s", p.Schema.String(), p.Domain.Domain, p.Port, p.Path)
}
