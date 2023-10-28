// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

import "errors"

// DB Errors
var ErrTransactionInProgress = errors.New("transaction already in progress")
var ErrNoCurrentTransaction = errors.New("no transaction currently active")

// Service errors
var ErrInvalidURL = errors.New("invalid url provided")
var ErrUnauthorisedDomain = errors.New("unauthorised domain")
var ErrInvalidSchema = errors.New("schema is not one of http, https")
