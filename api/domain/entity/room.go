package entity

type Room struct {
	Id     string `json:"id" dynamo:"id"`
	RoomId string `json:"room_id" dynamo:"room_id"`
}
