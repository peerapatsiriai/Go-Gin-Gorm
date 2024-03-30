package useradapter

import (
	"api/entities"
	userusecase "api/usecases/user"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) userusecase.UserRepository {
	return &GormUserRepository{db: db}
}

func (repo *GormUserRepository) SaveUser(user *entities.User) error {
	return repo.db.Create(user).Error
}

func (repo *GormUserRepository) FindOneUser(id int) (*entities.User, error) {
	var user entities.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *GormUserRepository) FindAllUser() ([]*entities.User, error) {
	var users []*entities.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *GormUserRepository) UpdateOneUser(user *entities.User) error {
	return repo.db.Save(user).Error
}

func (repo *GormUserRepository) DeleteOneUser(id int) error {
	return repo.db.Delete(&entities.User{}, id).Error
}
