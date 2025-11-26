package repository

import "get-getById-with-clean-architecture/model"

type UserRepository struct{
	users []model.User
}


func NewUserRepository() *UserRepository {
	users := []model.User{
		{ID: 1, Name: "User 1", Age: 20},
		{ID: 2, Name: "User 2", Age: 21},
		{ID: 3, Name: "User 3", Age: 22},
		{ID: 4, Name: "User 4", Age: 23},
		{ID: 5, Name: "User 5", Age: 24},
		{ID: 6, Name: "User 6", Age: 25},
		{ID: 7, Name: "User 7", Age: 26},
		{ID: 8, Name: "User 8", Age: 27},
		{ID: 9, Name: "User 9", Age: 28},
		{ID: 10, Name: "User 10", Age: 29},
	}

	return &UserRepository{users: users}
}

func (s *UserRepository) GetAllUsers() []model.User{
	return s.users
}

func (s *UserRepository) GetUserByID(id int) *model.User{
	for _, u := range s.users{
		if u.ID == id {
			return &u
		}
	}
	return nil
}
