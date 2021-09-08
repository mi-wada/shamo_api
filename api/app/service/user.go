package service

import (
	"api/app/lib"
	"api/app/model"
)

func CreateUser(user *model.User) *model.User {
	db := lib.ConnectDB()
	closer, _ := db.DB()
	defer closer.Close()

	db.Create(user)

	return user
}

func GetUsersByRoomId(roomId string) []model.User {
	db := lib.ConnectDB()
	closer, _ := db.DB()
	defer closer.Close()

	var users []model.User
	db.Find(&users, "Room_id = ?", roomId)

	return users
}
