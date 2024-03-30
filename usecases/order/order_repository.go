package orderusecase

import "api/entities"

type OrderRepository interface {
	SaveOrder(order *entities.Order) error
	// FindOneOrder(id int) (*entities.Order, error)
	FindAllOrder() ([]*entities.Order, error)
	// UpdateOneOrder(order *entities.Order) error
	// DeleteOneOrder(id int) error
}
