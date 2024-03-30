package orderadapter

import (
	"api/entities"
	orderusecase "api/usecases/order"

	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) orderusecase.OrderRepository {
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
