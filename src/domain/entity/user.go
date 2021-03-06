package entity

type User struct {
	Id         string `json:"id" dynamo:"id"`
	RoomId     string `json:"room_id" dynamo:"room_id"`
	Name       string `json:"name" dynamo:"name"`
	Email      string `json:"email" dynamo:"email"`
	PictureUrl string `json:"picture_url" dynamo:"picture_url"`
}
