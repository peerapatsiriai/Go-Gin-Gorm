package adapters

import (
	"api/entities"
	"api/usecases"

	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) usecases.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (repo *GormOrderRepository) SaveOrder(order *entities.Order) error {
	return repo.db.Create(order).Error
}

func (repo *GormOrderRepository) FindAllOrder() ([]*entities.Order, error) {
	var orders []*entities.Order
	if err := repo.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
