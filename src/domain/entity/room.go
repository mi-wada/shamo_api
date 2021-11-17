package entity

type Room struct {
	Id string `json:"id" dynamo:"id"`
	// Name   string `json:"name" dynamo:"name"`
}
