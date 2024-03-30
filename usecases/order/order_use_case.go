package orderusecase

import (
	"api/entities"
	"errors"
)

type OrderUseCase interface {
	Create(order *entities.Order) error
	// GetByID(id int) (*entities.Order, error)
	GetAll() ([]*entities.Order, error)
	// Update(order *entities.Order) error
	// Delete(id int) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderUseCase {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(order *entities.Order) error {
	if order.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if err := s.repo.SaveOrder(order); err != nil {
		return err
	}
	return nil
}

func (s *OrderService) GetAll() ([]*entities.Order, error) {
	orders, err := s.repo.FindAllOrder()
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, errors.New("no orders found")
	}

	return orders, nil
}
