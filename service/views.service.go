// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package service

import "github.com/SidingsMedia/stats/repository"

type ViewsService interface {
}

type viewsService struct {
    repo repository.ViewsRepository
}

func NewViewsService(repository repository.ViewsRepository) ViewsService {
    return &viewsService{
		repo: repository,
    }
}
