package entity

type User struct {
	Id     int    `json:"id" dynamo:"id"`
	Name   string `json:"name" dynamo:"name"`
	RoomId string `json:"room_id" dynamo:"room_id"`
}
