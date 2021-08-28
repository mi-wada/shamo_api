package service

import (
	"api/lib"
	"api/model"
)

func CreateRoom(room *model.Room) *model.Room {
	db := lib.ConnectDB()

	db.Create(room)

	return room
}
