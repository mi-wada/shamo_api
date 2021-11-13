package model

type User struct {
	Id      int    `dynamo:"id"`
	Name    string `dynamo:"namo"`
	Room_id string `dynamo:"room_id"`
}
