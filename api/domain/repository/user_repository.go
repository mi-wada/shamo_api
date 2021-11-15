package repository

import "api/domain/entity"

type UserRepository interface {
	GetAllByRoomId(roomId string) []entity.User
	Save(user *entity.User)
}
