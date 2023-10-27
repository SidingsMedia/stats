// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/SidingsMedia/stats/service"
)

type CounterController interface {
}

type counterController struct {
	service service.CounterService
}

func NewCounterController(engine *gin.Engine, service service.CounterService) {
    // controller := &counterController{
    //     service: service,
    // }
    // api := engine.Group("counter")
    // {
    // }
}
