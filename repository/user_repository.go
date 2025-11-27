package repository

import (
	"get-getById-with-clean-architecture/model"
	"gorm.io/gorm"
)

type UserRepository struct{
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) GetAllUsers() ([]model.User, error){
	var users []model.User
	result := s.db.Find(&users)

	return users, result.Error
}

func (s *UserRepository) GetUserByID(id int) (*model.User, error){
	var user model.User

	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *UserRepository) CreateUser(user *model.User) error{

	result := s.db.Create(user)

	return result.Error
}
