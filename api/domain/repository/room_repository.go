package repository

import "api/domain/entity"

type RoomRepository interface {
	Save(user *entity.Room)
}
