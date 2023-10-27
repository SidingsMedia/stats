// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package types

type Schema int64;

const (
	HTTP Schema = iota
	HTTPS
)

// Return the string representation of the schema in lower case. If an
// invalid value is set return "unknown"
func (s Schema) String() string {
	switch s {
		case HTTP:
			return "http"
		case HTTPS:
			return "https"
		default:
			return "unknown"
	}
}