package service

import (
	"api/app/lib"
	"api/app/model"
)

func CreateRoom() *model.Room {
	db := lib.ConnectDB()
	table := db.Table("shamo_room")

	roomId := lib.CreateRandomStr(32)
	room := model.Room{Id: 1, Room_id: roomId}
	table.Put(room).Run()

	return &room
}
