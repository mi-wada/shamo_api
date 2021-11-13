package model

type User struct {
	Id      int    `dynamo:"id"`
	Name    string `dynamo:"name"`
	Room_id string `dynamo:"room_id"`
}
