package userusecase

import "api/entities"

type UserUsecase interface {
	Create(user *entities.User) error
	GetByID(id int) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(user *entities.User) error
	Delete(id int) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserUsecase {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user *entities.User) error {
	if err := s.repo.SaveUser(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetByID(id int) (*entities.User, error) {
	user, err := s.repo.FindOneUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAll() ([]*entities.User, error) {
	users, err := s.repo.FindAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) Update(user *entities.User) error {
	if err := s.repo.UpdateOneUser(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Delete(id int) error {
	if err := s.repo.DeleteOneUser(id); err != nil {
		return err
	}
	return nil
}
