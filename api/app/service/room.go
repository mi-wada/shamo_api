package service

import (
	"api/app/lib"
	"api/app/model"
)

func CreateRoom() *model.Room {
	db := lib.ConnectDB()
	closer, _ := db.DB()
	defer closer.Close()

	roomId := lib.CreateRandomStr(32)
	room := model.Room{Room_id: roomId}

	db.Create(&room)

	return &room
}
