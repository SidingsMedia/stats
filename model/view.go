// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package model


type View struct {
	Page    string    `json:"code" binding:"required,url"`
}