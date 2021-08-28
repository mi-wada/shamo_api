package service

import (
	"api/lib"
	"api/model"
)

func CreateUser(user *model.User) *model.User {
	db := lib.ConnectDB()

	db.Create(user)

	return user
}

func GetUsersByRoomId(roomId string) []model.User {
	db := lib.ConnectDB()

	var users []model.User
	db.Find(&users, "Room_id = ?", roomId)

	return users
}
