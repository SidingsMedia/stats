// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package model

type View struct {
	Page      string `json:"page" binding:"required,url"`
	UserAgent string `header:"user-agent"`
}
