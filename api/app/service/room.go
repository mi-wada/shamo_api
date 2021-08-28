package service

import (
	"api/app/lib"
	"api/app/model"
)

func CreateRoom(room *model.Room) *model.Room {
	db := lib.ConnectDB()

	db.Create(room)

	return room
}
