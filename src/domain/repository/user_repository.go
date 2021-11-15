package repository

import "api/domain/entity"

type UserRepository interface {
	GetListByRoomId(roomId string) []entity.User
	Save(user *entity.User)
}
