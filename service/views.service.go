// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package service

import (
	"net/url"
	"strconv"

	"github.com/SidingsMedia/stats/domain"
	"github.com/SidingsMedia/stats/model"
	"github.com/SidingsMedia/stats/repository"
	"github.com/SidingsMedia/stats/types"
	"github.com/SidingsMedia/stats/util"
)

type ViewsService interface {
	AddView(model.View) error
}

type viewsService struct {
    r repository.ViewsRepository
}

func (s *viewsService) AddView(req model.View) error {
	pageUrl, err := url.Parse(req.Page)
	if err != nil {
		return util.ErrInvalidURL
	}

	host, err := s.r.GetDomain(pageUrl.Host)
	if err != nil {
		return err
	}
	if host == nil {
		// Domain doesn't exist
		return util.ErrUnauthorisedDomain
	}

	var schema types.Schema
	switch pageUrl.Scheme {
		case "http":
			schema = types.HTTP
			break
		case "https":
			schema = types.HTTPS
			break
		default:
			return util.ErrInvalidSchema	
	}

	var port uint16
	if pageUrl.Port() != "" {
		t, err := strconv.ParseUint(pageUrl.Port(), 10, 16)
		if err != nil {
			return err
		}
		port = uint16(t)
	} else {
		if schema == types.HTTP {
			port = 80
		} else {
			port = 443
		}
	}

	page := domain.Page{
		Domain: *host,
		Schema: schema,
		Port: port,
		Path: pageUrl.Path,
	}
	view := domain.View{
		Page: page,
		UserAgent: req.UserAgent,
	}

	return s.r.AddView(view)
}

func NewViewsService(repository repository.ViewsRepository) ViewsService {
    return &viewsService{
		r: repository,
    }
}
