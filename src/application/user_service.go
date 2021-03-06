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
	return s.user_repository.GetListByRoomId(roomId)
}

func (s *UserService) CreateUser(id, name, roomId, email, pictureUrl string) *entity.User {
	user := entity.User{
		Id:         id,
		Name:       name,
		RoomId:     roomId,
		Email:      email,
		PictureUrl: pictureUrl,
	}

	s.user_repository.Save(&user)

	return &user
}

func (s *UserService) Exist(id string) bool {
	return s.user_repository.Exist(id)
}
