package model

type Room struct {
	Id      int    `dynamo:"id"`
	Room_id string `dynamo:"room_id"`
}
