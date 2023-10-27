// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

// Runtime config
var (
	// Address to bind server to
	BindAddr string
	// Proxies from which to trust alternative client IP headers
	TrustedProxies []string
	// Timescale database username
	TimescaleUname string
	// Timescale database password
	TimescalePwd string
	// Timescale database address in <host>:<port> format
	TimescaleAddr string
	// Name of timescale database to connect to
	TimescaleName string
)

// Environment variables
const (
	BindAddrEnv       = "BIND_ADDRESS"
	TrustedProxiesEnv = "TRUSTED_PROXIES"
	TimescaleUnameEnv = "TIMESCALE_USERNAME"
	TimescalePwdEnv   = "TIMESCALE_PASSWORD"
	TimescaleAddrEnv  = "TIMESCALE_ADDRESS"
	TimescaleNameEnv  = "TIMESCALE_NAME"
)

// Defaults
const (
	DefaultBindAddr       = "[::]:3000"
	DefaultTrustedProxies = "*"
)
