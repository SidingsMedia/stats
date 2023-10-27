// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package domain

import "github.com/SidingsMedia/stats/types"

type View struct {
	Domain    string
	Schema    types.Schema
	Port      uint16
	Path      string
	UserAgent string
}
