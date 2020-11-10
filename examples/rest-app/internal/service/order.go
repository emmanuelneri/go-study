package service

import (
	"app/internal/repository"
	"app/pkg/domain"
)

type OrderService interface {
	Save(order *domain.Order) error
}

type OrderServiceImpl struct {
	repository repository.OrderRepository
}

func NewOrderServiceImpl(repository repository.OrderRepository) OrderService {
	return &OrderServiceImpl{repository: repository}
}

func (s OrderServiceImpl) Save(order *domain.Order) error {
	if err := order.CreationValidate(); err != nil {
		return err
	}

	return s.repository.Save(order)
}
