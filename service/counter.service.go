// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package service

import "github.com/SidingsMedia/stats/repository"

type CounterService interface {
}

type counterService struct {
    repo repository.CounterRepository
}

func NewCounterService(repository repository.CounterRepository) CounterService {
    return &counterService{
		repo: repository,
    }
}
