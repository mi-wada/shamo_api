package entity

import "time"

type Payment struct {
	Id        string    `json:"id"         dynamo:"id"`
	Price     int       `json:"price"      dynamo:"price"`
	RoomId    string    `json:"room_id"    dynamo:"room_id"`
	UserId    string    `json:"user_id"    dynamo:"user_id"`
	What      string    `json:"what"       dynamo:"what"`
	CreatedAt time.Time `json:"created_at" dynamo:"created_at"`
}
