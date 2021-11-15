package repository

import (
	"api/domain/entity"
	"api/infra/database"

	"github.com/guregu/dynamo"
)

type RoomRepository struct {
	db *dynamo.DB
}

func NewRoomRepository() *RoomRepository {
	room_repository := new(RoomRepository)
	room_repository.db = database.Connect()
	return room_repository
}

func (r *RoomRepository) Save(room *entity.Room) {
	table := r.db.Table("shamo_room")

	table.Put(room).Run()
}
