package model

type Payment struct {
	Id       string `dynamo:"id"`
	Price    int    `json:"price" dynamo:"price"`
	RoomId   string `json:"room_id" param:"roomId" dynamo:"room_id"`
	UserId   int    `json:"user_id" dynamo:"user_id"`
	What     string `json:"what" dynamo:"what"`
	Is_valid bool
}
