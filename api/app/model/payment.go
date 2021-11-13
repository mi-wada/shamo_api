package model

type Payment struct {
	Id       string `dynamo:"id"`
	Price    int    `json:"price" dynamo:"price"`
	Room_id  string `json:"room_id" param:"roomId" dynamo:"room_id"`
	User_id  string `json:"user_id" dynamo:"user_id"`
	What     string `json:"what" dynamo:"what"`
	Is_valid bool
}
