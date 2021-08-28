package main

import (
	"api/app/lib"
	"api/app/model"
	"fmt"
)

func createTable() {
	db := lib.ConnectDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(
		&model.User{},
		&model.Room{},
		&model.Payment{},
	)
}

func insertRoom() (roomId string) {
	db := lib.ConnectDB()
	roomId = lib.CreateRandomStr(32)
	room := model.Room{Room_id: roomId}
	db.Create(&room)
	fmt.Println("CREATED ROOM RECORD")
	return
}

func insertUser(roomId string) {
	db := lib.ConnectDB()
	var kahori = model.User{Name: "Kahori", Room_id: roomId}
	var mitsuaki = model.User{Name: "Mitsuaki", Room_id: roomId}
	fmt.Println("CREATED USER RECORD")
	db.Create(&kahori)
	db.Create(&mitsuaki)
}

func insertData() {
	roomId := insertRoom()
	insertUser(roomId)
}

func main() {
	fmt.Println(">>>>> START MIGRATION >>>>>")
	createTable()
	insertData()
	fmt.Println(">>>>> FINISH MIGRATION >>>>>")
}
