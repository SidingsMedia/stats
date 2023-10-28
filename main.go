// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/SidingsMedia/stats/controller"
	"github.com/SidingsMedia/stats/repository"
	"github.com/SidingsMedia/stats/service"
	"github.com/SidingsMedia/stats/util"
)

func init() {
	log.Println("Fetching environment variables")

	// This will fail in a docker container. Perhaps we need to check if
	// we are in a container and only run if not.
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	// Server settings
	util.BindAddr = util.SGetenv(util.BindAddrEnv, util.DefaultBindAddr)
	util.TrustedProxies = strings.Split(
		util.SGetenv(util.TrustedProxiesEnv, util.DefaultTrustedProxies),
		",",
	)

	// Timescale settings
	util.TimescaleUname = util.Mustgetenv(util.TimescaleUnameEnv)
	util.TimescalePwd = util.Mustgetenv(util.TimescalePwdEnv)
	util.TimescaleAddr = util.Mustgetenv(util.TimescaleAddrEnv)
	util.TimescaleName = util.Mustgetenv(util.TimescaleNameEnv)
}

func main() {
	timescale, dbContext, err := util.InitTimescaleDB(util.TimescaleAddr, util.TimescaleUname, util.TimescalePwd, util.TimescaleName)
	defer timescale.Close(dbContext)
	if err != nil {
		log.Fatalf("Failed to connect to timescale database: %s", err)
	}

	counterRepository := repository.NewViewsRepository(timescale)

	counterService := service.NewViewsService(counterRepository)

	engine := gin.Default()
	engine.Use(cors.Default())

	controller.NewViewsController(engine, counterService)

	// Set trusted proxies. If user has set it to * then we can just
	// ignore it as GIN trusts all by default
	if util.TrustedProxies[0] != "*" {
		if err := engine.SetTrustedProxies(util.TrustedProxies); err != nil {
			log.Fatalf("Failed to set trusted proxies. %s", err)
		}
		log.Printf("Trusting the following proxies: %s", util.TrustedProxies)
	}

	log.Printf("Starting server on %s\n", util.BindAddr)
	engine.Run(util.BindAddr)
}
