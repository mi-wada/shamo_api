package application

import (
	"api/domain/entity"
	"api/domain/repository"
)

type UserService struct {
	user_repository repository.UserRepository
}

func NewUserService(user_repository repository.UserRepository) *UserService {
	user_service := new(UserService)
	user_service.user_repository = user_repository
	return user_service
}

func (s *UserService) GetUserListByRoomId(roomId string) []entity.User {
	return s.user_repository.GetAllByRoomId(roomId)
}

func (s *UserService) CreateUser(id int, name string, roomId string) *entity.User {
	user := entity.User{
		Id:     id,
		Name:   name,
		RoomId: roomId,
	}

	s.user_repository.Save(&user)

	return &user
}
