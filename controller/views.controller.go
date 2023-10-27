// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/SidingsMedia/stats/service"
)

type ViewsController interface {
}

type viewsController struct {
	service service.ViewsService
}

func NewViewsController(engine *gin.Engine, service service.ViewsService) {
    // controller := &counterController{
    //     service: service,
    // }
    // api := engine.Group("counter")
    // {
    // }
}
