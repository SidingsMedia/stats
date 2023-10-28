// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/SidingsMedia/stats/model"
	"github.com/SidingsMedia/stats/service"
	"github.com/SidingsMedia/stats/util"
)

type ViewsController interface {
	AddView(*gin.Context)
}

type viewsController struct {
	service service.ViewsService
}

func (c *viewsController) AddView(ctx *gin.Context) {
	request := &model.View{}
	
	if err := ctx.ShouldBind(request); err != nil && errors.As(err, &validator.ValidationErrors{}){
		util.SendBadRequestFieldNames(ctx, err.(validator.ValidationErrors))
		return
	} else if err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GeneralError{
            Code: http.StatusBadRequest,
            Message: "Request was malformed",
        })
        return
    }

	if err := ctx.ShouldBindHeader(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GeneralError{
            Code: http.StatusBadRequest,
            Message: "Headers are malformed",
        })
        return
	}

	switch err := c.service.AddView(*request); err {
		case util.ErrInvalidSchema:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.BadRequest{
				Code: http.StatusBadRequest,
				Message: "Your request is malformed",
				Fields: []model.Fields{
					{
						Field: "page",
						Condition: "url scheme not one of http, https",
					},
				},
			})
			return
		case util.ErrUnauthorisedDomain:
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.GeneralError{
				Code: http.StatusForbidden,
				Message: "Unrecognized domain for page",
			})
			return
		case nil:
			ctx.Status(http.StatusNoContent)
			return
		default:
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GeneralError{
				Code: http.StatusInternalServerError,
				Message: "Internal server error",
			})
			return
	}
}

func NewViewsController(engine *gin.Engine, service service.ViewsService) {
    controller := &viewsController{
        service: service,
    }
    api := engine.Group("views")
    {
		api.POST("", controller.AddView)
    }
}
